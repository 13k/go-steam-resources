package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/go-steam-resources/steampb/steam"
	"github.com/13k/go-steam-resources/steampb/steam/client"
	"github.com/13k/go-steam-resources/steampb/underlords"
)

func TestProtobufUnderlords(t *testing.T) {
	assert.Equal(t, int32(2), steam.Default_CMsgProtoBufHeader_Eresult)

	assert.Equal(
		t,
		"k_EProtoExecutionSiteSteamClient",
		client.EProtoExecutionSite_k_EProtoExecutionSiteSteamClient.String(),
	)

	assert.Equal(
		t,
		"k_EMsgPlayerEliminated",
		underlords.EDACGameMessages_k_EMsgPlayerEliminated.String(),
	)
}
