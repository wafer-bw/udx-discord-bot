package chstrat

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

func TestChstratWrapper(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		request := &discord.InteractionRequest{}
		require.Panics(t, func() { chstratWrapper(request) })
	})
}

func TestChstrat(t *testing.T) {
	// t.Run("LIVE RUN", func(t *testing.T) {
	// 	err := godotenv.Load("../../.env")
	// 	if err != nil {
	// 		log.Println("Warning: could not load .env file")
	// 	}
	// 	request := &discord.InteractionRequest{
	// 		Data: &discord.ApplicationCommandInteractionData{
	// 			Options: []*discord.ApplicationCommandInteractionDataOption{
	// 				{Name: "symbol", Value: "SPY"},
	// 			},
	// 		},
	// 	}
	// 	response := chstratWrapper(request)
	// 	require.NotNil(t, response)
	// 	require.Equal(t, "", response.Data.Content)
	// })

	// t.Run("success", func(t *testing.T) {
	// 	now := time.Now()
	// 	now = now.AddDate(0, 0, 101)
	// 	symbol := "SPY"
	// 	expirations := tradier.Expirations{now.Format("2006-01-02")}
	// 	request := &discord.InteractionRequest{
	// 		Data: &discord.ApplicationCommandInteractionData{
	// 			Options: []*discord.ApplicationCommandInteractionDataOption{
	// 				{Name: "symbol", Value: symbol},
	// 			},
	// 		},
	// 	}

	// 	tapiMock.On("GetQuote", symbol, false).Return(&tradier.Quote{Last: 200.00}, nil).Times(1)
	// 	tapiMock.On("GetOptionExpirations", symbol, true, false).Return(expirations, nil).Times(1)

	// 	response := chstrat(request, tapiMock, time.Now())
	// 	fmt.Println(response.Data.Content)
	// 	require.Nil(t, response)
	// })
}
