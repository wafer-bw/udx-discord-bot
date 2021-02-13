package slashcommands

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/app/errs"
	"github.com/wafer-bw/udx-discord-bot/app/models"
	"github.com/wafer-bw/udx-discord-bot/app/slashcommands/slashcommand"
)

var mockCommandName = "hello"
var mockCommandResponseContent = "Hello World!"

var slashCommands = New([]slashcommand.SlashCommand{
	slashcommand.New(mockCommandName, &models.ApplicationCommand{Name: mockCommandName, Description: "desc"}, mockDo, true, []string{"11111"}),
})

func mockDo(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	return &models.InteractionResponse{
		Type: models.InteractionResponseTypeChannelMessageWithSource,
		Data: &models.InteractionApplicationCommandCallbackData{Content: mockCommandResponseContent},
	}, nil
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRun(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{Name: mockCommandName},
		}
		response, err := slashCommands.Run(request)
		require.NoError(t, err)
		require.Equal(t, response.Data.Content, mockCommandResponseContent)
	})

	t.Run("failure/command not implemented", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Name: "abc123xyz",
			}}
		_, err := slashCommands.Run(request)
		require.Equal(t, err, errs.ErrNotImplemented)
	})
}
