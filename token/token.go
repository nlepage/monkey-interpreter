package token

import (
	"fmt"

	"github.com/nlepage/monkey-interpreter/common"
)

// Type represents the type of token
type Type byte

func (t Type) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	}
	panic(fmt.Sprintf("Unknown token type %d", t))
}

// Token types
const (
	// Specials
	ILLEGAL Type = iota
	EOF

	// Identifiers and literals
	IDENT
	INT

	// Operators
	ASSIGN
	PLUS

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
)

// Token contains information about a token
type Token struct {
	Type    Type
	Literal string
	common.Position
}
