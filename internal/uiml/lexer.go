package uiml

import (
	"strings"
	"unicode"
)

// TokenKind represents a lexer token type.
type TokenKind int

const (
	TokEOF TokenKind = iota
	TokTagOpen      // <tag
	TokTagClose     // </tag>
	TokSelfClose    // />
	TokGT           // >
	TokAttrName
	TokAttrValue
	TokEquals       // =
	TokText         // text content between tags
	TokComment      // <!-- ... -->
)

// Token is a lexer token.
type Token struct {
	Kind  TokenKind
	Value string
	Line  int
}

// Lexer tokenizes UIML source.
type Lexer struct {
	src   []rune
	pos   int
	line  int
	// inTag is true while the lexer is between an opening tag's name
	// and its closing `>` or `/>`. In that window we tokenize attribute
	// names as identifiers (TokText) instead of falling through to
	// readTextContent (which slurps everything until the next `<` and
	// destroys attribute structure). Cleared on TokGT and TokSelfClose.
	inTag bool
}

// NewLexer creates a lexer for the given source.
func NewLexer(src string) *Lexer {
	return &Lexer{src: []rune(src), pos: 0, line: 1}
}

// Next returns the next token.
func (l *Lexer) Next() Token {
	l.skipWhitespacePreserveNewlines()

	if l.pos >= len(l.src) {
		return Token{Kind: TokEOF, Line: l.line}
	}

	// Comment
	if l.match("<!--") {
		return l.readComment()
	}

	// Closing tag
	if l.match("</") {
		l.advance(2)
		name := l.readTagName()
		l.skipWhitespace()
		if l.pos < len(l.src) && l.src[l.pos] == '>' {
			l.advance(1)
		}
		l.inTag = false
		return Token{Kind: TokTagClose, Value: name, Line: l.line}
	}

	// Self-closing />
	if l.match("/>") {
		l.advance(2)
		l.inTag = false
		return Token{Kind: TokSelfClose, Line: l.line}
	}

	// Opening tag
	if l.src[l.pos] == '<' {
		l.advance(1)
		name := l.readTagName()
		l.inTag = true
		return Token{Kind: TokTagOpen, Value: name, Line: l.line}
	}

	// >
	if l.src[l.pos] == '>' {
		l.advance(1)
		l.inTag = false
		return Token{Kind: TokGT, Line: l.line}
	}

	// =
	if l.src[l.pos] == '=' {
		l.advance(1)
		return Token{Kind: TokEquals, Line: l.line}
	}

	// Quoted string (attribute value)
	if l.src[l.pos] == '"' {
		return l.readQuotedString()
	}

	// Inside an opening tag — read an identifier as an attribute name.
	// readTextContent would otherwise slurp everything up to the next
	// `<`, which collapses `name="value">\n  ` into one TokText that
	// the parser cannot decompose. The parser's attribute loop expects
	// a clean stream of TokText (name) + TokEquals + TokAttrValue.
	if l.inTag {
		name := l.readTagName()
		if name != "" {
			return Token{Kind: TokText, Value: name, Line: l.line}
		}
		// Unknown character inside a tag — skip it and recurse so we
		// don't loop forever.
		l.advance(1)
		return l.Next()
	}

	// Text content (between tags)
	return l.readTextContent()
}

func (l *Lexer) readComment() Token {
	line := l.line
	l.advance(4) // skip <!--
	start := l.pos
	for l.pos < len(l.src) {
		if l.match("-->") {
			val := string(l.src[start:l.pos])
			l.advance(3)
			return Token{Kind: TokComment, Value: strings.TrimSpace(val), Line: line}
		}
		if l.src[l.pos] == '\n' {
			l.line++
		}
		l.pos++
	}
	return Token{Kind: TokComment, Value: string(l.src[start:]), Line: line}
}

func (l *Lexer) readTagName() string {
	start := l.pos
	for l.pos < len(l.src) {
		ch := l.src[l.pos]
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '-' || ch == '_' {
			l.pos++
		} else {
			break
		}
	}
	return string(l.src[start:l.pos])
}

func (l *Lexer) readQuotedString() Token {
	line := l.line
	l.advance(1) // skip opening "
	start := l.pos
	for l.pos < len(l.src) && l.src[l.pos] != '"' {
		if l.src[l.pos] == '\n' {
			l.line++
		}
		l.pos++
	}
	val := string(l.src[start:l.pos])
	if l.pos < len(l.src) {
		l.advance(1) // skip closing "
	}
	return Token{Kind: TokAttrValue, Value: val, Line: line}
}

func (l *Lexer) readTextContent() Token {
	line := l.line
	start := l.pos
	for l.pos < len(l.src) && l.src[l.pos] != '<' {
		if l.src[l.pos] == '\n' {
			l.line++
		}
		l.pos++
	}
	val := strings.TrimSpace(string(l.src[start:l.pos]))
	if val == "" {
		return l.Next()
	}
	return Token{Kind: TokText, Value: val, Line: line}
}

func (l *Lexer) match(s string) bool {
	runes := []rune(s)
	if l.pos+len(runes) > len(l.src) {
		return false
	}
	for i, r := range runes {
		if l.src[l.pos+i] != r {
			return false
		}
	}
	return true
}

func (l *Lexer) advance(n int) {
	for i := 0; i < n && l.pos < len(l.src); i++ {
		if l.src[l.pos] == '\n' {
			l.line++
		}
		l.pos++
	}
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.src) && (l.src[l.pos] == ' ' || l.src[l.pos] == '\t') {
		l.pos++
	}
}

func (l *Lexer) skipWhitespacePreserveNewlines() {
	for l.pos < len(l.src) && unicode.IsSpace(l.src[l.pos]) {
		if l.src[l.pos] == '\n' {
			l.line++
		}
		l.pos++
	}
}
