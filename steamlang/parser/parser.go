package parser

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ParseBytes ...
func ParseBytes(buf []byte, filename string) (*Node, error) {
	return Parse(bytes.NewBuffer(buf), filename)
}

// ParseString ...
func ParseString(str, filename string) (*Node, error) {
	return Parse(strings.NewReader(str), filename)
}

// ParseFile ...
func ParseFile(filename string) (*Node, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer f.Close()
	return Parse(f, filename)
}

// Parse ...
func Parse(r io.Reader, filename string) (*Node, error) {
	parser, err := NewParser(r, filename)

	if err != nil {
		return nil, err
	}

	return parser.Parse()
}

// Parser ...
type Parser struct {
	scanner   *Scanner
	root      *Node
	filename  string
	directory string
}

// NewParser ...
func NewParser(r io.Reader, filename string) (*Parser, error) {
	filename, err := filepath.Abs(filename)

	if err != nil {
		return nil, err
	}

	root, err := makeRoot()

	if err != nil {
		return nil, err
	}

	parser := &Parser{
		scanner:   NewScanner(r, filename),
		root:      root,
		filename:  filename,
		directory: filepath.Dir(filename),
	}

	return parser, nil
}

// Parse ...
func (p *Parser) Parse() (*Node, error) {
	var (
		token *Token
		err   error
	)

	token, err = p.scanner.Scan()

	for err == nil {
		switch token.Type {
		case TokenPreprocess:
			token, err = p.handlePreprocess(token)
		case TokenIdent:
			switch token.Value {
			case "class":
				token, err = p.handleClass()
			case "enum":
				token, err = p.handleEnum()
			default:
				err = fmt.Errorf("Unknown declaration %s", token)
			}
		default:
			err = fmt.Errorf("Unknown token %s", token)
		}

		if err == nil {
			token, err = p.scanner.Scan()
		}
	}

	if err != nil && err != io.EOF {
		if token != nil {
			err = fmt.Errorf("%s: %s", token.Position, err)
		}

		return nil, err
	}

	return p.root, nil
}

func (p *Parser) requireToken(typ TokenType) (t *Token, err error) {
	defer func() {
		debug("requireToken(%q) : t=%v, e=%v\n", typ, t, err)
	}()

	if t, err = p.scanner.Peek(); err != nil {
		return
	}

	if t.Type != typ {
		err = fmt.Errorf("Expected token %s<*>, got %s", typ, t)
		return
	}

	t, err = p.scanner.Scan()
	return
}

func (p *Parser) requireAnyToken(types ...TokenType) (t *Token, err error) {
	defer func() {
		debug("requireAnyToken(%q) : t=%v, e=%v\n", types, t, err)
	}()

	if t, err = p.scanner.Peek(); err != nil {
		return
	}

	if !t.IsAny(types...) {
		err = fmt.Errorf("Expected token of types %q, got %s", types, t)
		return
	}

	t, err = p.scanner.Scan()
	return
}

func (p *Parser) requireTokenValue(typ TokenType, value string) (t *Token, err error) {
	defer func() {
		debug("requireTokenValue(%q, %q) : t=%v, e=%v\n", typ, value, t, err)
	}()

	if t, err = p.scanner.Peek(); err != nil {
		return
	}

	if t.Type != typ || t.Value != value {
		err = fmt.Errorf("Expected token %s<%q>, got %s", typ, value, t)
		return
	}

	t, err = p.scanner.Scan()
	return
}

func (p *Parser) optionalToken(typ TokenType) (t *Token, err error) {
	defer func() {
		debug("optionalToken(%q) : t=%v, e=%v\n", typ, t, err)
	}()

	if t, err = p.scanner.Peek(); err != nil {
		return
	}

	if t.Type != typ {
		t = nil
		return
	}

	t, err = p.scanner.Scan()
	return
}

func (p *Parser) optionalTokenValue(typ TokenType, value string) (t *Token, err error) {
	defer func() {
		debug("optionalTokenValue(%q, %q) : t=%v, e=%v\n", typ, value, t, err)
	}()

	if t, err = p.scanner.Peek(); err != nil {
		return
	}

	if t.Type != typ || t.Value != value {
		t = nil
		return
	}

	t, err = p.scanner.Scan()
	return
}

func (p *Parser) optionalTokenAnyValue(typ TokenType, values ...string) (t *Token, err error) {
	defer func() {
		debug("optionalTokenAnyValue(%q, %q) : t=%v, e=%v\n", typ, values, t, err)
	}()

	if t, err = p.scanner.Peek(); err != nil {
		return
	}

	if t.Type != typ || !t.HasAnyValue(values...) {
		t = nil
		return
	}

	t, err = p.scanner.Scan()
	return
}

func (p *Parser) optionalTypeParamToken() (*Token, error) {
	token, err := p.optionalTokenValue(TokenTypeParam, "<")

	if err != nil {
		return token, err
	}

	if token != nil {
		token, err = p.requireAnyToken(TokenIdent, TokenInt, TokenFloat)

		if err != nil {
			return token, err
		}

		if t, err := p.requireTokenValue(TokenTypeParam, ">"); err != nil {
			return t, err
		}
	}

	return token, err
}

