package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var (
	rememberCategory   string
	rememberEntities   []string
	rememberConfidence string
	rememberSource     string
	rememberTags       []string
)

var rememberCmd = &cobra.Command{
	Use:   `remember "<text>"`,
	Short: "Capture a design learning",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := args[0]

		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		learning := &model.LearningEntity{
			BaseEntity: model.BaseEntity{
				Kind:        model.KindLearning,
				Name:        text,
				Description: text,
				Tags:        rememberTags,
			},
			Category:     model.LearningCategory(rememberCategory),
			EntityRefs:   rememberEntities,
			Source:       model.LearningSource(rememberSource),
			ConfLevel:    model.Confidence(rememberConfidence),
			DiscoveredAt: time.Now().UTC().Format(time.RFC3339),
		}

		// Truncate name if too long for slug
		if len(learning.Name) > 60 {
			learning.Name = learning.Name[:60]
		}

		filePath, err := store.Create(learning, "")
		if err != nil {
			return err
		}

		fmt.Printf("Remembered: %s\n", text)
		fmt.Printf("  Category: %s | Confidence: %s\n", rememberCategory, rememberConfidence)
		fmt.Printf("  File: %s\n", filePath)

		return nil
	},
}

var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "Manage learnings",
}

var learnListCmd = &cobra.Command{
	Use:   "list",
	Short: "List learnings",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		learnings, err := store.List(model.KindLearning)
		if err != nil {
			return err
		}

		if len(learnings) == 0 {
			fmt.Println("No learnings found.")
			return nil
		}

		for _, ewb := range learnings {
			l := ewb.Entity.(*model.LearningEntity)
			icon := "ℹ"
			switch l.Category {
			case model.CatGotcha:
				icon = "⚠"
			case model.CatConstraint:
				icon = "⚠"
			}
			fmt.Printf("  %s %-12s %-50s [%s]\n", icon, l.Category, l.Description, l.ConfLevel)
		}
		return nil
	},
}

var learnAboutCmd = &cobra.Command{
	Use:   "about <entity>",
	Short: "Show learnings for an entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		entitySlug := args[0]

		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		// Find the entity to get its ID
		entity, _, err := store.Get(entitySlug)
		if err != nil {
			return fmt.Errorf("entity not found: %s", entitySlug)
		}
		entityID := entity.GetBase().ID

		learnings, err := store.List(model.KindLearning)
		if err != nil {
			return err
		}

		found := 0
		for _, ewb := range learnings {
			l := ewb.Entity.(*model.LearningEntity)
			for _, ref := range l.EntityRefs {
				if ref == entityID || ref == entitySlug {
					icon := "ℹ"
					if l.Category == model.CatGotcha || l.Category == model.CatConstraint {
						icon = "⚠"
					}
					fmt.Printf("  %s %s: %s [%s]\n", icon, l.Category, l.Description, l.ConfLevel)
					found++
					break
				}
			}
		}

		if found == 0 {
			fmt.Printf("No learnings for '%s'\n", entitySlug)
		}
		return nil
	},
}

var learnSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search learning text",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := strings.ToLower(args[0])
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		learnings, _ := store.List(model.KindLearning)
		found := 0
		for _, ewb := range learnings {
			l := ewb.Entity.(*model.LearningEntity)
			if strings.Contains(strings.ToLower(l.Description), query) ||
				strings.Contains(strings.ToLower(l.Name), query) ||
				strings.Contains(strings.ToLower(ewb.Body), query) {
				icon := "ℹ"
				if l.Category == model.CatGotcha || l.Category == model.CatConstraint {
					icon = "⚠"
				}
				fmt.Printf("  %s %s: %s [%s]\n", icon, l.Category, l.Description, l.ConfLevel)
				fmt.Printf("    %s\n", store.FS.RelativePath(model.KindLearning, utils.Slugify(l.Name)))
				found++
			}
		}
		if found == 0 {
			fmt.Printf("No learnings matching '%s'\n", args[0])
		}
		return nil
	},
}

var learnPromoteKind string

var learnPromoteCmd = &cobra.Command{
	Use:   "promote <slug>",
	Short: "Promote learning to formal entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindLearning, args[0])
		if err != nil {
			return fmt.Errorf("learning not found: %s", args[0])
		}
		l := entity.(*model.LearningEntity)

		kind, ok := model.ParseEntityKind(learnPromoteKind)
		if !ok {
			return fmt.Errorf("unknown kind: %s", learnPromoteKind)
		}

		newEntity := model.NewEntityForKind(kind)
		nb := newEntity.GetBase()
		nb.Name = l.Name
		nb.Description = l.Description
		nb.Tags = l.Tags

		if dec, ok := newEntity.(*model.DecisionEntity); ok {
			dec.Statement = l.Description
			dec.Rationale = "Promoted from learning: " + l.Description
		}

		filePath, err := store.Create(newEntity, body)
		if err != nil {
			return err
		}

		l.PromotedTo = nb.ID
		store.Update(l, body)

		fmt.Printf("Promoted '%s' → %s\n", l.Name, kind)
		fmt.Printf("  New entity ID: %s\n", nb.ID)
		fmt.Printf("  File: %s\n", filePath)
		return nil
	},
}

var learnStaleCmd = &cobra.Command{
	Use:   "stale",
	Short: "Show learnings whose entities have changed",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		learnings, _ := store.List(model.KindLearning)
		found := 0
		for _, ewb := range learnings {
			l := ewb.Entity.(*model.LearningEntity)
			for _, ref := range l.EntityRefs {
				// Check if referenced entity still exists
				if _, _, err := store.Get(ref); err != nil {
					fmt.Printf("  ⚠ %s: referenced entity '%s' not found\n", l.Name, ref)
					found++
					break
				}
			}
		}
		if found == 0 {
			fmt.Println("All learnings are current. No stale references.")
		} else {
			fmt.Printf("\n%d learnings with stale references\n", found)
		}
		return nil
	},
}

func init() {
	rememberCmd.Flags().StringVar(&rememberCategory, "category", "gotcha", "learning category (gotcha, constraint, convention, context, dependency, performance, workaround)")
	rememberCmd.Flags().StringSliceVar(&rememberEntities, "entity", nil, "related entity slug or ID")
	rememberCmd.Flags().StringVar(&rememberConfidence, "confidence", "medium", "confidence level (high, medium, low)")
	rememberCmd.Flags().StringVar(&rememberSource, "source", "session-observation", "discovery source")
	rememberCmd.Flags().StringSliceVar(&rememberTags, "tag", nil, "tags")

	learnPromoteCmd.Flags().StringVar(&learnPromoteKind, "to", "decision", "target entity kind")
	learnPromoteCmd.MarkFlagRequired("to")

	learnCmd.AddCommand(learnListCmd, learnAboutCmd, learnSearchCmd, learnPromoteCmd, learnStaleCmd)

	rootCmd.AddCommand(rememberCmd)
	rootCmd.AddCommand(learnCmd)
}
