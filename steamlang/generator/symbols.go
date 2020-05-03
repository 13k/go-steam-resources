package generator

import (
	"fmt"

	"github.com/13k/go-steam-resources/v2/steamlang/parser"
)

var (
	symbolMapping = map[string]string{
		"byte":           "uint8",
		"short":          "int16",
		"ushort":         "uint16",
		"int":            "int32",
		"uint":           "uint32",
		"long":           "int64",
		"ulong":          "uint64",
		"ulong.MaxValue": "^uint64(0)",
		"SteamKit2.GC.Internal.CMsgProtoBufHeader": "steam.CMsgProtoBufHeader",
		"SteamKit2.Internal.CMsgProtoBufHeader":    "steam.CMsgProtoBufHeader",
	}
)

func translateSymbol(sym1 string) string {
	if sym2, ok := symbolMapping[sym1]; ok {
		return sym2
	}

	return sym1
}

func genSymbol(node *parser.Node) string {
	switch node.Type {
	case parser.Type, parser.Class, parser.Enum, parser.Literal:
		return node.Value
	case parser.Property:
		return newProp(node).Name()
	}

	panic(fmt.Errorf("Could not generate symbol for node %s<%q>", node.Type, node.Value))
}

func genValueSymbol(node *parser.Node) string {
	steamlangValue := genSymbol(node)
	return translateSymbol(steamlangValue)
}

func genTypeSymbol(node *parser.Node) string {
	var steamlangType string

	switch node.Type {
	case parser.Enum:
		steamlangType = genEnumType(node)
	case parser.Type:
		steamlangType = genSymbol(node)
	}

	return translateSymbol(steamlangType)
}

func genEnumType(enumNode *parser.Node) string {
	if enumNode.Ref != nil {
		// enum name<type> { ... }
		return genSymbol(enumNode.Ref)
	} else if enumNode.Annotation == "flags" {
		// enum name flags { ... }
		return steamLangTypeEnumFlags
	}

	return steamLangTypeEnumDefault
}
