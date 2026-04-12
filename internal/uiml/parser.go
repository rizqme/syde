package uiml

import "fmt"

// ParseError represents a parse error with line number.
type ParseError struct {
	Line    int
	Message string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("line %d: %s", e.Line, e.Message)
}

// ParseResult contains the AST and any errors.
type ParseResult struct {
	Nodes  []*Node
	Errors []ParseError
}

// Parse parses UIML source into an AST.
func Parse(source string) *ParseResult {
	p := &parser{
		lexer:  NewLexer(source),
		valid:  ValidTags(),
		result: &ParseResult{},
	}
	p.advance()
	p.result.Nodes = p.parseChildren("")
	return p.result
}

type parser struct {
	lexer   *Lexer
	current Token
	valid   map[string]NodeKind
	result  *ParseResult
}

func (p *parser) advance() {
	p.current = p.lexer.Next()
}

func (p *parser) parseChildren(parentTag string) []*Node {
	var nodes []*Node

	for {
		switch p.current.Kind {
		case TokEOF:
			return nodes

		case TokTagClose:
			if p.current.Value == parentTag {
				p.advance() // consume the closing tag
				return nodes
			}
			// Mismatched close tag
			p.result.Errors = append(p.result.Errors, ParseError{
				Line:    p.current.Line,
				Message: fmt.Sprintf("unexpected </%s>, expected </%s>", p.current.Value, parentTag),
			})
			p.advance()
			return nodes

		case TokTagOpen:
			node := p.parseElement()
			if node != nil {
				nodes = append(nodes, node)
			}

		case TokText:
			nodes = append(nodes, &Node{
				Kind: NodeTextContent,
				Text: p.current.Value,
				Line: p.current.Line,
			})
			p.advance()

		case TokComment:
			p.advance() // skip comments

		default:
			p.advance() // skip unexpected tokens
		}
	}
}

func (p *parser) parseElement() *Node {
	tagName := p.current.Value
	line := p.current.Line

	kind, ok := p.valid[tagName]
	if !ok {
		p.result.Errors = append(p.result.Errors, ParseError{
			Line:    line,
			Message: fmt.Sprintf("unknown tag: <%s>", tagName),
		})
		kind = NodeKind(tagName) // proceed anyway for error recovery
	}

	node := &Node{Kind: kind, Line: line}
	p.advance() // consume tag open

	// Parse attributes
	for {
		if p.current.Kind == TokGT || p.current.Kind == TokSelfClose || p.current.Kind == TokEOF {
			break
		}

		if p.current.Kind == TokText {
			// Boolean attribute or attribute name
			attrName := p.current.Value
			p.advance()

			if p.current.Kind == TokEquals {
				p.advance() // consume =
				if p.current.Kind == TokAttrValue {
					node.Attrs = append(node.Attrs, Attribute{Key: attrName, Value: p.current.Value})
					p.advance()
				}
			} else {
				// Boolean attribute
				node.Attrs = append(node.Attrs, Attribute{Key: attrName})
			}
			continue
		}

		if p.current.Kind == TokAttrValue {
			// Bare attribute value (shouldn't happen normally)
			node.Attrs = append(node.Attrs, Attribute{Key: p.current.Value})
			p.advance()
			continue
		}

		p.advance() // skip unexpected
	}

	// Self-closing
	if p.current.Kind == TokSelfClose {
		node.SelfClose = true
		p.advance()
		return node
	}

	if SelfClosingTags[kind] && p.current.Kind == TokGT {
		// Void element used as <tag ...> without />, still treat as self-closing
		node.SelfClose = true
		p.advance()
		return node
	}

	// Consume >
	if p.current.Kind == TokGT {
		p.advance()
	}

	// Parse children until closing tag
	node.Children = p.parseChildren(tagName)

	return node
}
