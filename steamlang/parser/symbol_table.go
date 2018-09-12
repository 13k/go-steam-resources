package parser

import (
	"fmt"
	"strings"
)

const symbolPathSep = "::"

func symbolPath(symbol string) []string {
	return strings.Split(symbol, symbolPathSep)
}

// SymbolTable ...
type SymbolTable map[string]*Node

func (t SymbolTable) add(node *Node) error {
	if _, ok := t[node.Value]; ok {
		return fmt.Errorf("%q already defined", node.Value)
	}

	t[node.Value] = node
	return nil
}

func (t SymbolTable) lookupString(sym string) *Node {
	return t.lookup(&Token{Type: TokenIdent, Value: sym})
}

func (t SymbolTable) lookup(token *Token) *Node {
	if token == nil {
		return nil
	}

	if token.Type == TokenIdent {
		return t.lookupPath(symbolPath(token.Value))
	}

	return &Node{Type: Literal, Qualifier: token.Type.String(), Value: token.Value}
}

func (t SymbolTable) lookupPath(path []string) *Node {
	if len(path) == 0 {
		return nil
	}

	if node, ok := t[path[0]]; ok {
		if len(path) > 1 {
			return node.symbols.lookupPath(path[1:])
		}

		return node
	}

	return nil
}
