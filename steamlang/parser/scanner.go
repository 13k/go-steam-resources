package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

// TokenType represents token types
type TokenType rune

// Token types
// Equivalent to text/scanner token types, except tokens not emitted
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
		return "TokenOperator"
	}

	return string(t)
}

// Token ...
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

// Scanner ...
type Scanner struct {
	scanner.Scanner
	err  error
	peek *Token
}

func isIdent(ch rune, i int) bool {
	return unicode.IsLetter(ch) ||
		(unicode.IsDigit(ch) && i > 0) ||
		(ch == '_' && i > 0) ||
		(ch == ':' && i > 0) ||
		(ch == '.' && i > 0)
}

// NewScanner ...
func NewScanner(r io.Reader, filename string) *Scanner {
	s := &Scanner{}
	s.Init(r)
	s.Error = s.onError
	s.Filename = filename
	s.IsIdentRune = isIdent
	s.Mode = scanner.ScanIdents |
		scanner.ScanFloats |
		scanner.ScanStrings |
		scanner.ScanComments |
		scanner.SkipComments

	return s
}

// NewScannerBytes ...
func NewScannerBytes(buf []byte, filename string) *Scanner {
	return NewScanner(bytes.NewBuffer(buf), filename)
}

// NewScannerString ...
func NewScannerString(str string, filename string) *Scanner {
	return NewScanner(strings.NewReader(str), filename)
}

func (s *Scanner) onError(ts *scanner.Scanner, msg string) {
	s.err = errors.New(msg)
}

func (s *Scanner) scan() (rune, error) {
	ch := s.Scanner.Scan()
	return ch, s.err
}

// Scan ...
func (s *Scanner) Scan() (token *Token, err error) {
	debug("Scan() : ")

	if s.peek != nil {
		peek := s.peek
		s.peek = nil
		debug("[peek]: t=%v e=%v\n", peek, err)
		return peek, nil
	}

	var t rune

	t, err = s.scan()
	token = &Token{
		Type:     TokenType(t),
		Value:    s.TokenText(),
		Position: s.Position,
	}

	if err != nil {
		debug("t=%v e=%v\n", token, err)
		return
	}

	switch t {
	case scanner.EOF:
		err = io.EOF

	case scanner.Ident, scanner.Int, scanner.Float:

	case scanner.String:
		token.Value = strings.Trim(token.Value, `"`)

	case '#':
		var t2 rune
		t2, err = s.scan()

		if err != nil {
			debug("t2=%v e=%v\n", token, err)
			return
		}

		if t2 != scanner.Ident {
			err = errors.New("Invalid char '#'")
			debug("t2=%v e=%v\n", token, err)
			return
		}

		token.Type = TokenPreprocess
		token.Value = s.TokenText()

	case '-':
		var t2 rune
		t2, err = s.scan()

		if err != nil {
			debug("t2=%v e=%v\n", token, err)
			return
		}

		if t2 != scanner.Int && t2 != scanner.Float {
			err = errors.New("Invalid char '-'")
			debug("t2=%v e=%v\n", token, err)
			return
		}

		token.Type = TokenType(t2)
		token.Value += s.TokenText()

	case ';':
		token.Type = TokenTerminator

	case '{', '}':
		token.Type = TokenScope

	case '<', '>':
		token.Type = TokenTypeParam

	case '=', '|':
		token.Type = TokenOperator

	default:
		err = fmt.Errorf("Invalid token %s", s.TokenText())
	}

	debug("t=%v e=%v\n", token, err)
	return
}

// Peek ...
func (s *Scanner) Peek() (*Token, error) {
	t, err := s.Scan()
	s.peek = t
	return t, err
}
