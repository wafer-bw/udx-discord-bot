package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/app/mocks"
)

var guildID = "1234567890"

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestListApplicationCommands(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		commandName := `hello`
		mockResponse := fmt.Sprintf(`[{"id": "12345","application_id": "12345","name": "%s","description": "says hello"}]`, commandName)
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// nolint
			w.Write([]byte(mockResponse))
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		commands, err := clientImpl.ListApplicationCommands(guildID)
		assert.NoError(t, err)
		assert.Equal(t, commands[0].Name, commandName)
	})

	t.Run("failure/unauthorized", func(t *testing.T) {
		mockResponse := `{"message": "401: Unauthorized", "code": 0}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			// nolint
			w.Write([]byte(mockResponse))
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		_, err := clientImpl.ListApplicationCommands(guildID)
		require.Error(t, err)
	})

	// todo
	// t.Run("failure/not found", func(t *testing.T) {
	// })
}
