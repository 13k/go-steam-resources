package parser

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FileSet is a map of filenames to io.Reader. It represents a set of files that are available to be
// parsed.
type FileSet map[string]io.Reader

func (s FileSet) importFile(base, filename string) (io.Reader, string, error) {
	if r, ok := s[filename]; ok {
		return r, filename, nil
	}

	if !filepath.IsAbs(filename) {
		filename = filepath.Join(filepath.Dir(base), filename)
	}

	f, err := os.Open(filename)

	if err != nil {
		return nil, filename, err
	}

	s[filename] = f
	return f, filename, nil
}

// ParseFile parses the file named filename within the FileSet fset. It returns the root Node of the
// file's AST.
//
// Any "#import" directives in the file's source can trigger loading of the file in disk and modify
// fset.
func ParseFile(fset FileSet, filename string) (*Node, error) {
	p, err := newParser(fset, filename)

	if err != nil {
		return nil, err
	}

	return p.parse()
}

type parser struct {
	fset    FileSet
	base    string
	root    *Node
	scanner Scanner
}

func newParser(fset FileSet, filename string) (*parser, error) {
	reader, ok := fset[filename]

	if !ok {
		return nil, fmt.Errorf("File %q not found in FileSet", filename)
	}

	scanner := NewScanner(reader, filename)
	root, err := makeRoot()

	if err != nil {
		return nil, err
	}

	p := &parser{
		fset:    fset,
		base:    filename,
		root:    root,
		scanner: scanner,
	}

	return p, nil
}

func (p *parser) parse() (*Node, error) {
	var (
		token *Token
		err   error
	)

	token, err = p.scanner.Scan()

	for err == nil {
		switch token.Type {
		case TokenPreprocess:
			token, err = p.handlePreprocess(p.root, token)
		case TokenIdent:
			switch token.Value {
			case "public":
				// ignore
			case "class":
				token, err = p.handleClass(p.root)
			case "enum":
				token, err = p.handleEnum(p.root)
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

func (p *parser) requireToken(typ TokenType) (t *Token, err error) {
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

func (p *parser) requireAnyToken(types ...TokenType) (t *Token, err error) {
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

func (p *parser) requireTokenValue(typ TokenType, value string) (t *Token, err error) {
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

func (p *parser) optionalToken(typ TokenType) (t *Token, err error) {
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

func (p *parser) optionalTokenValue(typ TokenType, value string) (t *Token, err error) {
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

func (p *parser) optionalTokenAnyValue(typ TokenType, values ...string) (t *Token, err error) {
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

func (p *parser) optionalTypeParamToken() (*Token, error) {
	token, err := p.optionalTokenValue(TokenTypeParam, "<")

	if err != nil {
		return token, err
	}

	if token != nil {
		token, err = p.requireAnyToken(TokenIdent, TokenInt, TokenFloat)

		if err != nil {
			return token, err
		}

		if t, tErr := p.requireTokenValue(TokenTypeParam, ">"); tErr != nil {
			return t, tErr
		}
	}

	return token, err
}

func (p *parser) handlePreprocess(root *Node, token *Token) (*Token, error) {
	switch token.Value {
	case "import":
		return p.handleImport(root)
	}

	return token, fmt.Errorf("Invalid macro #%q", token.Value)
}

func (p *parser) handleImport(root *Node) (*Token, error) {
	// current token is "#import"
	t, err := p.requireToken(TokenString)

	if err != nil {
		return t, err
	}

	_, filename, err := p.fset.importFile(p.base, t.Value)

	if err != nil {
		return t, err
	}

	subRoot, err := ParseFile(p.fset, filename)

	if err != nil {
		return t, err
	}

	for _, child := range subRoot.Children {
		if err := root.add(child); err != nil {
			return t, err
		}
	}

	return t, nil
}

func (p *parser) handleClass(parent *Node) (*Token, error) {
	// current token is "class"
	nameToken, err := p.requireToken(TokenIdent)

	if err != nil {
		return nameToken, err
	}

	node, err := makeNode(Class, nameToken.Value, nil)

	if err != nil {
		return nameToken, err
	}

	paramToken, err := p.optionalTypeParamToken()

	if err != nil {
		return paramToken, err
	}

	annotToken, err := p.optionalTokenAnyValue(TokenIdent, "removed")

	if err != nil {
		return annotToken, err
	}

	if annotToken != nil {
		node.Annotation = annotToken.Value
	}

	removed := node.Annotation == "removed"

	if !removed {
		if err = parent.add(node); err != nil {
			return nameToken, err
		}

		if paramToken != nil {
			node.Ref, err = node.lookup(paramToken)

			if err != nil {
				return paramToken, err
			}
		}
	}

	return p.parseInnerScope(node, !removed)
}

func (p *parser) handleEnum(parent *Node) (*Token, error) {
	// current token is "enum"
	nameToken, err := p.requireToken(TokenIdent)

	if err != nil {
		return nameToken, err
	}

	node, err := makeNode(Enum, nameToken.Value, nil)

	if err != nil {
		return nameToken, err
	}

	paramToken, err := p.optionalTypeParamToken()

	if err != nil {
		return paramToken, err
	}

	annotToken, err := p.optionalTokenAnyValue(TokenIdent, "flags", "removed")

	if err != nil {
		return annotToken, err
	}

	if annotToken != nil {
		node.Annotation = annotToken.Value
	}

	removed := node.Annotation == "removed"

	if !removed {
		if err = parent.add(node); err != nil {
			return nameToken, err
		}

		if paramToken != nil {
			node.Ref, err = node.lookup(paramToken)

			if err != nil {
				return paramToken, err
			}
		}
	}

	return p.parseInnerScope(node, !removed)
}

func (p *parser) parseInnerScope(parent *Node, lookup bool) (*Token, error) {
	scopeBegin, err := p.requireTokenValue(TokenScope, "{")

	if err != nil {
		return scopeBegin, err
	}

	scopeEnd, err := p.optionalTokenValue(TokenScope, "}")

	for err == nil && scopeEnd == nil {
		if t, tErr := p.parseProperty(parent, lookup); tErr != nil {
			return t, tErr
		}

		scopeEnd, err = p.optionalTokenValue(TokenScope, "}")
	}

	if err != nil {
		return scopeEnd, err
	}

	return p.requireToken(TokenTerminator)
}

func (p *parser) parseProperty(parent *Node, lookup bool) (*Token, error) {
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

	node, err := makeNode(Property, name.Value, nil)

	if err != nil {
		return name, err
	}

	if quali != nil {
		node.Qualifier = quali.Value
	}

	assign, err := p.optionalTokenValue(TokenOperator, "=")

	if err != nil {
		return assign, err
	}

	var defValues []*Token

	if assign != nil {
		for {
			defValue, dErr := p.scanner.Scan()

			if dErr != nil {
				return defValue, dErr
			}

			defValues = append(defValues, defValue)

			union, uErr := p.optionalTokenValue(TokenOperator, "|")

			if uErr != nil {
				return union, uErr
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

	removed := node.Annotation == "removed"

	if lookup && !removed {
		if err := parent.add(node); err != nil {
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

		for _, defValue := range defValues {
			if addErr := node.addDefault(defValue); addErr != nil {
				return defValue, addErr
			}
		}
	}

	return nil, nil
}
