package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var allTokenTypes = []TokenType{
	TokenIdent,
	TokenInt,
	TokenFloat,
	TokenString,
	TokenPreprocess,
	TokenTerminator,
	TokenScope,
	TokenTypeParam,
	TokenOperator,
}

func allTokenTypesExcept(t TokenType) (types []TokenType) {
	for _, typ := range allTokenTypes {
		if typ != t {
			types = append(types, typ)
		}
	}

	return
}

func TestTokenIsAny(t *testing.T) {
	token := &Token{}

	for _, typ := range allTokenTypes {
		token.Type = typ
		allExcept := allTokenTypesExcept(typ)

		assert.True(t, token.IsAny(allTokenTypes...))
		assert.True(t, token.IsAny(typ))
		assert.False(t, token.IsAny(allExcept...))
		assert.False(t, token.IsAny())
	}
}

func TestTokenHasAnyValue(t *testing.T) {
	token := &Token{Value: "c"}

	values := []string{"a", "b", "c", "d", "e"}

	assert.True(t, token.HasAnyValue(values...))
	assert.True(t, token.HasAnyValue("c"))

	values = []string{"a", "b", "d", "e"}

	assert.False(t, token.HasAnyValue(values...))
	assert.False(t, token.HasAnyValue())
}
