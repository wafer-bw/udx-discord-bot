package chstrat

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/disgoslash/models"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestChstrat(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Options: []*models.ApplicationCommandInteractionDataOption{
					{Name: "symbol", Value: "AAPL"},
					{Name: "Asset Class", Value: "stocks"},
				},
			},
		}

		response, err := chstrat(request)
		require.Nil(t, err)
		require.NotNil(t, response)
	})
}
