package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"path"
	"strconv"

	"github.com/13k/go-steam-resources/steamlang/parser"
)

const (
	steamLangTypeEnumFlags   = "int"
	steamLangTypeEnumDefault = "int32"
)

// GenGo implements a Go code generator.
// Package can be changed anytime, but client must call `Init` again after such change.
type GenGo struct {
	// Package name for the generated code.
	// Packages are considered equal between two runs of the generator if this string is the same
	Package string
	// Base import path of the generated protobuf packages
	ProtobufPackage string

	output io.Writer
	// Imports are packages that need be imported by the generated code.
	// They are added by the generator itself when it generates code that requires these packages
	imports genImports
	// Includes are snippets of static code that are used by the generated code.
	// They are added by the generator itself when it generates code that uses code in these snippets
	includes genIncludes

	*bytes.Buffer
}

// Generate runs the code generation for the given AST root node and writes it to the given Writer.
//
// It resets the generator state, so it can be reused for multiple files and/or multiple packages.
//
// It should only be called once for each package/file combination. If called multiple times for the
// same file and package, things will break.
func (g *GenGo) Generate(w io.Writer, root *parser.Node) error {
	if !root.IsRoot() {
		return fmt.Errorf("Given node must be a root node, received %s", root)
	}

	g.init(w)

	for _, node := range root.Children {
		if err := g.writeNode(node); err != nil {
			return err
		}
	}

	return g.flush()
}

// does NOT reset includes because they must be unique per package
func (g *GenGo) init(w io.Writer) {
	g.Buffer = new(bytes.Buffer)
	g.imports = nil
	g.output = w
}

func (g *GenGo) flush() error {
	// dump everything to a buffer

	buf := new(bytes.Buffer)

	if err := g.writeHeader(buf); err != nil {
		return err
	}

	if _, err := io.Copy(buf, g); err != nil {
		return err
	}

	// TODO: parse and format the code

	// write to output

	if _, err := io.Copy(g.output, buf); err != nil {
		return err
	}

	return nil
}

func (g *GenGo) writeHeader(w io.Writer) error {
	b := bufio.NewWriter(w)
	b.WriteString("// Generated code. DO NOT EDIT\n")
	// package {{Package}}
	b.WriteString("package ")
	b.WriteString(g.Package)
	b.WriteString("\n\n")

	if len(g.imports) > 0 {
		// import ( ... )
		b.WriteString("import (\n")

		for pkg, alias := range g.imports {
			// {{Ident}}{{ImportAlias}} {{ImportPathString}}
			b.WriteByte('\t')

			if alias != "" {
				b.WriteString(alias)
				b.WriteByte(' ')
			}

			b.WriteString(strconv.Quote(pkg))
			b.WriteByte('\n')
		}

		b.WriteString(")\n\n")
	}

	if g.includes != nil {
		for name, snippet := range g.includes.get(g.Package) {
			b.WriteString("// Code included from snippet ")
			b.WriteString(strconv.Quote(name))
			b.WriteByte('\n')
			b.WriteString(snippet)
			b.WriteString("\n// ---\n\n")
		}
	}

	return b.Flush()
}

func (g *GenGo) writeNode(node *parser.Node) error {
	switch {
	case node.IsClass():
		return g.writeClass(node)
	case node.IsEnum():
		return g.writeEnum(node)
	}

	return fmt.Errorf("Unknown node %s", node)
}

func (g *GenGo) writeEnum(enum *parser.Node) error {
	typeName := genTypeSymbol(enum)
	props := newPropSlice(enum.Children)

	// type {{EnumName}} {{EnumType}}
	g.s("type ").s(enum.Value).b(' ').sn(typeName)
	g.n()
	// const ( ... )
	g.sn("const (")

	for _, prop := range props {
		if err := g.writeEnumProperty(enum, prop); err != nil {
			return err
		}
	}

	g.sn(")")
	g.n()

	return g.writeEnumStringer(enum, props)
}

