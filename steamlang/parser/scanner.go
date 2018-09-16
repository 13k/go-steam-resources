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

type Scanner interface {
	Scan() (*Token, error)
	Peek() (*Token, error)
}

type scannerImpl struct {
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
func NewScanner(r io.Reader, filename string) Scanner {
	s := &scannerImpl{}
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
func NewScannerBytes(buf []byte, filename string) Scanner {
	return NewScanner(bytes.NewBuffer(buf), filename)
}

// NewScannerString ...
func NewScannerString(str string, filename string) Scanner {
	return NewScanner(strings.NewReader(str), filename)
}

func (s *scannerImpl) onError(ts *scanner.Scanner, msg string) {
	s.err = errors.New(msg)
}

func (s *scannerImpl) scan() (rune, error) {
	ch := s.Scanner.Scan()
	return ch, s.err
}

// Scan ...
func (s *scannerImpl) Scan() (token *Token, err error) {
	debug("Scan() : ")

	if s.peek != nil {
		token = s.peek
		s.peek = nil
		debug("[peek]: t=%v e=%v\n", token, err)
		return
	}

	t, err := s.scan()

	token = &Token{
		Value:    s.TokenText(),
		Position: s.Position,
	}

	if t < 0 {
		token.Type = TokenType(t)
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
func (s *scannerImpl) Peek() (*Token, error) {
	t, err := s.Scan()
	s.peek = t
	return t, err
}
