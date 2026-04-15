# UIML Wireframe Research Survey

> Research artifact for plan `uiml-wireframe-research`. Lives under
> `.syde/research/` which is outside the source-tree's view, so it
> doesn't trip orphan/sync gates. Free-form notes — not a syde entity.

## 1. Lexer state machine and the attribute bug

### What the lexer does today

`internal/uiml/lexer.go` defines a hand-written single-pass tokenizer:

```go
type Lexer struct {
    src  []rune
    pos  int
    line int
}

func (l *Lexer) Next() Token {
    l.skipWhitespacePreserveNewlines()

    if l.pos >= len(l.src) { return Token{Kind: TokEOF, …} }
    if l.match("<!--")     { return l.readComment() }
    if l.match("</")       { /* TokTagClose */ }
    if l.match("/>")       { /* TokSelfClose */ }
    if l.src[l.pos] == '<' { /* TokTagOpen */ }
    if l.src[l.pos] == '>' { return TokGT }
    if l.src[l.pos] == '=' { return TokEquals }
    if l.src[l.pos] == '"' { return l.readQuotedString() }

    return l.readTextContent()    // <-- the bug
}
```

`readTextContent` reads until the next `<`:

```go
func (l *Lexer) readTextContent() Token {
    line := l.line
    start := l.pos
    for l.pos < len(l.src) && l.src[l.pos] != '<' { … }
    val := strings.TrimSpace(string(l.src[start:l.pos]))
    if val == "" { return l.Next() }
    return Token{Kind: TokText, Value: val, Line: line}
}
```

### Trace: `<layout direction="vertical">`

Position the lexer at the start. Calling `Next()` repeatedly:

