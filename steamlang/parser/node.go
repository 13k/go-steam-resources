package parser

import (
	"fmt"
	"strings"
)

// Node types
const (
	Root NodeType = iota
	Class
	Enum
	Property
	Type
	Literal
)

var globals = []*Node{
	&Node{Type: Type, Value: "byte"},
	&Node{Type: Type, Value: "short"},
	&Node{Type: Type, Value: "ushort"},
	&Node{Type: Type, Value: "int"},
	&Node{Type: Type, Value: "uint"},
	&Node{Type: Type, Value: "long"},
	&Node{Type: Type, Value: "ulong"},
	&Node{Type: Literal, Value: "ulong.MaxValue"},
	&Node{Type: Literal, Value: "SteamKit2.GC.Internal.CMsgProtoBufHeader"},
	&Node{Type: Literal, Value: "SteamKit2.Internal.CMsgProtoBufHeader"},
}

// NodeType represents node types
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

// Node ...
type Node struct {
	Parent            *Node
	Children          []*Node
	Type              NodeType // Class              | Enum              | Property
	Value             string   // (Name)             | (Name)            | (Name)
	Ref               *Node    // type-param (Ident) | type-param (Type) | type (Type)
	RefParam          *Node    // nil                | nil               | type-param (FlagsOpt)
	Default           []*Node  // nil                | nil               | default-value (Default)
	Qualifier         string   // nil                | nil               | type-qualifier (Flags)
	Annotation        string   // "removed"          | "flags"           | "obsolete"|"removed" (Obsolete)
	AnnotationComment string   // nil                | nil               | reason<String> (Obsolete)

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

// IsRoot returns true if the node is a root node
func (node *Node) IsRoot() bool {
	return node.Type == Root
}

// IsClass returns true if the node is a root node
func (node *Node) IsClass() bool {
	return node.Type == Class
}

// IsEnum returns true if the node is a root node
func (node *Node) IsEnum() bool {
	return node.Type == Enum
}

// IsProperty returns true if the node is a root node
func (node *Node) IsProperty() bool {
	return node.Type == Property
}

// IsType returns true if the node is a root node
func (node *Node) IsType() bool {
	return node.Type == Type
}

// IsLiteral returns true if the node is a root node
func (node *Node) IsLiteral() bool {
	return node.Type == Literal
}

func (node *Node) add(child *Node) error {
	node.Children = append(node.Children, child)
	child.Parent = node
	return node.symbols.add(child)
}

// Lookup ...
func (node *Node) Lookup(sym string) *Node {
	if node := node.symbols.lookupString(sym); node != nil {
		return node
	}

	if node.Parent != nil {
		return node.Parent.Lookup(sym)
	}

	return nil
}

func (node *Node) lookup(token *Token) (*Node, error) {
	if node := node.symbols.lookup(token); node != nil {
		return node, nil
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

// String ...
func (node *Node) String() string {
	var b strings.Builder
	b.WriteString(node.Type.String())
	b.WriteString("<")
	b.WriteString(node.Value)
	b.WriteString(">")
	return b.String()
}

// QualifiedString ...
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

// AstString ...
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
