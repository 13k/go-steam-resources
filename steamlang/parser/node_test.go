package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodeLookup(t *testing.T) {
	root, err := makeRoot()
	require.NoError(t, err)

	class, err := makeNode(Class, "Msg", root)
	require.NoError(t, err)

	cProp, err := makeNode(Property, "cField1", class)
	require.NoError(t, err)

	enum, err := makeNode(Enum, "EMsg", root)
	require.NoError(t, err)

	eProp, err := makeNode(Property, "eField1", enum)
	require.NoError(t, err)

	lookup := root.Lookup("invalid")
	assert.Nil(t, lookup)

	lookup = root.Lookup("Msg")
	assert.Exactly(t, class, lookup)

	lookup = root.Lookup("Msg::cField1")
	assert.Exactly(t, cProp, lookup)

	lookup = root.Lookup("Msg::invalid")
	assert.Nil(t, lookup)

	lookup = cProp.Lookup("EMsg::eField1")
	assert.Exactly(t, eProp, lookup)
}
