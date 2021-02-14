package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	clientMocks "github.com/wafer-bw/udx-discord-bot/disgoslash/generatedmocks/disgoslash/client"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/models"
)

var clientMock = &clientMocks.Client{}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestList(t *testing.T) {
	clientResult := []*models.ApplicationCommand{
		{
			ID:          "1",
			Name:        "Name",
			Description: "Description",
			Options:     []*models.ApplicationCommandOption{},
		},
	}

	t.Run("success", func(t *testing.T) {
		expect := "1 - Name: Description\n"
		clientMock.On("ListApplicationCommands", "32472384723").Return(clientResult, nil).Times(1)
		res := list(clientMock, "32472384723", false)
		require.Equal(t, expect, res)
	})

	t.Run("success/verbose", func(t *testing.T) {
		expect := "[\n    {\n        \"id\": \"1\",\n        \"application_id\": \"\",\n        \"name\": \"Name\",\n        \"description\": \"Description\",\n        \"options\": []\n    }\n]"
		clientMock.On("ListApplicationCommands", "32472384723").Return(clientResult, nil).Times(1)
		res := list(clientMock, "32472384723", true)
		require.Equal(t, expect, res)
	})
}
