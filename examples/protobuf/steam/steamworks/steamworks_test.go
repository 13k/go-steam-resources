package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/go-steam-resources/protobuf/steam/steamworks"
)

func TestProtobufSteamSteamworks(t *testing.T) {
	assert.Equal(t, "k_EProtoExecutionSiteSteamClient", steamworks.EProtoExecutionSite_k_EProtoExecutionSiteSteamClient.String())
}
