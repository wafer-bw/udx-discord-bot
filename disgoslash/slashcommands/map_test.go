package slashcommands

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/models"
)

var mockCommandName = "hello"
var mockCommandResponseContent = "Hello World!"

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

func TestNewMap(t *testing.T) {
	slashCommands := NewMap(
		New(mockCommandName, &models.ApplicationCommand{Name: mockCommandName, Description: "desc"}, mockDo, true, []string{"11111"}),
	)
	require.Equal(t, 1, len(slashCommands))
}
