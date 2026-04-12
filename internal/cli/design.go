package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/uiml"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var (
	designType string
)

var designCmd = &cobra.Command{
	Use:   "design",
	Short: "Manage UI designs (UIML)",
}

var designCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new design entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		design := &model.DesignEntity{
			BaseEntity: model.BaseEntity{
				Kind:   model.KindDesign,
				Name:   name,
				Status: model.StatusDraft,
			},
			DesignType: model.DesignType(designType),
		}

		skeleton := `<screen name="` + name + `">
  <layout direction="vertical">
    <heading>` + name + `</heading>
    <text>Design your screen here</text>
  </layout>
</screen>
`
		filePath, err := store.Create(design, skeleton)
		if err != nil {
			return err
		}
		fmt.Printf("Created design: %s\n", name)
		fmt.Printf("  Type: %s\n", designType)
		fmt.Printf("  File: %s\n", filePath)
		return nil
	},
}

var designListCmd = &cobra.Command{
	Use:   "list",
	Short: "List designs",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		designs, err := store.List(model.KindDesign)
		if err != nil {
			return err
		}
		if len(designs) == 0 {
			fmt.Println("No designs found.")
			return nil
		}
		for _, ewb := range designs {
			d := ewb.Entity.(*model.DesignEntity)
			fmt.Printf("  %-12s %-30s %s\n", d.DesignType, d.Name, d.Status)
		}
		return nil
	},
}

var designShowCmd = &cobra.Command{
	Use:   "show <slug>",
	Short: "Render design as ASCII art",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		_, body, err := store.GetByKind(model.KindDesign, args[0])
		if err != nil {
			return fmt.Errorf("design not found: %s", args[0])
		}

		result := uiml.Parse(body)
		if len(result.Errors) > 0 {
			for _, e := range result.Errors {
				fmt.Fprintf(os.Stderr, "  WARN: %s\n", e)
			}
		}

		ascii := uiml.RenderASCII(result.Nodes, 80)
		fmt.Print(ascii)
		return nil
	},
}

var designPreviewCmd = &cobra.Command{
	Use:   "preview <slug>",
	Short: "Open design as HTML in browser",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		_, body, err := store.GetByKind(model.KindDesign, args[0])
		if err != nil {
			return fmt.Errorf("design not found: %s", args[0])
		}

		result := uiml.Parse(body)
		html := uiml.RenderHTML(result.Nodes)

		tmpFile := filepath.Join(os.TempDir(), "syde-preview-"+args[0]+".html")
		if err := os.WriteFile(tmpFile, []byte(html), 0644); err != nil {
			return err
		}

		fmt.Printf("Preview: %s\n", tmpFile)

		var openCmd *exec.Cmd
		switch runtime.GOOS {
		case "darwin":
			openCmd = exec.Command("open", tmpFile)
		case "linux":
			openCmd = exec.Command("xdg-open", tmpFile)
		case "windows":
			openCmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", tmpFile)
		}
		if openCmd != nil {
			openCmd.Start()
		}
		return nil
	},
}

var designExportFormat string

var designExportCmd = &cobra.Command{
	Use:   "export <slug>",
	Short: "Export design to a format",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		_, body, err := store.GetByKind(model.KindDesign, args[0])
		if err != nil {
			return fmt.Errorf("design not found: %s", args[0])
		}

		result := uiml.Parse(body)

		switch designExportFormat {
		case "html":
			fmt.Print(uiml.RenderHTML(result.Nodes))
		case "ascii":
			fmt.Print(uiml.RenderASCII(result.Nodes, 80))
		default:
			fmt.Print(uiml.RenderASCII(result.Nodes, 80))
		}
		return nil
	},
}

var designValidateCmd = &cobra.Command{
	Use:   "validate <slug>",
	Short: "Validate UIML syntax",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		_, body, err := store.GetByKind(model.KindDesign, args[0])
		if err != nil {
			return fmt.Errorf("design not found: %s", args[0])
		}

		errs := uiml.Validate(body)
		if len(errs) == 0 {
			fmt.Println("Validation passed. No errors.")
			return nil
		}
		for _, e := range errs {
			fmt.Printf("  line %d: %s\n", e.Line, e.Message)
		}
		fmt.Printf("\n%d errors\n", len(errs))
		return nil
	},
}

var designLinkCmd = &cobra.Command{
	Use:   "link <design-slug> <entity-slug>",
	Short: "Link design to component/flow/concept",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindDesign, args[0])
		if err != nil {
			return fmt.Errorf("design not found: %s", args[0])
		}
		d := entity.(*model.DesignEntity)

		target, _, err := store.Get(args[1])
		if err != nil {
			return fmt.Errorf("entity not found: %s", args[1])
		}
		tb := target.GetBase()

		switch tb.Kind {
		case model.KindComponent:
			d.ComponentRefs = append(d.ComponentRefs, tb.ID)
		case model.KindFlow:
			d.FlowRefs = append(d.FlowRefs, tb.ID)
		case model.KindConcept:
			d.ConceptRefs = append(d.ConceptRefs, tb.ID)
		default:
			return fmt.Errorf("can only link to component, flow, or concept (got %s)", tb.Kind)
		}

		d.Relationships = append(d.Relationships, model.Relationship{
			Target: tb.ID,
			Type:   model.RelVisualizes,
			Label:  fmt.Sprintf("Visualizes %s %s", tb.Kind, tb.Name),
		})

		store.Update(d, body)
		fmt.Printf("Linked design '%s' to %s '%s'\n", d.Name, tb.Kind, tb.Name)
		return nil
	},
}

func init() {
	_ = utils.Slugify
	designCreateCmd.Flags().StringVar(&designType, "type", "screen", "design type (screen, flow, component)")
	designExportCmd.Flags().StringVar(&designExportFormat, "format", "ascii", "export format (ascii, html)")
	designCmd.AddCommand(designCreateCmd, designListCmd, designShowCmd, designPreviewCmd, designExportCmd, designValidateCmd, designLinkCmd)
	rootCmd.AddCommand(designCmd)
}
