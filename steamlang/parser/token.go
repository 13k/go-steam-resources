package parser

import (
	"fmt"
	"text/scanner"
)

// TokenType represents token types
type TokenType rune

// Token types. Equivalent to text/scanner token types, except for the tokens that are not emitted.
const (
	_               TokenType = -(iota + 1) // EOF (not emitted)
	TokenIdent                              // Identifier
	TokenInt                                // Int
	TokenFloat                              // Float
	_                                       // Char (not emitted)
	TokenString                             // String
	_                                       // RawString (not emitted)
	_                                       // Comment (not emitted)
	TokenPreprocess                         // #import directives
	TokenTerminator                         // Terminator (";")
	TokenScope                              // Scope change ("{", "}")
	TokenTypeParam                          // Type parameter delimiter ("<", ">")
	TokenOperator                           // Operator ("=", "|")
)

func (t TokenType) String() string {
	switch t {
	case TokenIdent:
		return "Ident"
	case TokenInt:
		return "Int"
	case TokenFloat:
		return "Float"
	case TokenString:
		return "String"
	case TokenPreprocess:
		return "Preprocess"
	case TokenTerminator:
		return "Terminator"
	case TokenScope:
		return "Scope"
	case TokenTypeParam:
		return "TypeParam"
	case TokenOperator:
		return "Operator"
	}

	return string(t)
}

// Token represents a code token.
//
// It contains the token type, value (the string in the source code) and a position.
type Token struct {
	Type     TokenType
	Value    string
	Position scanner.Position
}

func (t *Token) String() string {
	return fmt.Sprintf("%s<%q>", t.Type, t.Value)
}

func (t *Token) IsAny(types ...TokenType) bool {
	for _, typ := range types {
		if typ == t.Type {
			return true
		}
	}

	return false
}

func (t *Token) HasAnyValue(values ...string) bool {
	for _, v := range values {
		if v == t.Value {
			return true
		}
	}

	return false
}
