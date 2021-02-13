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
	"github.com/wafer-bw/udx-discord-bot/app/models"
)

var guildID = "1234567890"

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestListApplicationCommands(t *testing.T) {
	t.Run("success/global", func(t *testing.T) {
		commandName := `hello`
		mockResponse := fmt.Sprintf(`[{"id": "12345","application_id": "12345","name": "%s","description": "says hello"}]`, commandName)
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			// nolint
			w.Write([]byte(mockResponse))
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		commands, err := clientImpl.ListApplicationCommands("")
		assert.NoError(t, err)
		assert.Equal(t, commands[0].Name, commandName)
	})
	t.Run("success/guild", func(t *testing.T) {
		commandName := `hello`
		mockResponse := fmt.Sprintf(`[{"id": "12345","application_id": "12345","name": "%s","description": "says hello"}]`, commandName)
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
}

func TestDeleteApplicationCommand(t *testing.T) {
	t.Run("success/global", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		err := clientImpl.DeleteApplicationCommand("", "12345")
		assert.NoError(t, err)
	})
	t.Run("success/guild", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		err := clientImpl.DeleteApplicationCommand("12345", "12345")
		assert.NoError(t, err)
	})
	t.Run("failure/unauthorized", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		err := clientImpl.DeleteApplicationCommand("12345", "12345")
		assert.Error(t, err)
	})
}

func TestCreateApplicationCommand(t *testing.T) {
	t.Run("success/global", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		err := clientImpl.CreateApplicationCommand("", &models.ApplicationCommand{})
		assert.NoError(t, err)
	})
	t.Run("success/guild", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		err := clientImpl.CreateApplicationCommand("12345", &models.ApplicationCommand{})
		assert.NoError(t, err)
	})
	t.Run("failure/unauthorized", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		}))
		defer func() { mockServer.Close() }()
		mocks.Conf.DiscordAPI.BaseURL = mockServer.URL
		clientImpl := New(&Deps{}, mocks.Conf)

		err := clientImpl.CreateApplicationCommand("12345", &models.ApplicationCommand{})
		assert.Error(t, err)
	})
}