func (g *GenGo) writeEnumProperty(enum *parser.Node, prop *prop) error {
	// {{Indent}}{{EnumName}}_{{PropertyName}} {{EnumName}}
	g.t().s(prop.Name()).b(' ').s(enum.Value)

	if prop.Default != nil {
		// = {{PropertyValue1}} | {{PropertyValue2}} ...
		g.s(" = ").s(prop.DefaultValuesSymbols())
	}

	if prop.Annotation != "" {
		// // {{Annotation}}
		g.s(" // ").s(prop.Annotation)

		if prop.AnnotationComment != "" {
			// : {{AnnotationComment}}
			g.s(": ").s(prop.AnnotationComment)
		}
	}

	g.n()

	return nil
}

func (g *GenGo) writeEnumStringer(enum *parser.Node, props propSlice) error {
	g.importPkg("fmt")
	g.importPkg("sort")
	g.importPkg("strings")

	enumNameMap := fmt.Sprintf("%s_name", enum.Value)
	// var {{EnumName}}_name = map[{{EnumName}}]string{
	g.s("var ").s(enumNameMap).s(" = map[").s(enum.Value).sn("]string{")

	for _, prop := range props.WithoutEnumValueDuplicates() {
		// {{EnumName}}_{{PropertyName}}: "{{EnumName}}_{{PropertyName}}",
		g.t().s(prop.Name()).s(": ").s(strconv.Quote(prop.Name())).sn(",")
	}

	g.sn("}")
	g.n()

	// func (e {{EnumName}}) String() string {
	//   if s, ok := {{EnumName}}_name[e]; ok {
	//     return s
	//   }
	//   var flags []string
	//   for k, v := range {{EnumName}}_name {
	//     if e&k != 0 {
	//       flags = append(flags, v)
	//     }
	//   }
	//   if flags == nil {
	//     return fmt.Sprintf("%d", e)
	//   }
	//   sort.Strings(flags)
	//   return strings.Join(flags, " | ")
	// }
	g.s("func (e ").s(enum.Value).sn(") String() string {")
	g.t().s("if s, ok := ").s(enumNameMap).sn("[e]; ok {")
	g.t().t().sn("return s")
	g.t().sn("}")
	g.t().sn("var flags []string")
	g.t().s("for k, v := range ").s(enumNameMap).sn(" {")
	g.t().t().sn("if e&k != 0 {")
	g.t().t().t().sn("flags = append(flags, v)")
	g.t().t().sn("}")
	g.t().sn("}")
	g.t().sn("if flags == nil {")
	g.t().t().sn("return fmt.Sprintf(\"%d\", e)")
	g.t().sn("}")
	g.t().sn("sort.Strings(flags)")
	g.t().sn("return strings.Join(flags, \" | \")")
	g.sn("}")
	g.n()

	return nil
}

func (g *GenGo) writeClass(node *parser.Node) error {
	props := newPropSlice(node.Children)
	constProps := props.Constants()
	noConstProps := props.NoConstants()
	propsWithCtor := props.WithConstructor()

	if err := g.writeClassConstants(node, constProps); err != nil {
		return err
	}

	if err := g.writeClassDefinition(node, noConstProps); err != nil {
		return err
	}

	if err := g.writeClassConstructor(node, propsWithCtor); err != nil {
		return err
	}

	if err := g.writeClassEMsg(node); err != nil {
		return err
	}

	if err := g.writeClassSerializer(node, noConstProps); err != nil {
		return err
	}

	if err := g.writeClassDeserializer(node, noConstProps); err != nil {
		return err
	}

	return nil
}

func (g *GenGo) writeClassConstants(class *parser.Node, props propSlice) error {
	if len(props) > 0 {
		g.sn("const (")

		for _, prop := range props {
			// {{ClassName}}_{{PropertyName}} {{PropertyType}} = {{PropertyValue1}} | {{PropertyValue2}} ...
			g.t().s(prop.Name()).b(' ').s(prop.Type()).s(" = ").sn(prop.DefaultValuesSymbols())
		}

		g.sn(")")
		g.n()
	}

	return nil
}

