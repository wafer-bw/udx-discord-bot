package extrinsicrisk

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/disgoslash/discord"
	"github.com/wafer-bw/udx-discord-bot/common/utils"
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
					{Name: "share", String: utils.StringPointer("101.80")},
					{Name: "strike", String: utils.StringPointer("87.5")},
					{Name: "ask", String: utils.StringPointer("19")},
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
					{Name: "share", String: utils.StringPointer("a")},
					{Name: "strike", String: utils.StringPointer("87.5")},
					{Name: "ask", String: utils.StringPointer("19")},
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
					{Name: "share", String: utils.StringPointer("101.80")},
					{Name: "strike", String: utils.StringPointer("a")},
					{Name: "ask", String: utils.StringPointer("19")},
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
					{Name: "share", String: utils.StringPointer("101.80")},
					{Name: "strike", String: utils.StringPointer("87.5")},
					{Name: "ask", String: utils.StringPointer("a")},
				},
			},
		}
		expect := "Error parsing command :cry:"

		response := extrinsicRisk(request)
		require.Equal(t, expect, response.Data.Content)
	})
}