| Call | Position before | Token returned | Notes |
|------|-----------------|----------------|-------|
| 1 | `<layout direction="vertical">` | `TokTagOpen "layout"` | `<` consumed; `readTagName` reads `layout` |
| 2 | ` direction="vertical">` | `TokText "direction=\"vertical\">"` (or similar) | **Bug.** `skipWhitespacePreserveNewlines` skips the space; `d` is not `<`/`>`/`=`/`"`, so falls through to `readTextContent`, which slurps until the next `<`. The whole attribute name + `=` + quoted value + `>` becomes one TokText. |
| 3 | (depends on what's next) | `TokTagOpen "heading"` or whatever follows | The `>` and `\n` were already consumed inside the text token in step 2. |

The parser's attribute loop (`parser.go` `parseElement`) expects to see distinct
`TokText` (attribute name) + `TokEquals` + `TokAttrValue` tokens. Instead it
sees one fat TokText whose value contains everything, treats it as a boolean
attribute, then on the next iteration it sees TokTagOpen for the child element,
calls `p.advance()` to skip "unexpected", and **consumes the child's open tag**.
By the time the parser tries to recurse into `parseChildren("layout")`, the
child element is gone and the matching `</heading>` shows up before the
`<heading>` was ever opened — hence the confusing `unexpected </heading>,
expected </layout>` error.

### Minimal fix

Add an `inTag bool` to the lexer. Set true when reading an open tag's name,
clear on `>` or `/>`. When `inTag` is true, `Next()` must NOT call
`readTextContent`; instead it tokenizes an identifier (the attribute name).

```go
type Lexer struct {
    src   []rune
    pos   int
    line  int
    inTag bool   // NEW
}

func (l *Lexer) Next() Token {
    l.skipWhitespacePreserveNewlines()
    if l.pos >= len(l.src) { return Token{Kind: TokEOF, Line: l.line} }

    // Always-checked structural tokens first
    if l.match("<!--") { return l.readComment() }
    if l.match("</")   { /* unchanged; clears inTag is unnecessary */ }
    if l.match("/>")   { l.inTag = false; /* unchanged */ }

    if l.src[l.pos] == '<' {
        l.advance(1)
        name := l.readTagName()
        l.inTag = true                                 // NEW
        return Token{Kind: TokTagOpen, Value: name, Line: l.line}
    }
    if l.src[l.pos] == '>' {
        l.advance(1)
        l.inTag = false                                // NEW
        return Token{Kind: TokGT, Line: l.line}
    }
    if l.src[l.pos] == '=' { /* unchanged */ }
    if l.src[l.pos] == '"' { return l.readQuotedString() }

    if l.inTag {                                       // NEW
        // Inside an opening tag — read an attribute name (identifier).
        name := l.readTagName()  // reuses the identifier reader
        if name != "" {
            return Token{Kind: TokText, Value: name, Line: l.line}
        }
        // Unknown char inside tag — skip and recurse so we don't loop.
        l.advance(1)
        return l.Next()
    }

    return l.readTextContent()
}
```

That's the entire patch. ~10 lines. The parser's attribute loop is already
correct (see section 2) — we only need the lexer to emit cleaner tokens.

### Acceptance trace

After the fix, `<layout direction="vertical">` produces:

| # | Token | Value | inTag after |
|---|-------|-------|-------------|
| 1 | TokTagOpen | `layout` | true |
| 2 | TokText | `direction` | true |
| 3 | TokEquals | — | true |
| 4 | TokAttrValue | `vertical` | true |
| 5 | TokGT | — | false |

Which is exactly what `parseElement`'s attribute loop already handles.

---

## 2. Parser confirmation

`internal/uiml/parser.go` `parseElement` lines 88-160:

```go
node := &Node{Kind: kind, Line: line}
p.advance() // consume tag open

// Parse attributes
for {
    if p.current.Kind == TokGT || p.current.Kind == TokSelfClose || p.current.Kind == TokEOF {
        break
    }
    if p.current.Kind == TokText {
        attrName := p.current.Value
        p.advance()
        if p.current.Kind == TokEquals {
            p.advance()
            if p.current.Kind == TokAttrValue {
                node.Attrs = append(node.Attrs, Attribute{Key: attrName, Value: p.current.Value})
                p.advance()
            }
        } else {
            node.Attrs = append(node.Attrs, Attribute{Key: attrName})
        }
        continue
    }
    if p.current.Kind == TokAttrValue { … bare value, edge case … }
    p.advance() // skip unexpected
}
```

This loop is correct. It handles `name="value"` and bare boolean attributes.
The **problem is purely lexer-side** — the parser never gets the
right tokens today. With the lexer fix in section 1, this code path starts
working without modification.

`parseChildren(parentTag)` lines 44-86 is similarly correct: matches
`</parentTag>`, recurses on TokTagOpen, accumulates TokText as
NodeTextContent. No changes needed.

---

## 3. AST tag catalog

From `internal/uiml/ast.go` `ValidTags()` map. Filtered to wireframe-relevant
kinds:

### Structure
| Tag | NodeKind | Notes |
|---|---|---|
| `screen` | NodeScreen | Root container; `name` and `width` attributes |
| `layout` | NodeLayout | Generic flex container; `direction` attribute (`horizontal` / `vertical`) |
| `grid` | NodeGrid | CSS grid; `cols` and `gap` attributes |
| `stack` | NodeStack | Stacked vertical container (alias for `layout direction="vertical"`) |
| `navbar` | NodeNavbar | Top bar; flex-row by default in renderer |
| `sidebar` | NodeSidebar | Side rail; `width` attribute, vertical block in renderer |
| `main` | NodeMain | Main content region; flex-1 in renderer |
| `footer` | NodeFooter | Bottom bar |
| `section` | NodeSection | Labelled section with optional `title` |
| `card` | NodeCard | Bordered container with padding |
| `panel` | NodePanel | Generic panel |
| `modal` | NodeModal | Modal dialog |
| `drawer` | NodeDrawer | Slide-in drawer |
| `tabs` / `tab` | NodeTabs / NodeTab | Tab strip |
| `columns` | NodeColumns | (currently used as table column-header) |
| `row` | NodeRow | (currently used as table row) |

### Content / interactive (less wireframe-critical)
`heading`, `text`, `paragraph`, `label`, `breadcrumb`, `divider`, `spacer`,
`button`, `button-group`, `link`, `input`, `textarea`, `select`, `option`,
`checkbox`, `radio`, `toggle`, `slider`, `file-upload`, `search`, `table`,
`list`, `item`, `metric`, `trend`, `badge`, `tag`, `avatar`, `icon`, `image`,
`placeholder`, `chart`, `progress`, `skeleton`, `logo`, `nav`, `menu`,
`pagination`, `stepper`, `step`, `alert`, `toast`, `tooltip`, `empty-state`.

### Self-closing (no children)
`divider`, `spacer`, `input`, `slider`, `metric`, `icon`, `image`,
`placeholder`, `progress`, `skeleton`, `responsive`.

---

## 4. Per-NodeKind RenderHTML baseline

From `internal/uiml/render_html.go`. Captures what each tag emits **today**,
to be replaced by the wireframe renderer.

| Tag | Today's HTML | Honoured attrs |
|---|---|---|
| screen | `<div class="mx-auto border border-border rounded-lg overflow-hidden">` + optional name header bar | `name`, `width` |
| navbar | `<nav class="flex items-center gap-4 px-4 py-3 border-b border-border">` | (none) |
| sidebar | `<aside class="border-r border-border p-3">` | `width` (inline style) |
| layout | `<div class="flex flex-col min-h-0 flex-1">` (or `flex-row` if `direction="horizontal"`) | `direction` |
| main | `<main class="flex-1 p-6 overflow-y-auto">` | (none) |
| grid | `<div class="grid grid-cols-{N}" style="gap:{Npx}">` | `cols`, `gap` |
| card | `<div class="bg-card border border-border rounded-lg p-4">` | (none) |
| section | `<section class="mt-6">` + optional `<h3>` title | `title` |
| heading | `<h2 class="text-xl font-semibold">` (or `h3` for level=2/3) | `level` |
| text / paragraph | `<p class="text-sm text-muted">` | (none) |
| button | `<button class="px-3 py-1.5 rounded-md border bg-card hover:bg-…">` | `variant` (primary/outline/danger) |
| input | `<input type="…" placeholder="…">` styled | `type`, `placeholder` |
| list / item | (no list wrapper today; item renders as `<a class="px-3 py-1.5 rounded-md hover:…">`) | `active` |
| menu | `<nav class="flex flex-col gap-1">` | (none) |
| metric | `<div><div class="text-xs text-muted">{label}</div><div class="text-2xl">{value}</div></div>` | `label`, `value` |
| badge | `<span class="inline-flex px-2 py-0.5 rounded-full">` | `color` |
| divider | `<hr class="border-border" />` | (none) |
| spacer | `<div class="flex-1"></div>` | (none) |

### Critical observations

1. **Everything uses real Tailwind UI styles** (`bg-card`, `hover:`, brand
   colours). It looks like a low-fidelity rendering of the actual app, not a
   wireframe. There's no visual cue saying "this is a sketch, not a real UI."
2. **Region boundaries are invisible** without dashed borders or labels —
   the user can't tell which region is the sidebar and which is the main panel.
3. **Layout depends on attributes**. Without the lexer fix from section 1,
   `<layout direction="horizontal">` doesn't parse, so every wireframe stacks
   vertically. That's the immediate cause of "stacked headings" in the
   screenshot the user reported.
4. **`columns` and `row` are misused** — they currently render as `<thead>`
   and `<tr>` table rows, which only makes sense inside `<table>`. They are
   NOT general-purpose horizontal containers. Don't reach for them.
5. **No region labels.** Even when layout works, the renderer never prints the
   tag name (e.g. "sidebar", "main") inside the region, so wireframes look
   anonymous.

The wireframe renderer has to: (a) use sketch styles (dashed borders, mono font,
greyscale), (b) print tag-name labels in the corner of every structural region,
(c) replace content tags (heading, text, button) with placeholder bars instead
of styled real text, (d) respect the layout attributes that the lexer fix
enables.

---

*Sections 5-7 (visual language exploration, tag mapping, follow-up plan) are
filled in by phases 3-5 of this research plan.*

---

## 5. Visual language reference (user-supplied)

The user supplied a reference image showing the exact wireframe style they
want. This pre-empts phase 3's "explore 4-5 candidates" approach — we already
know the target. Phase 3 becomes "build to this spec, validate with the
screenshot loop, iterate until matching".

### Style spec

- **Palette**: two tones only.
  - Background: light grey, ~`#e5e5e5` (slightly off-white).
  - Strokes / labels: dark charcoal grey, ~`#4a4a4a`.
  - No brand colours, no gradients, no shadows.
- **Strokes**: solid (NOT dashed), ~3-4px weight, rounded line caps.
- **Outer container**: thick rounded-rectangle border framing the whole
  screen. Looks like a printed wireframe sheet.
- **Region dividers**: thick horizontal and vertical bars that cut the
  outer container into regions. Each region is its own rectangle.
- **Text labels**: bold UPPERCASE sans-serif in the same charcoal grey
  ("LOGO", "HEADLINE", "HOURS", "LOCATION", "MENU", "CONTACT").
  Acts as section header inside the region.
- **Body text**: rendered as 2-3 horizontal grey lines (line placeholders)
  — never real prose. ~2px thick, ~60% width, varying lengths.
- **Image / media placeholders**: rectangles with a bold ✕ through them
  (two diagonals from corner to corner). The classic wireframe image stub.
- **Buttons**: thick-bordered rounded rectangles, centered UPPERCASE label
  inside, no fill.
- **Icons**: simple line-art inside a small bordered box (e.g. location
  pin inside a square).
- **List items with thumbnails**: small ✕-rectangle next to the label +
  optional secondary value (PRICE etc.) on the right, separated by a
  horizontal line.
- **Font**: heavy sans-serif (system-ui or Inter Bold). Never serif.
  Never script / handwritten.

### What this means for the renderer

The new `RenderWireframeHTML` function (implementation plan) needs to map
each NodeKind to one of these primitives:

- Structural (`screen`, `layout`, `grid`, `sidebar`, `main`, `panel`,
  `card`, `navbar`, `footer`, `section`) → bordered region with
  optional UPPERCASE label in the top-left corner.
- `heading` → bold UPPERCASE label, larger.
- `text` / `paragraph` → 2-3 horizontal grey line placeholders.
- `button` → thick-bordered rounded rect with centered UPPERCASE label.
- `image` / `placeholder` → ✕-through rectangle.
- `icon` → small bordered square with a glyph inside.
- `list` → vertical stack of items, each with optional ✕-thumbnail +
  label + horizontal line under it.
- `item` → one row of the above pattern.
- `divider` → thick horizontal bar.

`<screen>`, `<layout direction="horizontal">`, and `<grid cols="N">` need
the lexer fix from section 1 to actually produce horizontal flow. With the
fix in place, the wireframe naturally looks like the reference.

### Phase 3 scope (revised)

- Build ONE sandbox HTML (`/tmp/wireframe-sandbox.html`) matching the
  spec above. The scene is the same components-inbox layout that
  appeared in the user's earlier broken screenshot.
- Screenshot via `scripts/wireframe-shot.sh` (point Chrome at the local
  file via `file://` URL).
- Compare to the user's reference. If divergent on any axis above,
  iterate the CSS until matching.
- Document the final CSS (border weights, exact greys, font stack,
  spacing) in this section under "Final CSS spec".

Phase 4 then translates that CSS into per-NodeKind Go template strings
for `RenderWireframeHTML`.

### Second reference (mobile app wireframes)

The user supplied a second reference: a sheet of mobile-app wireframes
(login screen, category list with side drawer, product grid, product
detail with rating, order list, order tracking with stepper). These
refine the first reference:

- **Strokes are finer** (~1-2px) than the first reference (~3-4px).
  The first reference is a printed "marketing wireframe" style; the
  second is closer to typical UX-tool low-fi (Whimsical, Balsamiq Mockups
  drawn with the cleaner stencil).
- **Background is pure white**, not light grey.
- **Real text is used for labels** ("Username", "Password", "Categories",
  "Electronics", "Model Name", "Rs. 999"). Text isn't always UPPERCASE.
  Section headers ("Categories", "Details") are bold or slightly heavier.
- **Body prose is omitted** entirely — there's no long lorem text. Lines
  in the wireframe always carry meaningful labels.
- **X-rect image placeholders** are everywhere: product thumbnails,
  category cards, profile photos, even hero images. Same primitive as
  reference 1.
- **Buttons** are bordered rounded-rects with centered text label, no
  fill (e.g. "Login", "SignUp", "View all"). Same as reference 1.
- **Form inputs** are just labels above thin horizontal underlines
  (no full bordered box).
- **Search bars** are full-width rounded-rect borders with "Search"
  placeholder text inside.
- **Lists** are vertical stacks of (X-rect thumbnail + name + secondary
  text + maybe price + maybe ★★★☆☆ rating) rows separated by light
  horizontal dividers.
- **Side drawers** show a list of categories with an active state
  rendered as a grey fill behind the selected row.
- **Stepper / progress indicators** use circles connected by lines plus
  a label per step.
- **Top nav bars** are horizontal rows with `← back`, title, then a
  cluster of icons on the right (search, cart). Icons are simple line
  art.

### Synthesized spec (final)

Combining both references, the wireframe renderer should produce:

- **Palette**: charcoal `#3f3f46` strokes on white `#ffffff` background.
  Optional very-light-grey `#f4f4f5` fills for active states or panel
  backgrounds. No other colours.
- **Stroke weight**: 2px solid for region borders and dividers, 1.5px
  solid for X-placeholders and underlines, 1px solid for thin sub-lines.
  Rounded line caps (`stroke-linecap: round` on SVG; `border-radius`
  for rectangles).
- **Outer container**: 2px border, ~12px corner radius.
- **Region dividers**: 2px solid lines splitting the container into
  rectangles. Each region's optional label sits at top-left in bold.
- **Labels**: real text in heavy sans-serif (`system-ui, -apple-system,
  Inter, sans-serif`), `font-weight: 600` for region headers, `500` for
  item names. Color `#3f3f46`. Mix of cases — UPPERCASE only for top-level
  region badges; sentence/title case for actual content labels.
- **Body text placeholders**: when a `<text>` or `<paragraph>` is empty,
  render 2-3 horizontal grey lines (~2px thick, varying widths 60-90%).
  When it has content, render the content directly in a smaller grey
  font (`#71717a`, ~12px).
- **Image / placeholder primitive**: a bordered rectangle with two
  diagonal lines forming an `✕`. Aspect ratio defaults to square; the
  `<image>` and `<placeholder>` tags get this treatment.
- **Button primitive**: bordered rounded-rect (8px radius), centered
  label text in heavier weight, no fill. Hover/active states ignored.
- **List primitive**: each `<item>` is a row containing (optional small
  ✕-thumbnail) + label + (optional right-aligned secondary text) with a
  thin horizontal divider underneath.
- **Icon primitive**: small bordered square containing a tag-name letter
  or simple glyph (✓, ✕, →, ⚙, 🔍 — but rendered as line art via Unicode
  fallback or inline SVG).
- **Stepper primitive** (`<stepper>` / `<step>`): vertical chain of
  bordered circles connected by short vertical lines, each with a label
  to the right.
- **Active state**: where a `<item active>` or similar is set, fill the
  row background with `#f4f4f5` and bold the label.

### What changes vs. section 5 (first reference only)

| Aspect | First ref | Second ref | Final spec |
|---|---|---|---|
| Background | light grey `#e5e5e5` | white `#ffffff` | white `#ffffff` |
| Stroke weight | 3-4px | 1-2px | 2px regions, 1.5px placeholders, 1px sub-lines |
| Label case | all UPPERCASE | mixed (title/sentence) | UPPERCASE for region badges only; title/sentence for content |
| Body text | line placeholders only | real text where present, omitted otherwise | content if present, line placeholders if empty |
| Outer corner radius | 8-12px | 8-12px | 12px |
| Form inputs | not shown | label + underline | label + underline |
| Lists | not heavily used | central pattern | central pattern with ✕-thumbnail + label + secondary |

Phase 3 builds the sandbox to THIS final spec, not the first-reference-only
spec from earlier in this section.

---

## 6. Final CSS spec (validated via screenshot loop)

Built `/tmp/wireframe-sandbox.html` and screenshotted at 1440x900 via
the helper from phase 1. Visual result matches both user references on
every axis except a minor cosmetic bleed where the active-row fill
extends past the outer rounded corner — fixed in the implementation
plan by clipping with `overflow: hidden` on the column.

Sandbox screenshot: `/tmp/wf-sandbox.png` (44 KB, 1440x900).

### CSS variables (translate to Go template constants)

```css
--stroke:        #3f3f46;   /* charcoal — every border, label, glyph */
--stroke-light:  #a1a1aa;   /* secondary text + thin dividers */
--bg:            #ffffff;   /* page + region fill */
--fill-active:   #f4f4f5;   /* selected item row background */
--w-region:      2px;       /* outer border + region splitters */
--w-thin:        1.5px;     /* placeholders, button outlines, X-cross */
--w-hair:        1px;       /* item dividers */
--radius:        12px;      /* outer container */
--radius-btn:    8px;       /* button corners */
--radius-thumb:  4px;       /* small thumbnail / region badge */
--font:          system-ui, -apple-system, "Inter", sans-serif;
```

### Per-primitive rules

#### Outer screen container (`<screen>`)
```css
border: 2px solid #3f3f46;
border-radius: 12px;
overflow: hidden;
background: #ffffff;
display: flex;          /* if direction=horizontal — row */
                        /* default direction is column */
```

#### Region (`<sidebar>`, `<main>`, `<panel>`, `<card>`, `<section>`)
```css
border-right: 2px solid #3f3f46;   /* (or border-bottom for vertical) */
padding: 16px;
display: flex;
flex-direction: column;
gap: 12px;
overflow: hidden;       /* clip active-row fill */
```

Each region gets a small UPPERCASE label chip in the top-left corner
showing the tag name (`SIDEBAR`, `LIST`, `MAIN`, `DETAIL`). The label
is rendered as:

```css
.wf-region-label {
  font: 700 10px/1.2 system-ui;
  text-transform: uppercase;
  letter-spacing: 0.06em;
  border: 1.5px solid #3f3f46;
  border-radius: 4px;
  padding: 3px 6px;
  align-self: flex-start;
}
```

#### Heading (`<heading>`)
```css
font: 700 16px/1.2 system-ui;
color: #3f3f46;
margin: 0;
```

#### Section title (sub-header inside a region)
```css
font: 700 11px/1.2 system-ui;
text-transform: uppercase;
letter-spacing: 0.06em;
color: #a1a1aa;
margin: 0;
```

#### Body text (`<text>` / `<paragraph>`)
- If the node has content: render as small grey text (`font: 400 11px/1.3
  system-ui; color: #a1a1aa`).
- If empty: render 2-3 horizontal grey bars with varying widths:

```html
<div class="wf-text-lines">
  <span></span><span></span><span></span>
</div>
```
```css
.wf-text-lines { display: flex; flex-direction: column; gap: 6px; }
.wf-text-lines > span { height: 2px; background: #a1a1aa; border-radius: 2px; }
.wf-text-lines > span:nth-child(1) { width: 80%; }
.wf-text-lines > span:nth-child(2) { width: 95%; }
.wf-text-lines > span:nth-child(3) { width: 60%; }
```

#### Image / placeholder (`<image>`, `<placeholder>`)
```css
.wf-placeholder {
  border: 1.5px solid #3f3f46;
  border-radius: 6px;
  aspect-ratio: 16 / 9;
  position: relative;
}
/* two diagonal lines forming the X */
.wf-placeholder::before, .wf-placeholder::after {
  content: ""; position: absolute; top: 50%; left: 0; right: 0;
  height: 1.5px; background: #3f3f46;
}
.wf-placeholder::before { transform: translateY(-50%) rotate(28deg); }
.wf-placeholder::after  { transform: translateY(-50%) rotate(-28deg); }
```

#### Button (`<button>`)
```css
border: 1.5px solid #3f3f46;
border-radius: 8px;
padding: 8px 14px;
font: 600 12px/1 system-ui;
color: #3f3f46;
background: transparent;
display: inline-block;
```

#### List + item (`<list>` / `<item>`)
```css
.wf-item {
  display: flex; align-items: center; gap: 10px;
  padding: 10px 0;
  border-bottom: 1px solid #a1a1aa;
}
.wf-item.active {
  background: #f4f4f5;
  font-weight: 700;
  margin: 0 -16px;       /* full-bleed */
  padding-left: 16px;
  padding-right: 16px;
}
.wf-item-thumb {
  width: 32px; height: 32px;
  border: 1.5px solid #3f3f46;
  border-radius: 4px;
  position: relative;
}
.wf-item-thumb::before, .wf-item-thumb::after { /* X cross */ }
.wf-item-label { flex: 1; font: 500 13px/1.2 system-ui; color: #3f3f46; }
.wf-item-meta  { font: 400 11px/1.2 system-ui; color: #a1a1aa; }
```

`<item active>` gets the `.active` styling. Bare `<item>` is the regular
row. `<list>` is just a vertical stack with no own styling — it's the
parent for items.

#### Form input (`<input>`)
```css
display: flex;
flex-direction: column;
gap: 4px;
/* label above, thin underline below — no full bordered box */
```
- Label = `font: 500 12px system-ui; color: #3f3f46`.
- Underline = `height: 22px; border-bottom: 1.5px solid #3f3f46`.

#### Menu item / nav row (`<item>` inside `<menu>` or `<sidebar>`)
```css
.wf-menu-item {
  font: 500 13px/1.2 system-ui;
  color: #3f3f46;
  padding: 8px 6px;
  border-radius: 4px;
}
.wf-menu-item.active {
  background: #f4f4f5;
  font-weight: 700;
}
```

#### Divider (`<divider>`)
```css
height: 0;
border-top: 2px solid #3f3f46;
margin: 4px 0;
```

#### Stack / Layout containers (`<stack>`, `<layout>`)
```css
display: flex;
flex-direction: column;  /* or row for layout direction="horizontal" */
gap: 8px;
```

#### Grid (`<grid cols="N">`)
```css
display: grid;
grid-template-columns: repeat(N, 1fr);
gap: 16px;
```

### What's NOT styled (yet)

The sandbox doesn't cover these tags — the implementation plan should
add minimal renderers for them so screen contracts can use them:

- `<icon>` — small bordered square with a glyph inside
- `<stepper>` / `<step>` — vertical chain of bordered circles + labels
- `<navbar>` — horizontal flex row with logo + items + icons
- `<footer>` — horizontal flex row, top border, low-emphasis text
- `<badge>` — small rounded-rect with text, no fill
- `<tabs>` / `<tab>` — horizontal row with bottom border, active tab has heavier bottom border
- `<modal>` — ignored (out of scope; use `<panel>` instead)
- `<table>` — render as a list of bordered rows, no hover

### Open issues

- **Active-row clip bleed**: when `<item active>` sits inside a column
  with rounded corners, the full-bleed background can extend past the
  rounded outer container. Fix: outer container needs `overflow: hidden`
  AND every region needs `overflow: hidden` too. Already in the spec.
- **Custom region label**: should the label show the literal tag name
  (`SIDEBAR`, `MAIN`) or accept an override via attribute (`label="Kinds"`
  or `name="Kinds"`)? Recommended: support `name` attribute, fall back
  to the tag name. Document in the implementation plan.
- **Self-closing void tags** (`<image/>`, `<input/>`, `<placeholder/>`)
  need to render WITHOUT children. The renderer's existing self-closing
  detection handles this.

---

## 7. Tag → wireframe mapping table

This is the lookup the implementation plan's `RenderWireframeHTML`
switch statement needs. Every NodeKind that should produce wireframe
output is listed; tags not in this table are passed through with
`renderChildren()` only (no styling).

| Tag | Wireframe DOM | Honoured attributes | Label chip? |
|---|---|---|---|
| `screen` | `<div class="wf-screen">` outer container, flex direction from attr | `name`, `direction`, `width` | yes (top-left, "SCREEN" or `name`) |
| `layout` | `<div class="wf-layout">` flex container | `direction` (default vertical), `gap` | no |
| `grid` | `<div class="wf-grid">` CSS grid | `cols` (required), `gap` | no |
| `stack` | `<div class="wf-stack">` flex column | `gap` | no |
| `sidebar` | `<aside class="wf-region wf-sidebar">` | `width` (default 200), `name` | yes ("SIDEBAR") |
| `main` | `<main class="wf-region wf-main">` flex 1 | `name` | yes ("MAIN") |
| `panel` | `<section class="wf-region wf-panel">` | `width`, `name` | yes ("PANEL") |
| `card` | `<div class="wf-region wf-card">` | `name` | yes ("CARD") |
| `section` | `<section class="wf-section">` heading + content | `title` | yes if `title` |
| `navbar` | `<nav class="wf-navbar">` flex row top border | (none) | yes ("NAV") |
| `footer` | `<footer class="wf-footer">` flex row top border | (none) | yes ("FOOTER") |
| `heading` | `<h2 class="wf-heading">` text | `level` (1/2/3 → font size) | no |
| `text` / `paragraph` | text content if non-empty, else 3 line bars | (none) | no |
| `label` | inline grey label | (none) | no |
| `button` | `<button class="wf-button">` bordered rounded | `variant` (ignored — always outline) | no |
| `button-group` | `<div class="wf-row">` | (none) | no |
| `input` | `<div class="wf-field">` label + thin underline | `placeholder`, `type` | no |
| `search` | `<div class="wf-field wf-field-bordered">` rounded full-width box | `placeholder` | no |
| `list` | `<div class="wf-list">` vertical stack | (none) | no |
| `item` | `<div class="wf-item">` thumb + label + meta + bottom border | `active` (background fill) | no |
| `menu` | `<nav class="wf-menu">` flex column | (none) | no |
| `image` | `<div class="wf-placeholder">` X-rect | `aspect`, `width`, `height` | no |
| `placeholder` | same as `image` | (none) | no |
| `icon` | `<div class="wf-icon">` small bordered square + glyph char | `glyph` (single char) | no |
| `divider` | `<hr class="wf-divider">` | (none) | no |
| `spacer` | `<div class="wf-spacer">` flex-grow filler | (none) | no |
| `metric` | label above + bold value below, no border | `label`, `value` | no |
| `badge` | small bordered rounded chip | `color` (ignored — monochrome) | no |
| `tabs` | `<div class="wf-tabs">` row with bottom border | (none) | no |
| `tab` | `<div class="wf-tab">` bordered top, active has bottom-border heavier | `active` | no |
| `stepper` | `<div class="wf-stepper">` vertical chain | (none) | no |
| `step` | bordered circle + label row | `active`, `done` | no |
| `breadcrumb` | inline row of items separated by `›` | (none) | no |
| `progress` | thin horizontal bar with filled portion | `value` (0-100) | no |
| `table` | rendered as a styled list (rows = items) | (none) | no |

### Worked UIML examples

These three examples cover the patterns used by the 13 backfilled screen
contracts. The implementation plan's re-backfill task uses these as the
canonical templates.

#### Example 1 — Overview screen (grid of metric cards)

```xml
<screen name="Overview" direction="vertical">
  <navbar>
    <heading>syde Overview</heading>
  </navbar>
  <main>
    <grid cols="4">
      <card name="Systems"><metric label="systems" value="3"/></card>
      <card name="Components"><metric label="components" value="22"/></card>
      <card name="Contracts"><metric label="contracts" value="82"/></card>
      <card name="Concepts"><metric label="concepts" value="11"/></card>
    </grid>
    <section title="Recent activity">
      <list>
        <item><label>storage-engine updated</label></item>
        <item><label>web-spa file added</label></item>
        <item><label>concept Order created</label></item>
      </list>
    </section>
  </main>
</screen>
```

Renders as: outer rounded container, top nav row with the heading,
4-column grid of bordered metric cards, then a labelled "Recent
activity" section with a vertical list of plain-label items.

#### Example 2 — Components inbox (sidebar + list + detail)

```xml
<screen name="Components Inbox" direction="horizontal">
  <sidebar name="Kinds" width="200">
    <menu>
      <item>Systems</item>
      <item active="true">Components</item>
      <item>Contracts</item>
      <item>Concepts</item>
      <item>Flows</item>
      <item>Decisions</item>
    </menu>
  </sidebar>
  <panel name="List" width="360">
    <heading>Components</heading>
    <button-group>
      <button>Filter</button>
      <button>Sort</button>
    </button-group>
    <list>
      <item active="true"><image/><label>CLI Commands</label><label>42 files</label></item>
      <item><image/><label>Storage Engine</label><label>8 files</label></item>
      <item><image/><label>Query Engine</label><label>5 files</label></item>
      <item><image/><label>Audit Engine</label><label>8 files</label></item>
      <item><image/><label>HTTP API</label><label>3 files</label></item>
    </list>
  </panel>
  <main name="Detail">
    <heading>CLI Commands</heading>
    <text/>
    <section title="Files"><text/></section>
    <section title="Relationships">
      <list>
        <item><image/><label>syde CLI</label><label>belongs_to</label></item>
        <item><image/><label>Storage Engine</label><label>depends_on</label></item>
      </list>
    </section>
    <button-group>
      <button>Edit</button>
      <button>Delete</button>
    </button-group>
  </main>
</screen>
```

This is the exact scene rendered in the sandbox screenshot. Everything
above the Edit/Delete buttons is structural; the empty `<text/>` tags
become 3-line-bar placeholders.

#### Example 3 — ERD canvas screen

```xml
<screen name="Concepts ERD" direction="vertical">
  <navbar>
    <heading>Concepts</heading>
    <button>List</button>
    <button>ERD</button>
  </navbar>
  <main>
    <placeholder aspect="16/9"/>
  </main>
</screen>
```

Single big X-rect placeholder for the React Flow canvas, with the toggle
buttons in the top nav. ~5 lines of UIML.

### Missing primitives (proposed additions)

The following tags would help wireframes but don't currently exist in
the AST. Implementation plan can either add them or document workarounds:

- `<region>` — a generic rectangle with optional label. Right now you
  have to pick a semantic tag (sidebar, main, panel, card) even if the
  region doesn't match any. **Workaround**: use `<panel>`.
- `<canvas>` — large central placeholder for graph/erd views.
  **Workaround**: use `<placeholder aspect="16/9"/>` (proposed in
  example 3).
- `<chip>` — small bordered rounded label (badges already cover this,
  rename them in docs).
- `<thumbnail>` — the small ✕-rect that appears in list items.
  **Workaround**: use `<image/>` inside `<item>` — the implementation
  plan's `<item>` renderer detects and shrinks the first `<image>` child.
- `<rating>` — star rating. **Workaround**: use `<text>★★★☆☆</text>`.
- `<step>` connectors — the vertical lines between stepper circles need
  to be drawn. **Workaround**: render via CSS `::before` on each step.

None of these are blockers — the existing tag set is sufficient for the
13 backfilled screens.

---

## 8. Implementation plan (next session)

Run this shell block to spin up the implementation plan from this
research. Every task carries `--affected-entity` and `--affected-file`
so `task done` clears drift automatically. The plan has 5 phases, one
task each.

```bash
set -e

syde plan create "uiml-wireframe-render" \
  --background "Research plan uiml-wireframe-research produced .syde/research/uiml-survey.md with the lexer fix, visual language spec, per-tag mapping, and 3 worked UIML examples. This plan executes the implementation: fix the lexer, add a wireframe HTML renderer matching the user's reference style, rewire FormatJSON for screen contracts, re-backfill the 13 existing screen contracts, and clean up the skill docs that warn about attribute-free UIML." \
  --objective "Screen contract detail panels in the dashboard render UIML wireframes that look like classic mid-fidelity wireframes (charcoal-on-white, X-rect placeholders, bordered button labels, region badges) instead of the current stacked headings. The UIML lexer accepts attributes correctly so syde design create skeletons stop producing parse warnings." \
  --scope "In-scope: (1) lexer inTag flag fix (~10 lines in internal/uiml/lexer.go); (2) new uiml.RenderWireframeHTML function in internal/uiml/render_wireframe.go matching the CSS spec in .syde/research/uiml-survey.md section 6; (3) FormatJSON contract case calls RenderWireframeHTML for contract_kind=screen instead of RenderHTML; (4) re-backfill all 13 screen contracts using the worked examples in section 7; (5) drop the attribute-free caveat from skill/SKILL.md, skill/references/entity-spec.md, skill/references/commands.md. Out-of-scope: changing the existing RenderHTML used by design entities, adding new UIML tags from the missing primitives list, syde wireframe shot CLI command, lexer test suite (smoke tests via syde sync check are enough)."

syde plan add-phase uiml-wireframe-render --name "Lexer inTag fix" \
  --description "Add inTag bool flag to internal/uiml/lexer.go so attributes parse correctly" \
  --objective "<layout direction=\"horizontal\"> and <grid cols=\"3\"> parse without errors and the existing syde design create skeleton stops producing warnings" \
  --changes "internal/uiml/lexer.go: add inTag field on Lexer struct; in Next(), set true after reading TokTagOpen, clear on TokGT/TokSelfClose; when inTag is true and the current char isn't a structural token, read an identifier (reusing readTagName) and emit it as TokText; never call readTextContent in the inTag branch." \
  --details "See section 1 of .syde/research/uiml-survey.md for the exact diff. ~10 lines. Parser is unchanged — the existing parseElement attribute loop already handles distinct identifier/equals/value tokens correctly." \
  --notes "Smoke test: run syde design create Smoke and verify syde design show smoke produces zero parse warnings."

syde task create "Add inTag flag and identifier branch to Lexer.Next" --plan uiml-wireframe-render --phase phase_1 \
  --objective "Lexer tokenizes attributes correctly when inside an opening tag" \
  --details "internal/uiml/lexer.go: add inTag bool to Lexer struct. In Next(): on TokTagOpen branch set inTag=true; on TokGT branch clear inTag=false; on TokSelfClose branch clear inTag=false. Add a new branch before readTextContent: if inTag is true, read an identifier via readTagName and return as TokText. The existing TokEquals and readQuotedString branches already work — they fire before the readTextContent fallthrough." \
  --acceptance "syde design create Foo && syde design show foo produces zero parse warnings. Creating a screen contract with --wireframe '<screen direction=\"horizontal\"><sidebar width=\"200\"><heading>Sidebar</heading></sidebar><main><heading>Main</heading></main></screen>' passes syde sync check." \
  --affected-entity uiml-parser \
  --affected-file internal/uiml/lexer.go

syde plan add-phase uiml-wireframe-render --name "RenderWireframeHTML new function" \
  --description "internal/uiml/render_wireframe.go implementing the visual language spec from research section 6" \
  --objective "uiml.RenderWireframeHTML(nodes []*Node) returns a self-contained HTML document matching the user-reference wireframe style" \
  --changes "internal/uiml/render_wireframe.go (NEW). Mirrors render_html.go's structure but emits the CSS classes from research section 6 and the per-NodeKind rules from section 7. Uses an inline <style> block with the wf- class palette; no Tailwind dependency." \
  --details "Self-contained HTML doc. <style> block declares all wf- classes from section 6. Switch on NodeKind in renderNodeWireframe (mirrors renderNodeHTML but maps to wf- classes). Region tags (screen/sidebar/main/panel/card/section/navbar/footer) get a top-left UPPERCASE label chip showing the tag name (or name attribute if present). Empty <text>/<paragraph> nodes render as 3 line bars; non-empty render the text content. <image>/<placeholder> render as the X-rect primitive. <button> as bordered rounded label. <list>/<item> as the inbox row pattern with optional <image> first-child detected as a thumbnail. <input> as label + thin underline. Active items get the wf-active fill." \
  --notes "RenderHTML stays untouched — design entities continue to use the realistic Tailwind preview. Only screen contracts use the new wireframe renderer."

syde task create "Implement RenderWireframeHTML matching research section 6 spec" --plan uiml-wireframe-render --phase phase_2 \
  --objective "internal/uiml/render_wireframe.go exists and produces the wireframe HTML matching the sandbox screenshot at /tmp/wf-sandbox.png" \
  --details "New file internal/uiml/render_wireframe.go. Function: func RenderWireframeHTML(nodes []*Node) string. Self-contained <html><head><style>...wf classes...</style></head><body>...rendered nodes...</body></html>. Per-NodeKind switch handling screen, layout, grid, stack, sidebar, main, panel, card, section, navbar, footer, heading, text, paragraph, label, button, button-group, input, search, list, item, menu, image, placeholder, icon, divider, spacer, metric, badge, tabs, tab, stepper, step, breadcrumb, progress, table. Empty text nodes get the 3-bar placeholder. Items with an <image> child get the small ✕-thumb pattern. Active items get the fill. See section 6 for exact CSS values and section 7 for the per-tag DOM template." \
  --acceptance "Calling uiml.RenderWireframeHTML on the example 2 UIML from research section 7 produces HTML that, when screenshotted, matches /tmp/wf-sandbox.png within obvious tolerances (region badges visible, ✕ thumbnails, line placeholders, active row fill, charcoal-on-white)." \
  --affected-entity uiml-parser \
  --affected-file internal/uiml/render_wireframe.go

syde plan add-phase uiml-wireframe-render --name "FormatJSON rewire" \
  --description "FormatJSON contract case calls RenderWireframeHTML for screen contracts" \
  --objective "Dashboard contract detail panel renders screen wireframes via the new wireframe renderer instead of RenderHTML" \
  --changes "internal/query/formatter.go FormatJSON case *model.ContractEntity: when ContractKind=='screen', set entityMap['wireframe_html'] = uiml.RenderWireframeHTML(parsedNodes) instead of uiml.RenderHTML(parsedNodes). wireframe_ascii unchanged. wireframe (raw source) unchanged." \
  --details "One-line change. The existing dashboard EntityDetail.tsx contract case already mounts wireframe_html via dangerouslySetInnerHTML — it doesn't care which renderer produced it." \
  --notes "Verify by screenshotting any screen contract via scripts/wireframe-shot.sh after the change."

syde task create "Switch screen contracts to RenderWireframeHTML in FormatJSON" --plan uiml-wireframe-render --phase phase_3 \
  --objective "Dashboard /api/<proj>/entity/<slug> for any screen contract returns wireframe_html generated by RenderWireframeHTML" \
  --details "internal/query/formatter.go: replace uiml.RenderHTML(res.Nodes) with uiml.RenderWireframeHTML(res.Nodes) inside the if e.ContractKind == \"screen\" block. Leave wireframe_ascii alone (still uses RenderASCII). Leave non-screen contract paths and design entity paths alone." \
  --acceptance "scripts/wireframe-shot.sh components-inbox-screen-c5jh produces a PNG that matches /tmp/wf-sandbox.png within obvious tolerances." \
  --affected-entity query-engine \
  --affected-file internal/query/formatter.go

syde plan add-phase uiml-wireframe-render --name "Re-backfill 13 screen contracts" \
  --description "Rewrite every screen contract --wireframe using attribute-rich UIML matching the 3 worked examples" \
  --objective "Every screen contract in the project carries a wireframe that renders correctly under RenderWireframeHTML and looks like the scene it represents" \
  --changes "Bash script /tmp/syde-rebackfill-screens.sh: for each of the 13 screen contracts (overview-screen, file-tree-screen, graph-screen, plan-view-screen, task-board-screen, learning-feed-screen, concepts-erd-screen, systems-inbox-screen, components-inbox-screen, contracts-inbox-screen, concepts-inbox-screen, flows-inbox-screen, decisions-inbox-screen), syde update <slug> --wireframe '<new UIML>' using the patterns from research section 7." \
  --details "Pattern reuse: Overview uses example 1 (grid of metric cards). Inbox screens (Systems/Components/Contracts/Concepts/Flows/Decisions) use example 2 (sidebar + list + detail). Plan View / Task Board / Learning Feed / File Tree / Graph each get a tailored UIML based on their actual layout. ERD screen uses example 3 (placeholder canvas with toggle nav). Each --wireframe should be inline, single-line shell-quoted." \
  --acceptance "syde sync check passes with zero wireframe parse errors. scripts/wireframe-shot.sh on each of the 13 contracts produces a PNG whose layout matches the contract's intent (sidebar+list+detail for inbox screens, grid for overview, single canvas for graph/erd)." \
  --notes "This is the rare bulk-update case where a shell script is appropriate per the SKILL.md guidance."

syde task create "Re-backfill 13 screen contracts with attribute-rich UIML" --plan uiml-wireframe-render --phase phase_4 \
  --objective "All 13 screen contracts have wireframes that render correctly under the new renderer" \
  --details "Write /tmp/syde-rebackfill-screens.sh containing 13 syde update calls. Each updates the --wireframe field to the appropriate pattern from research section 7. Run the script. Verify with syde sync check + a few spot screenshots." \
  --acceptance "syde sync check passes; spot-screenshotting at least Components Inbox + Overview + Concepts ERD shows wireframes that match their respective patterns." \
  --affected-entity web-spa

syde plan add-phase uiml-wireframe-render --name "Skill docs cleanup" \
  --description "Drop the attribute-free UIML caveat from SKILL.md, entity-spec.md, commands.md" \
  --objective "Skill docs no longer warn agents to avoid UIML attributes" \
  --changes "skill/SKILL.md screen contract section: remove the 'stick to attribute-free structural tags' paragraph; replace with a pointer to the UIML tag vocabulary (research section 7 mapping table can be inlined or summarized). skill/references/entity-spec.md and skill/references/commands.md: same treatment." \
  --details "Search for 'attribute-free' in skill/ — three occurrences. Remove them and replace with positive guidance: 'Use <layout direction=\"horizontal|vertical\">, <grid cols=\"N\">, and the structural region tags. The wireframe renderer adds region badges automatically.' Add a brief table of the most common tags." \
  --notes "After this lands, future agents authoring wireframes get the full UIML vocabulary instead of being told to avoid attributes."

syde task create "Drop attribute-free UIML caveat from skill docs" --plan uiml-wireframe-render --phase phase_5 \
  --objective "Skill docs guide agents to use UIML attributes freely now that the lexer fix is shipped" \
  --details "skill/SKILL.md screen contract subsection: remove the 'stick to attribute-free structural tags' bullet and the surrounding caveat paragraph. Replace with positive guidance pointing at the new wireframe vocabulary (sidebar/main/panel/card/grid/layout). skill/references/entity-spec.md wireframe row: drop the 'avoid tag attributes' clause. skill/references/commands.md --wireframe flag: same. Add one short paragraph or mini table summarising the most common tags and their effect." \
  --acceptance "grep -ri 'attribute-free' skill/ returns zero hits. Reading the updated screen contract section produces a copy-pasteable example using <layout direction=\"horizontal\">." \
  --affected-entity skill-installer \
  --affected-file skill/SKILL.md \
  --affected-file skill/references/entity-spec.md \
  --affected-file skill/references/commands.md
```

That block is runnable as-is in a future session. After it completes,
run `syde plan approve uiml-wireframe-render` then walk through phase 1
→ 5 with `syde task start` / `syde task done` per the existing workflow.
