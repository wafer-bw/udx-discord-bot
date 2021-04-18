package leap

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/disgoslash/discord"
	"github.com/wafer-bw/udx-discord-bot/common/config"
	tapiMocks "github.com/wafer-bw/udx-discord-bot/generatedmocks/common/apis/tradier"
)

var tapiMock *tapiMocks.ClientInterface
var confMock *config.Config

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	tapiMock = &tapiMocks.ClientInterface{}
	confMock = &config.Config{Tradier: &config.TradierConfig{Endpoint: "https://sandbox.tradier.com/v1", Token: "1"}}
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestLeapWrapper(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		request := &discord.InteractionRequest{}
		require.Panics(t, func() { leapWrapper(request) })
	})
}

func TestLeap(t *testing.T) {
	// TODO
}
