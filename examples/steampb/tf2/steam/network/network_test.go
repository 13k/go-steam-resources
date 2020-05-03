package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/go-steam-resources/v2/steampb/tf2/steam/network"
)

func TestProtobufTf2SteamNetwork(t *testing.T) {
	assert.Equal(
		t,
		"k_ESteamDatagramMsg_ConnectRequest",
		network.ESteamDatagramMsgID_k_ESteamDatagramMsg_ConnectRequest.String(),
	)
}
