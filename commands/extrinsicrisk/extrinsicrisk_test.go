package extrinsicrisk

import (
	"encoding/json"
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

func TestExtrinsicRisk(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: json.RawMessage(`"101.80"`)},
					{Name: "strike", Value: json.RawMessage(`"87.5"`)},
					{Name: "ask", Value: json.RawMessage(`"19"`)},
				},
			},
		}
		expect := "4.62%"

		response := extrinsicRisk(request)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/missing options", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicRisk(request)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid share value", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: json.RawMessage(`"a"`)},
					{Name: "strike", Value: json.RawMessage(`"87.5"`)},
					{Name: "ask", Value: json.RawMessage(`"19"`)},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicRisk(request)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid strike value", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: json.RawMessage(`"101.80"`)},
					{Name: "strike", Value: json.RawMessage(`"a"`)},
					{Name: "ask", Value: json.RawMessage(`"19"`)},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicRisk(request)
		require.Equal(t, expect, response.Data.Content)
	})

	t.Run("failure/invalid ask value", func(t *testing.T) {
		request := &discord.InteractionRequest{
			Data: &discord.ApplicationCommandInteractionData{
				Options: []*discord.ApplicationCommandInteractionDataOption{
					{Name: "share", Value: json.RawMessage(`"101.80"`)},
					{Name: "strike", Value: json.RawMessage(`"87.5"`)},
					{Name: "ask", Value: json.RawMessage(`"a"`)},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicRisk(request)
		require.Equal(t, expect, response.Data.Content)
	})
}
