package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	src = `
#import "enum.steamd"

// This is not a real SteamLanguage class
class Data<Enum::Field> {
	A = 1;
	B = 3.1417;
	C = -2;
	D = -123; removed "For pressing ceremonial reasons"
	E = Enum::Flag1 | Enum::Flag2;
};
`

	stream = []struct {
		t TokenType
		v string
	}{
		{TokenPreprocess, "import"},
		{TokenString, "enum.steamd"},
		{TokenIdent, "class"},
		{TokenIdent, "Data"},
		{TokenTypeParam, "<"},
		{TokenIdent, "Enum::Field"},
		{TokenTypeParam, ">"},
		{TokenScope, "{"},
		{TokenIdent, "A"},
		{TokenOperator, "="},
		{TokenInt, "1"},
		{TokenTerminator, ";"},
		{TokenIdent, "B"},
		{TokenOperator, "="},
		{TokenFloat, "3.1417"},
		{TokenTerminator, ";"},
		{TokenIdent, "C"},
		{TokenOperator, "="},
		{TokenInt, "-2"},
		{TokenTerminator, ";"},
		{TokenIdent, "D"},
		{TokenOperator, "="},
		{TokenInt, "-123"},
		{TokenTerminator, ";"},
		{TokenIdent, "removed"},
		{TokenString, "For pressing ceremonial reasons"},
		{TokenIdent, "E"},
		{TokenOperator, "="},
		{TokenIdent, "Enum::Flag1"},
		{TokenOperator, "|"},
		{TokenIdent, "Enum::Flag2"},
		{TokenTerminator, ";"},
		{TokenScope, "}"},
		{TokenTerminator, ";"},
	}
)

func TestScannerScan(t *testing.T) {
	s := NewScannerString(src, "")

	for _, tc := range stream {
		token, err := s.Scan()

		assert.Nil(t, err)
		assert.Equalf(t, tc.t, token.Type, "Token: %s. Expected type: %q, actual type: %q", token, tc.t, token.Type)
		assert.Equalf(t, tc.v, token.Value, "Token: %s", token)
	}
}

func TestScannerPeek(t *testing.T) {
	s := NewScannerString(src, "")

	for _, tc := range stream {
		peek1, err := s.Peek()

		assert.NoError(t, err)
		assert.Equalf(t, tc.t, peek1.Type, "Token: %s. Expected type: %q, actual type: %q", peek1, tc.t, peek1.Type)
		assert.Equalf(t, tc.v, peek1.Value, "Token: %s", peek1)

		peek2, err := s.Peek()

		assert.NoError(t, err)
		assert.Equalf(t, tc.t, peek2.Type, "Token: %s. Expected type: %q, actual type: %q", peek2, tc.t, peek2.Type)
		assert.Equalf(t, tc.v, peek2.Value, "Token: %s", peek2)

		assert.Equal(t, peek1, peek2, "Expected token: %s, actual token: %q", peek1, peek2)

		token, err := s.Scan()

		assert.NoError(t, err)
		assert.Equalf(t, tc.t, token.Type, "Token: %s. Expected type: %q, actual type: %q", token, tc.t, token.Type)
		assert.Equalf(t, tc.v, token.Value, "Token: %s", token)

		assert.Equal(t, peek1, token, "Expected token: %s, actual token: %q", peek1, token)
	}
}
