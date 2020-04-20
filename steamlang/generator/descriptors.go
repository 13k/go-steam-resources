package generator

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/13k/go-steam-resources/steamlang/parser"
	"github.com/iancoleman/strcase"
)

type prop struct {
	*parser.Node

	typ       string
	typeProp  *prop
	name      string
	values    string
	styp      string
	sftyp     string
	ctor      string
	protoLen  *prop
	arraySize *int64
	enumValue *int64
}

func newProp(node *parser.Node) *prop {
	if node == nil {
		return nil
	}

	if !node.IsProperty() {
		return nil
	}

	return &prop{Node: node}
}

func (p *prop) Parent() *parser.Node {
	return p.Node.Parent
}

func (p *prop) FromEnum() bool {
	return p.Parent() != nil && p.Parent().IsEnum()
}

func (p *prop) FromClass() bool {
	return p.Parent() != nil && p.Parent().IsClass()
}

func (p *prop) IsConst() bool {
	return p.Qualifier == "const"
}

func (p *prop) IsProto() bool {
	return p.Qualifier == "proto"
}

func (p *prop) IsBoolMarshal() bool {
	return p.Qualifier == "boolmarshal"
}

func (p *prop) IsSteamIDMarshal() bool {
	return p.Qualifier == "steamidmarshal"
}

func (p *prop) IsGameIDMarshal() bool {
	return p.Qualifier == "gameidmarshal"
}

func (p *prop) IsProtoMask() bool {
	return p.Qualifier == "protomask"
}

func (p *prop) IsProtoMaskGC() bool {
	return p.Qualifier == "protomaskgc"
}

func (p *prop) IsSlice() bool {
	return strings.HasPrefix(p.Type(), "[]")
}

func (p *prop) IsArray() bool {
	return p.ArraySize() > 0
}

func (p *prop) IsMap() bool {
	return strings.HasPrefix(p.Type(), "map[")
}

func (p *prop) HasEnumType() bool {
	return p.TypeProp() != nil && p.TypeProp().FromEnum()
}

// ThisProperty.Ref -> Node.Ref -> Node.Ref -> ... -> Go type
func (p *prop) DereferenceType() string {
	var next, prev *parser.Node
	next = p.Node

	for next != nil {
		prev = next
		next = next.Ref
	}

	return genTypeSymbol(prev)
}

// TypeProp returns a prop of this prop's type (Ref)
// It will be nil if this prop has no type or the type is not a Property itself
func (p *prop) TypeProp() *prop {
	if p.typeProp == nil {
		p.typeProp = newProp(p.Ref)
	}

	return p.typeProp
}

// Type translates the property's SteamLang type to Go type
func (p *prop) Type() string {
	if p.typ == "" {
		typ := genSymbol(p.Node.Ref)

		switch {
		case p.IsBoolMarshal():
			if typ == "byte" {
				p.typ = "bool"
			} else {
				p.typ = translateSymbol(typ)
			}
		case p.IsArray():
			if typ == "byte" {
				p.typ = fmt.Sprintf("[%d]byte", p.ArraySize())
			} else {
				typ = translateSymbol(typ)
				p.typ = fmt.Sprintf("[%d]%s", p.ArraySize(), typ)
			}
		default:
			p.typ = translateSymbol(typ)
		}
	}

	return p.typ
}

// Name translates this Property's SteamLang name to a valid Go name
func (p *prop) Name() string {
	if p.name == "" {
		if p.FromEnum() || (p.FromClass() && p.IsConst()) {
			p.name = fmt.Sprintf("%s_%s", p.Parent().Value, p.Value)
		} else {
			p.name = strcase.ToCamel(p.Value)
		}
	}

	return p.name
}

// ArraySize extracts the integer in this Property's type-parameter (RefParam).
// If present, this Property is an array and it returns a positive integer. Otherwise returns 0.
func (p *prop) ArraySize() int64 {
	if p.arraySize == nil {
		p.arraySize = new(int64)

		if p.RefParam != nil {
			if size, err := strconv.ParseInt(p.RefParam.Value, 10, 64); err == nil {
				*p.arraySize = size
			}
		}
	}

	return *p.arraySize
}

// ProtoLengthField returns the Class field (as a prop) that stores the proto message length.
//
// Classes that contain a proto message also have a field that stores the proto message length.
//
// The proto Property (Qualifier == "proto") declaring the proto message field will have a
// type-param (RefParam) referencing the message length field.
//
// For example:
//
//     class MsgGCHdrProtoBuf
//     {
//     	protomaskgc uint msg = 0;
//     	int headerLength;
//
//     	proto<headerLength> SteamKit2.GC.Internal.CMsgProtoBufHeader proto;
//     };
//
// The `proto` field of type `SteamKit2.GC.Internal.CMsgProtoBufHeader` has a `proto` qualifier and
// a type-param `headerLength`. This type-param is a reference to the class field `headerLength`.
//
// When the class is serialized/deserialized, it will use the value in `headerLength` to indicate
// the size of the message stored in `proto`.
//
func (p *prop) ProtoLengthField() *prop {
	if !p.IsProto() {
		return nil
	}

	if p.protoLen == nil {
		p.protoLen = newProp(p.RefParam)
	}

	return p.protoLen
}

// DefaultValuesSymbols translates the Property's values in Default.
//
// Enum properties can have multiple values and they will be OR'ed together.
//
// Class properties will always have only one value.
func (p *prop) DefaultValuesSymbols() string {
	if p.values == "" {
		if p.Default == nil {
			return p.values
		}

		var values []string

		for _, node := range p.Default {
			values = append(values, genValueSymbol(node))
		}

		p.values = strings.Join(values, " | ")
	}

	return p.values
}

