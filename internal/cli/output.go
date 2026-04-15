package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputFormat values accepted by --format on read commands.
const (
	FormatRich    = "rich"
	FormatCompact = "compact"
	FormatJSON    = "json"
)

// Envelope is the shared schema every --format json payload uses. kind
// identifies the command family (list, query, health, tree, ...) so
// consumers can dispatch on a single field. data holds the command-
// specific body. Meta is optional metadata (counts, timing).
type Envelope struct {
	Kind string      `json:"kind"`
	Data interface{} `json:"data"`
	Meta *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Count int `json:"count,omitempty"`
}

// Emit writes payload to stdout in the requested format. When format is
// FormatJSON it serializes an Envelope; otherwise it delegates to the
// provided rich callback (which owns all the pretty-printing the
// command does today). If the rich callback is nil, the payload is
// dumped via Go's default %+v formatter — fine for compact one-liners
// but commands with a proper rich view should always supply a callback.
func Emit(kind, format string, data interface{}, meta *Meta, rich func()) error {
	switch format {
	case FormatJSON:
		env := Envelope{Kind: kind, Data: data, Meta: meta}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(env)
	default:
		if rich != nil {
			rich()
			return nil
		}
		fmt.Printf("%+v\n", data)
		return nil
	}
}