func (p *Parser) handlePreprocess(token *Token) (*Token, error) {
	switch token.Value {
	case "import":
		return p.handleImport()
	}

	return token, fmt.Errorf("Invalid macro #%q", token.Value)
}

func (p *Parser) handleImport() (*Token, error) {
	// current token is "#import"
	t, err := p.requireToken(TokenString)

	if err != nil {
		return t, err
	}

	filename := filepath.Join(p.directory, t.Value)
	root, err := ParseFile(filename)

	if err != nil {
		return t, err
	}

	for _, child := range root.Children {
		if err := p.root.add(child); err != nil {
			return t, err
		}
	}

	return t, nil
}

func (p *Parser) handleClass() (*Token, error) {
	// current token is "class"
	nameToken, err := p.requireToken(TokenIdent)

	if err != nil {
		return nameToken, err
	}

	node, err := makeNode(Class, nameToken.Value, p.root)

	if err != nil {
		return nameToken, err
	}

	paramToken, err := p.optionalTypeParamToken()

	if err != nil {
		return paramToken, err
	}

	if paramToken != nil {
		node.Ref, err = node.lookup(paramToken)

		if err != nil {
			return paramToken, err
		}
	}

	if annotToken, err := p.optionalTokenAnyValue(TokenIdent, "removed"); err != nil {
		return annotToken, err
	} else if annotToken != nil {
		node.Annotation = annotToken.Value
	}

	return p.parseInnerScope(node)
}

func (p *Parser) handleEnum() (*Token, error) {
	// current token is "enum"
	nameToken, err := p.requireToken(TokenIdent)

	if err != nil {
		return nameToken, err
	}

	node, err := makeNode(Enum, nameToken.Value, p.root)

	if err != nil {
		return nameToken, err
	}

	paramToken, err := p.optionalTypeParamToken()

	if err != nil {
		return paramToken, err
	}

	if paramToken != nil {
		node.Ref, err = node.lookup(paramToken)

		if err != nil {
			return paramToken, err
		}
	}

	if annotToken, err := p.optionalTokenAnyValue(TokenIdent, "flags"); err != nil {
		return annotToken, err
	} else if annotToken != nil {
		node.Annotation = annotToken.Value
	}

	return p.parseInnerScope(node)
}

func (p *Parser) parseInnerScope(node *Node) (*Token, error) {
	scopeBegin, err := p.requireTokenValue(TokenScope, "{")

	if err != nil {
		return scopeBegin, err
	}

	scopeEnd, err := p.optionalTokenValue(TokenScope, "}")

	for err == nil && scopeEnd == nil {
		if t, err := p.parseScopeProperty(node); err != nil {
			return t, err
		}

		scopeEnd, err = p.optionalTokenValue(TokenScope, "}")
	}

	if err != nil {
		return scopeEnd, err
	}

	return p.requireToken(TokenTerminator)
}

func (p *Parser) parseScopeProperty(parent *Node) (*Token, error) {
	def1, err := p.requireToken(TokenIdent)

	if err != nil {
		return def1, err
	}

	param, err := p.optionalTypeParamToken()

	if err != nil {
		return param, err
	}

	def2, err := p.optionalToken(TokenIdent)

	if err != nil {
		return def2, err
	}

	def3, err := p.optionalToken(TokenIdent)

	if err != nil {
		return def3, err
	}

	var name, typ, quali *Token

	if def3 != nil {
		name = def3
		typ = def2
		quali = def1
	} else if def2 != nil {
		name = def2
		typ = def1
	} else {
		name = def1
	}

	node, err := makeNode(Property, name.Value, parent)

	if err != nil {
		return name, err
	}

	if typ != nil {
		node.Ref, err = node.lookup(typ)

		if err != nil {
			return typ, err
		}
	}

	if param != nil {
		node.RefParam, err = node.lookup(param)

		if err != nil {
			return param, err
		}
	}

	if quali != nil {
		node.Qualifier = quali.Value
	}

	assign, err := p.optionalTokenValue(TokenOperator, "=")

	if err != nil {
		return assign, err
	}

	if assign != nil {
		for {
			defValue, err := p.scanner.Scan()

			if err != nil {
				return defValue, err
			}

			err = node.addDefault(defValue)

			if err != nil {
				return defValue, err
			}

			union, err := p.optionalTokenValue(TokenOperator, "|")

			if err != nil {
				return union, err
			}

			if union == nil {
				break
			}
		}
	}

	term, err := p.requireToken(TokenTerminator)

	if err != nil {
		return term, err
	}

	annotIdent, err := p.optionalTokenAnyValue(TokenIdent, "obsolete", "removed")

	if err != nil {
		return annotIdent, err
	}

	if annotIdent != nil {
		node.Annotation = annotIdent.Value
		annot, err := p.optionalToken(TokenString)

		if err != nil {
			return annot, err
		}

		if annot != nil {
			node.AnnotationComment = annot.Value
		}
	}

	return nil, nil
}
