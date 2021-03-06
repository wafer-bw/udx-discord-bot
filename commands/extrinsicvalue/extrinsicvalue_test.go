package extrinsicvalue

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/disgoslash/discord"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestExtrinsicValue(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: "101.80"},
					{Name: "strike", Value: "87.5"},
					{Name: "ask", Value: "19"},
				},
			},
		}
		expect := "4.62%"

		response := extrinsicValue(request)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/missing options", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicValue(request)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid share value", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: "a"},
					{Name: "strike", Value: "87.5"},
					{Name: "ask", Value: "19"},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicValue(request)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid strike value", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: "101.80"},
					{Name: "strike", Value: "a"},
					{Name: "ask", Value: "19"},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicValue(request)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid ask value", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: "101.80"},
					{Name: "strike", Value: "87.5"},
					{Name: "ask", Value: "a"},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicValue(request)
		require.Equal(t, expect, response.Data.Content)
	})
}