func (g *GenGo) writeClassDefinition(class *parser.Node, props propSlice) error {
	// type {{ClassName}} struct {
	g.s("type ").s(class.Value).sn(" struct {")

	for _, prop := range props {
		// {{Ident}}{{PropertyName}} {{PropertyStructFieldType}}
		g.t().s(prop.Name()).b(' ').sn(prop.StructFieldType())

		if pkg := prop.ProtobufPackage(); pkg != "" {
			importPath := path.Join(g.ProtobufPackage, pkg)
			g.importPkg(importPath)
		}
	}

	g.sn("}")
	g.n()

	return nil
}

func (g *GenGo) writeClassConstructor(class *parser.Node, props propSlice) error {
	// func New{{ClassName}}() *{{ClassName}} {
	//   return &{{ClassName}}{
	//     {{PropertyName}}: {{PropertyConstructor}},
	//     ...
	//   }
	// }
	g.s("func New").s(class.Value).s("() ").sptr(class.Value).sn(" {")
	g.t().s("return ")

	if len(props) == 0 {
		g.snew(class.Value).n()
	} else {
		g.saddr(class.Value).sn("{")

		for _, prop := range props {
			g.t().t().s(prop.Name()).s(": ").s(prop.Constructor()).sn(",")
		}

		g.t().sn("}")
	}

	g.sn("}")
	g.n()

	return nil
}

func (g *GenGo) writeClassEMsg(class *parser.Node) error {
	prop := newProp(class.Ref)

	if prop == nil || !prop.FromEnum() || prop.Parent().Value != "EMsg" {
		return nil
	}

	// func (m *{{ClassName}}) GetEMsg() EMsg {
	//   return {{PropertyName}}
	// }
	g.s("func (m ").sptr(class.Value).sn(") GetEMsg() EMsg {")
	g.t().s("return ").sn(prop.Name())
	g.sn("}")
	g.n()

	return nil
}

func (g *GenGo) writeClassSerializer(class *parser.Node, props propSlice) error {
	g.importPkg("encoding/binary")
	g.importPkg("io")
	g.importPkg("github.com/golang/protobuf/proto")
	g.include("emsgmask")
	g.include("protomask")

	// func (m *Class) Serialize(w io.Writer) (err error) {
	g.s("func (m ").sptr(class.Value).sn(") Serialize(w io.Writer) (err error) {")

	for i, prop := range props {
		dataN := fmt.Sprintf("data%d", i)

		switch prop.SerializationKind() {
		case "proto":
			lenField := prop.ProtoLengthField()
			if lenField == nil {
				return fmt.Errorf("Expected %s to have a proto length field", prop.QualifiedString())
			}

			// var dataN []byte
			// if dataN, err := proto.Marshal(m.Property); err != nil { return }
			// m.ProtoLengthField = ProtoLengthFieldType(len(dataN))
			// if _, err = w.Write(dataN); err != nil { return }
			g.t().s("var ").s(dataN).sn(" []byte")
			g.t().s("if ").s(dataN).s(", err = proto.Marshal(").sfield("m", prop.Name()).sn("); err != nil { return }")
			g.t().sfield("m", lenField.Name()).s(" = ").s(lenField.Type()).s("(").slen(dataN).sn(")")
			g.t().s("if _, err = w.Write(").s(dataN).sn("); err != nil { return }")

		case "protomask":
			// dataN := MaskProto(uint32(m.Property))
			// if err = binary.Write(w, binary.LittleEndian, data); err != nil { return }
			g.t().s(dataN).s(" := MaskProto(uint32(").sfield("m", prop.Name()).sn("))")
			g.t().s("if err = binary.Write(w, binary.LittleEndian, ").s(dataN).sn("); err != nil { return }")

		default:
			// if err = binary.Write(w, binary.LittleEndian, m.Property); err != nil { return }
			g.t().s("if err = binary.Write(w, binary.LittleEndian, ").sfield("m", prop.Name()).sn("); err != nil { return }")
		}
	}

	g.t().sn("return")
	g.sn("}")
	g.n()

	return nil
}

