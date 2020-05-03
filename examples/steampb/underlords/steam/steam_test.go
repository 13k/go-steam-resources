package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/go-steam-resources/v2/steampb/underlords/steam"
)

func TestProtobufUnderlordsSteam(t *testing.T) {
	assert.Equal(t, int32(2), steam.Default_CMsgProtoBufHeader_Eresult)
}
