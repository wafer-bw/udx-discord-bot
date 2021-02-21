package extrinsicrisk

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

func TestExtrinsicRisk(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Options: []*models.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: "101.80"},
					{Name: "strike", Value: "87.5"},
					{Name: "ask", Value: "19"},
				},
			},
		}
		expect := "4.62%"

		response, err := extrinsicRisk(request)
		require.Nil(t, err)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/missing options", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Options: []*models.ApplicationCommandInteractionDataOption{},
			},
		}
		expect := "Error parsing command :cry:"

		response, err := extrinsicRisk(request)
		require.Nil(t, err)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid share value", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Options: []*models.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: "a"},
					{Name: "strike", Value: "87.5"},
					{Name: "ask", Value: "19"},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response, err := extrinsicRisk(request)
		require.Nil(t, err)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid strike value", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Options: []*models.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: "101.80"},
					{Name: "strike", Value: "a"},
					{Name: "ask", Value: "19"},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response, err := extrinsicRisk(request)
		require.Nil(t, err)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid ask value", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Options: []*models.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: "101.80"},
					{Name: "strike", Value: "87.5"},
					{Name: "ask", Value: "a"},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response, err := extrinsicRisk(request)
		require.Nil(t, err)
		require.Equal(t, expect, response.Data.Content)
	})
}