func (g *GenGo) writeClassDeserializer(class *parser.Node, props propSlice) error {
	g.importPkg("encoding/binary")
	g.importPkg("io")
	g.importPkg("github.com/golang/protobuf/proto")
	g.include("emsgmask")
	g.include("protomask")

	// func (m *Class) Deserialize(r io.Reader) (err error) {
	g.s("func (m ").sptr(class.Value).sn(") Deserialize(r io.Reader) (err error) {")

	for i, prop := range props {
		dataN := fmt.Sprintf("data%d", i)

		switch prop.SerializationKind() {
		case "proto":
			lenField := prop.ProtoLengthField()
			if lenField == nil {
				return fmt.Errorf("Expected %s to have a proto length field", prop.QualifiedString())
			}

			// dataN := make([]byte, m.ProtoLengthField, m.ProtoLengthField)
			// if _, err = io.ReadFull(r, dataN); err != nil { return }
			// if err = proto.Unmarshal(dataN, m.Property); err != nil { return }
			g.t().s(dataN).s(" := make([]byte, ").sfield("m", lenField.Name()).s(", ").sfield("m", lenField.Name()).sn(")")
			g.t().s("if _, err = io.ReadFull(r, ").s(dataN).sn("); err != nil { return }")
			g.t().s("if err = proto.Unmarshal(").s(dataN).s(", ").sfield("m", prop.Name()).sn("); err != nil { return }")

		case "protomask":
			cast := fmt.Sprintf("MakeEMsg(%s)", dataN)
			if prop.Type() != "EMsg" {
				cast = fmt.Sprintf("%s(%s)", prop.Type(), cast)
			}

			// var dataN uint32
			// if err = binary.Read(r, binary.LittleEndian, &dataN); err != nil { return }
			// m.Property = PropertyType(MakeEMsg(dataN))
			g.t().s("var ").s(dataN).sn(" uint32")
			g.t().s("if err = binary.Read(r, binary.LittleEndian, ").saddr(dataN).sn("); err != nil { return }")
			g.t().sfield("m", prop.Name()).s(" = ").sn(cast)

		default:
			/*
				typ := prop.DereferenceType()
				if typ == "" {
					return fmt.Errorf("Failed to dereference type of %s", prop.QualifiedShortString())
				}
				methodTyp := strcase.ToCamel(typ)

				// if m.Property, err = rwu.Read{{Type}}(r); err != nil { return }
				g.t().s("if ").sfield("m", prop.Name()).s(", err = rwu.Read").s(methodTyp).sn("(r); err != nil { return }")
			*/

			// if err = binary.Read(r, binary.LittleEndian, &m.Property); err != nil { return }
			g.t().s("if err = binary.Read(r, binary.LittleEndian, &").sfield("m", prop.Name()).sn("); err != nil { return }")
		}
	}

	g.t().sn("return")
	g.sn("}")
	g.n()

	return nil
}

func (g *GenGo) importPkg(pkg string) *GenGo {
	return g.importPkgAlias(pkg, "")
}

func (g *GenGo) importPkgAlias(pkg, alias string) *GenGo {
	if g.imports == nil {
		g.imports = make(genImports)
	}

	g.imports.add(pkg, alias)
	return g
}

func (g *GenGo) include(name string) *GenGo {
	if g.includes == nil {
		g.includes = make(genIncludes)
	}

	g.includes.add(g.Package, name)
	return g
}

func (g *GenGo) b(b byte) *GenGo {
	g.WriteByte(b)
	return g
}

func (g *GenGo) bn(b byte) *GenGo {
	g.b(b)
	return g.n()
}

func (g *GenGo) n() *GenGo {
	return g.b('\n')
}

func (g *GenGo) t() *GenGo {
	return g.b('\t')
}

func (g *GenGo) s(s string) *GenGo {
	g.WriteString(s)
	return g
}

func (g *GenGo) sn(s string) *GenGo {
	g.s(s)
	return g.n()
}

func (g *GenGo) sptr(s string) *GenGo {
	return g.b('*').s(s)
}

func (g *GenGo) saddr(s string) *GenGo {
	return g.b('&').s(s)
}

func (g *GenGo) slen(s string) *GenGo {
	return g.s("len(").s(s).b(')')
}

func (g *GenGo) snew(s string) *GenGo {
	return g.s("new(").s(s).b(')')
}

func (g *GenGo) sfield(s, f string) *GenGo {
	return g.s(s).b('.').s(f)
}