// StructFieldType returns this Property's appropriate Go type to be used in structs.
func (p *prop) StructFieldType() string {
	if p.sftyp == "" {
		if p.IsProto() {
			p.sftyp = "*" + p.Type()
		} else {
			p.sftyp = p.Type()
		}
	}

	return p.sftyp
}

// Constructor returns Go code that initializes a value of this Property.
// Properties without a constructor don't need to be initialized (their zero-value doesn't require
// explicit initialization).
func (p *prop) Constructor() string {
	if p.ctor == "" {
		switch {
		case p.IsConst(), p.IsArray():
			return ""
		case p.IsProto(): // pointer to proto struct
			p.ctor = fmt.Sprintf("&%s{}", p.Type())
		case p.IsSlice(), p.IsMap(): // slice or map
			p.ctor = fmt.Sprintf("make(%s)", p.Type())
		case p.Default == nil:
			return ""
		default: // anything with a default value
			p.ctor = p.DefaultValuesSymbols()
		}
	}

	return p.ctor
}

// SerializationKind returns the kind of serialization/deserialization required by this Property.
func (p *prop) SerializationKind() string {
	if p.styp == "" {
		switch {
		case p.IsBoolMarshal():
			p.styp = "bool"
		case p.IsSteamIDMarshal(), p.IsGameIDMarshal():
			p.styp = "uint64"
		case p.IsProtoMask(), p.IsProtoMaskGC():
			p.styp = "protomask"
		case p.IsProto():
			p.styp = "proto"
		case p.IsArray(), p.IsSlice():
			p.styp = "slice"
		case p.HasEnumType():
			p.styp = "enum"
		}
	}

	return p.styp
}

// ProtobufPackage returns the generated protobuf package this Property references.
// This is a hack for hardcoded shit in SteamLang files
func (p *prop) ProtobufPackage() string {
	if !p.IsProto() {
		return ""
	}

	i := strings.IndexByte(p.Type(), '.')

	if i < 0 {
		i = 0
	}

	pkg := p.Type()[0:i]

	switch pkg {
	case "steam", "artifact", "csgo", "dota2", "tf2", "underlords":
		return pkg
	case "steamclient":
		return "steam/client"
	case "steamworks":
		return "dota2/steamworks"
	default:
		return ""
	}
}

// EvalEnumValue returns the evaluated value of this Property's value (Default).
func (p *prop) EvalEnumValue() int64 {
	if !p.FromEnum() {
		return 0
	}

	if p.enumValue == nil {
		var result int64

		for _, node := range p.Default {
			var parsed int64

			switch {
			case node.IsLiteral():
				s := node.Value

				if i, err := strconv.ParseInt(s, 0, 64); err == nil {
					parsed = i
				} else {
					panic(fmt.Errorf("Could not parse enum value %q into an integer", s))
				}
			case node.IsProperty():
				parsed = newProp(node).EvalEnumValue()
			default:
				panic(fmt.Errorf("Cannot infer enum value for node %s", node.QualifiedString()))
			}

			result |= parsed
		}

		p.enumValue = new(int64)
		*p.enumValue = result
	}

	return *p.enumValue
}

type propSlice []*prop

func newPropSlice(nodes []*parser.Node) (props propSlice) {
	for _, node := range nodes {
		if !node.IsProperty() {
			continue
		}

		props = append(props, newProp(node))
	}

	return
}

func (s propSlice) Constants() (props propSlice) {
	for _, prop := range s {
		if prop.IsConst() {
			props = append(props, prop)
		}
	}

	return
}

func (s propSlice) NoConstants() (props propSlice) {
	for _, prop := range s {
		if !prop.IsConst() {
			props = append(props, prop)
		}
	}

	return
}

func (s propSlice) Protos() (props propSlice) {
	for _, prop := range s {
		if prop.IsProto() {
			props = append(props, prop)
		}
	}

	return
}

func (s propSlice) WithConstructor() (props propSlice) {
	for _, prop := range s {
		if prop.Constructor() != "" {
			props = append(props, prop)
		}
	}

	return
}

// groups props by evaluated enum value
func (s propSlice) byEnumValue() map[int64]propSlice {
	m := make(map[int64]propSlice)

	for _, prop := range s {
		v := prop.EvalEnumValue()
		m[v] = append(m[v], prop)
	}

	return m
}

type propSorterByAnnotation struct {
	props propSlice
}

func (s *propSorterByAnnotation) Len() int {
	return len(s.props)
}

func (s *propSorterByAnnotation) Swap(i, j int) {
	s.props[i], s.props[j] = s.props[j], s.props[i]
}

// less annotations is "better"
func (s *propSorterByAnnotation) Less(i, j int) bool {
	return s.props[i].Annotation < s.props[j].Annotation
}

// WithoutEnumValueDuplicates filters out properties that have the same evaluated enum value.
// It assumes all properties in the slice belong to an enum.
// When duplicates are detected, properties without negative annotations (removed, deprecated) are
// given higher priority.
func (s propSlice) WithoutEnumValueDuplicates() (props propSlice) {
	byValue := s.byEnumValue()

	for _, group := range byValue {
		sorter := &propSorterByAnnotation{props: group}
		// use stable sorting to keep the order of equal annotations
		sort.Stable(sorter)
	}

	selected := make(map[int64]bool)

	for _, prop := range s {
		v := prop.EvalEnumValue()

		if !selected[v] {
			selected[v] = true
			props = append(props, byValue[v][0])
		}
	}

	return
}
