package parser

import (
	"fmt"
	"strings"
)

// Node types.
const (
	Root NodeType = iota
	Class
	Enum
	Property
	Type
	Literal
)

var globals = []*Node{
	{Type: Type, Value: "byte"},
	{Type: Type, Value: "short"},
	{Type: Type, Value: "ushort"},
	{Type: Type, Value: "int"},
	{Type: Type, Value: "uint"},
	{Type: Type, Value: "long"},
	{Type: Type, Value: "ulong"},
	{Type: Literal, Value: "ulong.MaxValue"},
	{Type: Literal, Value: "SteamKit2.GC.Internal.CMsgProtoBufHeader"},
	{Type: Literal, Value: "SteamKit2.Internal.CMsgProtoBufHeader"},
}

// NodeType represents node types.
type NodeType int

func (t NodeType) String() string {
	switch t {
	case Root:
		return "Root"
	case Class:
		return "Class"
	case Enum:
		return "Enum"
	case Property:
		return "Property"
	case Type:
		return "Type"
	case Literal:
		return "Literal"
	}

	return ""
}

// Node represents an AST node.
//
// It belongs to a Parent (can be nil), has multiple Children (can be empty) and has a Type and a
// Value.
//
// The meaning of the fields Ref, RefParam, Default, Qualifier and Annotation* change
// depending on the value of the Type field, according to the following table.
//
// The values in the table are meta-variables in the EBNF syntax. Values in parenthesis are the
// equivalent fields in the original SteamKit's AST nodes.
//
//	|-------------------|--------------------|-------------------|----------------------------------------|
//	| Type              | Class              | Enum              | Property                               |
//	|-------------------|--------------------|-------------------|----------------------------------------|
// 	| Ref               | class-emsg (Ident) | enum-type (Type)  | property-type (Type)                   |
// 	| RefParm           | nil                | nil               | property-type-param (FlagsOpt)         |
// 	| Default           | nil                | nil               | property-values (Default)              |
// 	| Qualifier         | ""                 | ""                | property-qualifier (Flags)             |
// 	| Annotation        | class-annotation   | enum-annotation   | property-annotation-value (Obsolete)   |
// 	| AnnotationComment | ""                 | ""                | property-annotation-comment (Obsolete) |
//	|-------------------|--------------------|-------------------|----------------------------------------|
//
type Node struct {
	Parent            *Node
	Children          []*Node
	Type              NodeType
	Value             string
	Ref               *Node
	RefParam          *Node
	Default           []*Node
	Qualifier         string
	Annotation        string
	AnnotationComment string

	symbols SymbolTable
}

func makeNode(typ NodeType, value string, parent *Node) (*Node, error) {
	node := &Node{
		Type:    typ,
		Value:   value,
		symbols: make(SymbolTable),
	}

	if parent != nil {
		if err := parent.add(node); err != nil {
			return nil, err
		}
	}

	return node, nil
}

func makeRoot() (*Node, error) {
	node, err := makeNode(Root, "", nil)

	if err != nil {
		return nil, err
	}

	for _, global := range globals {
		if err := node.symbols.add(global); err != nil {
			return nil, err
		}
	}

	return node, nil
}

// IsRoot returns true if the node is a Root node
func (node *Node) IsRoot() bool {
	return node.Type == Root
}

// IsClass returns true if the node is a Class node
func (node *Node) IsClass() bool {
	return node.Type == Class
}

// IsEnum returns true if the node is an Enum node
func (node *Node) IsEnum() bool {
	return node.Type == Enum
}

// IsProperty returns true if the node is a Property node
func (node *Node) IsProperty() bool {
	return node.Type == Property
}

// IsType returns true if the node is a Type node
func (node *Node) IsType() bool {
	return node.Type == Type
}

// IsLiteral returns true if the node is a Literal node
func (node *Node) IsLiteral() bool {
	return node.Type == Literal
}

func (node *Node) add(child *Node) error {
	node.Children = append(node.Children, child)
	child.Parent = node
	return node.symbols.add(child)
}

// Lookup returns a node referenced by the given symbol, starting from the context of the current
// node and up until the root node.
//
// It returns nil if no node is found.
func (node *Node) Lookup(sym string) *Node {
	if n := node.symbols.lookupString(sym); n != nil {
		return n
	}

	if node.Parent != nil {
		return node.Parent.Lookup(sym)
	}

	return nil
}

func (node *Node) lookup(token *Token) (*Node, error) {
	if n := node.symbols.lookup(token); n != nil {
		return n, nil
	}

	if node.Parent != nil {
		return node.Parent.lookup(token)
	}

	return nil, fmt.Errorf("%q is undefined", token.Value)
}

func (node *Node) addDefault(token *Token) error {
	defNode, err := node.lookup(token)

	if err != nil {
		return err
	}

	node.Default = append(node.Default, defNode)

	return nil
}

// String returns a formatted string of the node's type and value
func (node *Node) String() string {
	var b strings.Builder
	b.WriteString(node.Type.String())
	b.WriteString("<")
	b.WriteString(node.Value)
	b.WriteString(">")
	return b.String()
}

// QualifiedString returns a formatted string of all ascending nodes until the root, joined by "::"
func (node *Node) QualifiedString() string {
	var path []string

	for n := node; n != nil; n = n.Parent {
		if n.IsRoot() {
			break
		}

		path = append([]string{n.String()}, path...)
	}

	return strings.Join(path, "::")
}

// AstString returns a formatted AST string
func (node *Node) AstString() string {
	return node.astString(0)
}

func mapNodesToQString(nodes []*Node) []string {
	result := make([]string, len(nodes))
	for i, n := range nodes {
		result[i] = n.QualifiedString()
	}
	return result
}

func (node *Node) astString(level int) string {
	var b strings.Builder
	indent := strings.Repeat(" ", level)
	var detailsSlice []string

	if node.Ref != nil {
		detailsSlice = append(detailsSlice, node.Ref.QualifiedString())
	}

	if node.Qualifier != "" {
		detailsSlice = append(detailsSlice, node.Qualifier)
	}

	b.WriteString(indent)
	b.WriteString("(")
	b.WriteString(node.Type.String())
	b.WriteString(" ")
	b.WriteString(node.Value)

	if detailsSlice != nil {
		b.WriteString("{")
		b.WriteString(strings.Join(detailsSlice, ", "))
		b.WriteString("}")
	}

	if node.Default != nil {
		b.WriteString(" = ")
		b.WriteString(strings.Join(mapNodesToQString(node.Default), "|"))
	}

	for _, child := range node.Children {
		b.WriteString("\n")
		b.WriteString(child.astString(level + 1))
	}

	if len(node.Children) > 0 {
		b.WriteString("\n")
		b.WriteString(indent)
	}

	b.WriteString(")")
	return b.String()
}
