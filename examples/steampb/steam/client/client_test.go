package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/go-steam-resources/v2/steampb/steam/client"
)

func TestProtobufSteamClient(t *testing.T) {
	assert.Equal(
		t,
		"k_EProtoExecutionSiteSteamClient",
		client.EProtoExecutionSite_k_EProtoExecutionSiteSteamClient.String(),
	)
}
