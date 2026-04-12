package uiml

// NodeKind represents the type of AST node.
type NodeKind string

const (
	// Structure
	NodeScreen  NodeKind = "screen"
	NodeLayout  NodeKind = "layout"
	NodeGrid    NodeKind = "grid"
	NodeStack   NodeKind = "stack"
	NodeNavbar  NodeKind = "navbar"
	NodeSidebar NodeKind = "sidebar"
	NodeMain    NodeKind = "main"
	NodeFooter  NodeKind = "footer"
	NodeSection NodeKind = "section"
	NodeCard    NodeKind = "card"
	NodePanel   NodeKind = "panel"
	NodeModal   NodeKind = "modal"
	NodeDrawer  NodeKind = "drawer"
	NodeTabs    NodeKind = "tabs"
	NodeTab     NodeKind = "tab"

	// Content
	NodeHeading     NodeKind = "heading"
	NodeText        NodeKind = "text"
	NodeParagraph   NodeKind = "paragraph"
	NodeLabel       NodeKind = "label"
	NodeBreadcrumb  NodeKind = "breadcrumb"
	NodeDivider     NodeKind = "divider"
	NodeSpacer      NodeKind = "spacer"

	// Interactive
	NodeButton      NodeKind = "button"
	NodeButtonGroup NodeKind = "button-group"
	NodeLink        NodeKind = "link"
	NodeInput       NodeKind = "input"
	NodeTextarea    NodeKind = "textarea"
	NodeSelect      NodeKind = "select"
	NodeOption      NodeKind = "option"
	NodeCheckbox    NodeKind = "checkbox"
	NodeRadio       NodeKind = "radio"
	NodeToggle      NodeKind = "toggle"
	NodeSlider      NodeKind = "slider"
	NodeFileUpload  NodeKind = "file-upload"
	NodeSearch      NodeKind = "search"

	// Data display
	NodeTable       NodeKind = "table"
	NodeColumns     NodeKind = "columns"
	NodeRow         NodeKind = "row"
	NodeList        NodeKind = "list"
	NodeItem        NodeKind = "item"
	NodeMetric      NodeKind = "metric"
	NodeTrend       NodeKind = "trend"
	NodeBadge       NodeKind = "badge"
	NodeTag         NodeKind = "tag"
	NodeAvatar      NodeKind = "avatar"
	NodeIcon        NodeKind = "icon"
	NodeImage       NodeKind = "image"
	NodePlaceholder NodeKind = "placeholder"
	NodeChart       NodeKind = "chart"
	NodeProgress    NodeKind = "progress"
	NodeSkeleton    NodeKind = "skeleton"

	// Navigation
	NodeLogo       NodeKind = "logo"
	NodeNav        NodeKind = "nav"
	NodeMenu       NodeKind = "menu"
	NodePagination NodeKind = "pagination"
	NodeStepper    NodeKind = "stepper"
	NodeStep       NodeKind = "step"

	// Feedback
	NodeAlert      NodeKind = "alert"
	NodeToast      NodeKind = "toast"
	NodeTooltip    NodeKind = "tooltip"
	NodeEmptyState NodeKind = "empty-state"
	NodeLoading    NodeKind = "loading"

	// Meta
	NodeNote       NodeKind = "note"
	NodeTodo       NodeKind = "todo"
	NodeVariant    NodeKind = "variant"
	NodeResponsive NodeKind = "responsive"
	NodeAt         NodeKind = "at"

	// Text content (not a tag)
	NodeTextContent NodeKind = "text-content"
)

// Attribute is a key-value attribute on a node.
type Attribute struct {
	Key   string
	Value string // empty string means boolean attribute
}

// Node represents a UIML AST node.
type Node struct {
	Kind       NodeKind
	Attrs      []Attribute
	Children   []*Node
	Text       string // for text-content and inline text
	Line       int
	SelfClose  bool
}

// Attr gets an attribute value by key. Returns empty string if not found.
func (n *Node) Attr(key string) string {
	for _, a := range n.Attrs {
		if a.Key == key {
			if a.Value == "" {
				return "true" // boolean attribute
			}
			return a.Value
		}
	}
	return ""
}

// HasAttr checks if a boolean attribute is present.
func (n *Node) HasAttr(key string) bool {
	for _, a := range n.Attrs {
		if a.Key == key {
			return true
		}
	}
	return false
}

// SelfClosingTags are tags that don't have children.
var SelfClosingTags = map[NodeKind]bool{
	NodeDivider:     true,
	NodeSpacer:      true,
	NodeInput:       true,
	NodeSlider:      true,
	NodeMetric:      true,
	NodeIcon:        true,
	NodeImage:       true,
	NodePlaceholder: true,
	NodeProgress:    true,
	NodeSkeleton:    true,
	NodeResponsive:  true,
}

// ValidTags returns all valid UIML tag names.
func ValidTags() map[string]NodeKind {
	return map[string]NodeKind{
		"screen": NodeScreen, "layout": NodeLayout, "grid": NodeGrid, "stack": NodeStack,
		"navbar": NodeNavbar, "sidebar": NodeSidebar, "main": NodeMain, "footer": NodeFooter,
		"section": NodeSection, "card": NodeCard, "panel": NodePanel, "modal": NodeModal,
		"drawer": NodeDrawer, "tabs": NodeTabs, "tab": NodeTab,
		"heading": NodeHeading, "text": NodeText, "paragraph": NodeParagraph, "label": NodeLabel,
		"breadcrumb": NodeBreadcrumb, "divider": NodeDivider, "spacer": NodeSpacer,
		"button": NodeButton, "button-group": NodeButtonGroup, "link": NodeLink,
		"input": NodeInput, "textarea": NodeTextarea, "select": NodeSelect, "option": NodeOption,
		"checkbox": NodeCheckbox, "radio": NodeRadio, "toggle": NodeToggle, "slider": NodeSlider,
		"file-upload": NodeFileUpload, "search": NodeSearch,
		"table": NodeTable, "columns": NodeColumns, "row": NodeRow, "list": NodeList, "item": NodeItem,
		"metric": NodeMetric, "trend": NodeTrend, "badge": NodeBadge, "tag": NodeTag,
		"avatar": NodeAvatar, "icon": NodeIcon, "image": NodeImage, "placeholder": NodePlaceholder,
		"chart": NodeChart, "progress": NodeProgress, "skeleton": NodeSkeleton,
		"logo": NodeLogo, "nav": NodeNav, "menu": NodeMenu, "pagination": NodePagination,
		"stepper": NodeStepper, "step": NodeStep,
		"alert": NodeAlert, "toast": NodeToast, "tooltip": NodeTooltip,
		"empty-state": NodeEmptyState, "loading": NodeLoading,
		"note": NodeNote, "todo": NodeTodo, "variant": NodeVariant,
		"responsive": NodeResponsive, "at": NodeAt,
	}
}
