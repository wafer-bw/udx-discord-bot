package commands

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/app/errs"
	"github.com/wafer-bw/udx-discord-bot/app/mocks"
	"github.com/wafer-bw/udx-discord-bot/app/models"
)

var mockCommandName = "hello"
var mockCommandDescription = "says hello"
var mockCommandResponseContent = "Hello World!"

var commandsImpl = New(&Deps{SlashCommands: []models.SlashCommand{
	NewSlashCommand(mockCommandName, &models.ApplicationCommand{Name: mockCommandName, Description: mockCommandDescription}, mockAction),
}}, mocks.Conf)

func mockAction(request *models.InteractionRequest) (*models.InteractionResponse, error) {
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

func TestNewSlashCommand(t *testing.T) {
	name := "HelloWorld"
	command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
	slashCommand := NewSlashCommand(name, command, mockAction)
	require.Equal(t, strings.ToLower(name), slashCommand.Name)
	require.Equal(t, name, slashCommand.Command.Name)
}

func TestRun(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Name: mockCommandName,
			}}
		response, err := commandsImpl.Run(request)
		require.NoError(t, err)
		require.Equal(t, response.Data.Content, mockCommandResponseContent)
	})

	t.Run("failure/command not implemented", func(t *testing.T) {
		request := &models.InteractionRequest{
			Data: &models.ApplicationCommandInteractionData{
				Name: "abc123xyz",
			}}
		_, err := commandsImpl.Run(request)
		require.Equal(t, err, errs.ErrNotImplemented)
	})
}
