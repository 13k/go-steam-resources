package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/go-steam-resources/steamlang"
)

func TestSteamlang(t *testing.T) {
	assert.Equal(t, steamlang.EMsg(1), steamlang.EMsg_Multi)
	assert.Equal(t, steamlang.EResult(1), steamlang.EResult_OK)
}
