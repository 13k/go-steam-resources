package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	enumSrc = `
		enum Enum<byte> flags {
			Field = 0;
			Flag1 = 1;
			Flag2 = 2;
			Flag3 = 4;
			NoField = -1; removed "For pressing ceremonial reasons"
			Flag4 = 6;
		};
	`

	classSrc = `
		#import "enum.steamd"

		// This is not a real SteamLanguage class
		class Data<Enum::Field> {
			uint a = 1;
			int z; removed "For pressing ceremonial reasons"
			ulong b = -2; obsolete "Not used"
			Enum c = Enum::Flag1 | Enum::Flag2;
			proto<a> SteamKit2.GC.Internal.CMsgProtoBufHeader d;
		};
	`
)

func TestParseFile(t *testing.T) {
	fset := FileSet{
		"class.steamd": strings.NewReader(classSrc),
		"enum.steamd":  strings.NewReader(enumSrc),
	}

	root, err := ParseFile(fset, "class.steamd")
	assert.NoError(t, err)
	assert.Equal(t, Root, root.Type)

	require.Len(t, root.Children, 2)

	// Enum

	enum := root.Children[0]
	assert.Exactly(t, root, enum.Parent)
	require.Len(t, enum.Children, 5)
	assert.Equal(t, Enum, enum.Type)
	assert.Equal(t, "Enum", enum.Value)
	require.NotNil(t, enum.Ref)
	assert.Equal(t, Type, enum.Ref.Type)
	assert.Equal(t, "byte", enum.Ref.Value)
	assert.Nil(t, enum.RefParam)
	assert.Nil(t, enum.Default)
	assert.Equal(t, "", enum.Qualifier)
	assert.Equal(t, "flags", enum.Annotation)
	assert.Equal(t, "", enum.AnnotationComment)

	// Enum::Field

	prop := enum.Children[0]
	assert.Exactly(t, enum, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "Field", prop.Value)
	assert.Nil(t, prop.Ref)
	assert.Nil(t, prop.RefParam)
	require.Len(t, prop.Default, 1)
	assert.Equal(t, "", prop.Qualifier)
	assert.Equal(t, "", prop.Annotation)
	assert.Equal(t, "", prop.AnnotationComment)

	propDefault := prop.Default[0]
	assert.Equal(t, Literal, propDefault.Type)
	assert.Equal(t, "0", propDefault.Value)

	// Enum::Flag1

	prop = enum.Children[1]
	assert.Exactly(t, enum, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "Flag1", prop.Value)
	assert.Nil(t, prop.Ref)
	assert.Nil(t, prop.RefParam)
	require.Len(t, prop.Default, 1)
	assert.Equal(t, "", prop.Qualifier)
	assert.Equal(t, "", prop.Annotation)
	assert.Equal(t, "", prop.AnnotationComment)

	propDefault = prop.Default[0]
	assert.Equal(t, Literal, propDefault.Type)
	assert.Equal(t, "1", propDefault.Value)

	// Enum::Flag2

	prop = enum.Children[2]
	assert.Exactly(t, enum, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "Flag2", prop.Value)
	assert.Nil(t, prop.Ref)
	assert.Nil(t, prop.RefParam)
	require.Len(t, prop.Default, 1)
	assert.Equal(t, "", prop.Qualifier)
	assert.Equal(t, "", prop.Annotation)
	assert.Equal(t, "", prop.AnnotationComment)

	propDefault = prop.Default[0]
	assert.Equal(t, Literal, propDefault.Type)
	assert.Equal(t, "2", propDefault.Value)

	// Enum::Flag3

	prop = enum.Children[3]
	assert.Exactly(t, enum, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "Flag3", prop.Value)
	assert.Nil(t, prop.Ref)
	assert.Nil(t, prop.RefParam)
	require.Len(t, prop.Default, 1)
	assert.Equal(t, "", prop.Qualifier)
	assert.Equal(t, "", prop.Annotation)
	assert.Equal(t, "", prop.AnnotationComment)

	propDefault = prop.Default[0]
	assert.Equal(t, Literal, propDefault.Type)
	assert.Equal(t, "4", propDefault.Value)

	// Enum::Flag4

	prop = enum.Children[4]
	assert.Exactly(t, enum, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "Flag4", prop.Value)
	assert.Nil(t, prop.Ref)
	assert.Nil(t, prop.RefParam)
	require.Len(t, prop.Default, 1)
	assert.Equal(t, "", prop.Qualifier)
	assert.Equal(t, "", prop.Annotation)
	assert.Equal(t, "", prop.AnnotationComment)

	propDefault = prop.Default[0]
	assert.Equal(t, Literal, propDefault.Type)
	assert.Equal(t, "6", propDefault.Value)

	// Class

	class := root.Children[1]
	assert.Exactly(t, root, class.Parent)
	require.Len(t, class.Children, 4)
	assert.Equal(t, Class, class.Type)
	assert.Equal(t, "Data", class.Value)
	require.NotNil(t, class.Ref)
	assert.Equal(t, Property, class.Ref.Type)
	assert.Equal(t, "Field", class.Ref.Value)
	assert.Exactly(t, enum, class.Ref.Parent)
	assert.Nil(t, class.RefParam)
	assert.Nil(t, class.Default)
	assert.Equal(t, "", class.Qualifier)
	assert.Equal(t, "", class.Annotation)
	assert.Equal(t, "", class.AnnotationComment)

	// Class::a

	prop = class.Children[0]
	assert.Exactly(t, class, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "a", prop.Value)
	require.NotNil(t, prop.Ref)
	assert.Equal(t, Type, prop.Ref.Type)
	assert.Equal(t, "uint", prop.Ref.Value)
	assert.Nil(t, prop.RefParam)
	require.Len(t, prop.Default, 1)
	assert.Equal(t, "", prop.Qualifier)
	assert.Equal(t, "", prop.Annotation)
	assert.Equal(t, "", prop.AnnotationComment)

	propDefault = prop.Default[0]
	assert.Equal(t, Literal, propDefault.Type)
	assert.Equal(t, "1", propDefault.Value)

	// Class::b

	prop = class.Children[1]
	assert.Exactly(t, class, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "b", prop.Value)
	require.NotNil(t, prop.Ref)
	assert.Equal(t, Type, prop.Ref.Type)
	assert.Equal(t, "ulong", prop.Ref.Value)
	assert.Nil(t, prop.RefParam)
	require.Len(t, prop.Default, 1)
	assert.Equal(t, "", prop.Qualifier)
	assert.Equal(t, "obsolete", prop.Annotation)
	assert.Equal(t, "Not used", prop.AnnotationComment)

	propDefault = prop.Default[0]
	assert.Equal(t, Literal, propDefault.Type)
	assert.Equal(t, "-2", propDefault.Value)

	// Class::c

	prop = class.Children[2]
	assert.Exactly(t, class, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "c", prop.Value)
	require.NotNil(t, prop.Ref)
	assert.Equal(t, Enum, prop.Ref.Type)
	assert.Equal(t, "Enum", prop.Ref.Value)
	assert.Nil(t, prop.RefParam)
	require.Len(t, prop.Default, 2)
	assert.Equal(t, "", prop.Qualifier)
	assert.Equal(t, "", prop.Annotation)
	assert.Equal(t, "", prop.AnnotationComment)

	propDefault = prop.Default[0]
	assert.Equal(t, Property, propDefault.Type)
	assert.Equal(t, "Flag1", propDefault.Value)
	assert.Exactly(t, enum, propDefault.Parent)

	propDefault = prop.Default[1]
	assert.Equal(t, Property, propDefault.Type)
	assert.Equal(t, "Flag2", propDefault.Value)
	assert.Exactly(t, enum, propDefault.Parent)

	// Class::d

	prop = class.Children[3]
	assert.Exactly(t, class, prop.Parent)
	assert.Len(t, prop.Children, 0)
	assert.Equal(t, Property, prop.Type)
	assert.Equal(t, "d", prop.Value)
	require.NotNil(t, prop.Ref)
	assert.Equal(t, Literal, prop.Ref.Type)
	assert.Equal(t, "SteamKit2.GC.Internal.CMsgProtoBufHeader", prop.Ref.Value)
	require.NotNil(t, prop.RefParam)
	assert.Equal(t, Property, prop.RefParam.Type)
	assert.Equal(t, "a", prop.RefParam.Value)
	assert.Exactly(t, class, prop.RefParam.Parent)
	assert.Nil(t, prop.Default)
	assert.Equal(t, "proto", prop.Qualifier)
	assert.Equal(t, "", prop.Annotation)
	assert.Equal(t, "", prop.AnnotationComment)
}
