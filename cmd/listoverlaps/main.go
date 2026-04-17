package main

import (
    "fmt"
    "os"
    "path/filepath"
    "sort"

    "github.com/feedloop/syde/internal/audit"
    "github.com/feedloop/syde/internal/model"
    "github.com/feedloop/syde/internal/storage"
)

func main() {
    root, _ := filepath.Abs(".syde")
    if _, err := os.Stat(root); err != nil {
        fmt.Fprintln(os.Stderr, "run from project root")
        os.Exit(1)
    }
    store, err := storage.NewStore(root)
    if err != nil { panic(err) }
    defer store.Close()

    allReqs, listErr := store.List(model.KindRequirement)
    if listErr != nil { panic(listErr) }

    type reqEntry struct{ req *model.RequirementEntity; terms map[string]bool }
    var active []reqEntry
    var sets []map[string]bool
    for _, ewb := range allReqs {
        r, ok := ewb.Entity.(*model.RequirementEntity)
        if !ok || r.RequirementStatus != model.RequirementActive || r.Statement == "" {
            continue
        }
        t := audit.SignificantTerms(r.Statement)
        if len(t) == 0 { continue }
        active = append(active, reqEntry{r, t})
        sets = append(sets, t)
    }
    corpus := audit.NewTFIDFCorpus(sets)

    type pair struct{ a, b string; sim float64; acked bool }
    var out []pair
    for i := 0; i < len(active); i++ {
        for j := i + 1; j < len(active); j++ {
            sim := corpus.TFIDFSimilarity(active[i].terms, active[j].terms)
            if sim <= 0.5 { continue }
            aSlug := active[i].req.GetBase().CanonicalSlug()
            bSlug := active[j].req.GetBase().CanonicalSlug()
            acked := false
            for _, ao := range active[i].req.AuditedOverlaps { if ao.Slug == bSlug { acked = true; break } }
            for _, ao := range active[j].req.AuditedOverlaps { if ao.Slug == aSlug { acked = true; break } }
            out = append(out, pair{aSlug, bSlug, sim, acked})
        }
    }
    sort.Slice(out, func(i,j int) bool { return out[i].sim > out[j].sim })
    for _, p := range out {
        ack := ""
        if p.acked { ack = " [acked]" }
        fmt.Printf("%.0f%%\t%s\t%s%s\n", p.sim*100, p.a, p.b, ack)
    }
    fmt.Fprintf(os.Stderr, "Total pairs >50%%: %d\n", len(out))
}
