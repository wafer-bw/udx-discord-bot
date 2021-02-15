package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/docopt/docopt-go"
	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/errs"
	clientMocks "github.com/wafer-bw/udx-discord-bot/disgoslash/generatedmocks/disgoslash/client"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/models"
)

var guildID = "guild-id"
var commandID = "command-id"
var clientMock = &clientMocks.Client{}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestLoadEnv(t *testing.T) {
	require.Panics(t, func() { loadEnv() })
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
		clientMock.On("ListApplicationCommands", commandID).Return(clientResult, nil).Times(1)
		res, err := list(clientMock, commandID, false)
		require.NoError(t, err)
		require.Equal(t, expect, res)
	})
	t.Run("success/verbose", func(t *testing.T) {
		expect := "[\n    {\n        \"id\": \"1\",\n        \"application_id\": \"\",\n        \"name\": \"Name\",\n        \"description\": \"Description\",\n        \"options\": []\n    }\n]"
		clientMock.On("ListApplicationCommands", commandID).Return(clientResult, nil).Times(1)
		res, err := list(clientMock, commandID, true)
		require.NoError(t, err)
		require.Equal(t, expect, res)
	})
	t.Run("failure", func(t *testing.T) {
		expect := ""
		clientMock.On("ListApplicationCommands", commandID).Return(clientResult, errs.ErrUnauthorized).Times(1)
		res, err := list(clientMock, commandID, true)
		require.Error(t, err)
		require.Equal(t, expect, res)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		clientMock.On("DeleteApplicationCommand", guildID, commandID).Return(nil).Times(1)
		err := delete(clientMock, guildID, commandID)
		require.NoError(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		clientMock.On("DeleteApplicationCommand", guildID, commandID).Return(errs.ErrUnauthorized).Times(1)
		err := delete(clientMock, guildID, commandID)
		require.Error(t, err)
	})
}

func TestCreate(t *testing.T) {
	cmd := &models.ApplicationCommand{}
	t.Run("success", func(t *testing.T) {
		clientMock.On("CreateApplicationCommand", guildID, cmd).Return(nil).Times(1)
		err := create(clientMock, guildID, cmd)
		require.NoError(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		clientMock.On("CreateApplicationCommand", guildID, cmd).Return(errs.ErrUnauthorized).Times(1)
		err := create(clientMock, guildID, cmd)
		require.Error(t, err)
	})
}

func TestUsage(t *testing.T) {
	var usageTestTable = []struct {
		argv      []string    // Given command line args
		validArgs bool        // Are they supposed to be valid?
		opts      docopt.Opts // Expected options parsed
	}{
		{
			[]string{"list"},
			true,
			docopt.Opts{
				"list":                true,
				"register":            false,
				"unregister":          false,
				"--help":              false,
				"--verbose":           false,
				"-v":                  false,
				"<guildID>":           nil,
				"<commandID>":         nil,
				"<command-json-path>": nil,
			},
		},
		{
			[]string{"list", "-v"},
			true,
			docopt.Opts{
				"list":                true,
				"register":            false,
				"unregister":          false,
				"--help":              false,
				"--verbose":           false,
				"-v":                  true,
				"<guildID>":           nil,
				"<commandID>":         nil,
				"<command-json-path>": nil,
			},
		},
		{
			[]string{"unregister", commandID},
			true,
			docopt.Opts{
				"list":                false,
				"register":            false,
				"unregister":          true,
				"--help":              false,
				"--verbose":           false,
				"-v":                  false,
				"<guildID>":           nil,
				"<commandID>":         commandID,
				"<command-json-path>": nil,
			},
		},
		{
			[]string{"unregister", commandID, guildID},
			true,
			docopt.Opts{
				"list":                false,
				"register":            false,
				"unregister":          true,
				"--help":              false,
				"--verbose":           false,
				"-v":                  false,
				"<guildID>":           guildID,
				"<commandID>":         commandID,
				"<command-json-path>": nil,
			},
		},
		{
			[]string{"register", "./somepath", guildID},
			true,
			docopt.Opts{
				"list":                false,
				"register":            true,
				"unregister":          false,
				"--help":              false,
				"--verbose":           false,
				"-v":                  false,
				"<guildID>":           guildID,
				"<commandID>":         nil,
				"<command-json-path>": "./somepath",
			},
		},
		{
			[]string{"register", "./somepath"},
			true,
			docopt.Opts{
				"list":                false,
				"register":            true,
				"unregister":          false,
				"--help":              false,
				"--verbose":           false,
				"-v":                  false,
				"<guildID>":           nil,
				"<commandID>":         nil,
				"<command-json-path>": "./somepath",
			},
		},
		{
			[]string{"blah", "blah", "blah"},
			false,
			docopt.Opts{},
		},
	}

	for _, tt := range usageTestTable {
		validArgs := true
		parser := &docopt.Parser{
			HelpHandler: func(err error, _ string) {
				if err != nil {
					validArgs = false // Triggered usage, args were invalid.
				}
			},
		}
		opts, err := parser.ParseArgs(usage, tt.argv, "")
		if validArgs != tt.validArgs {
			t.Fail()
		}
		if tt.validArgs && err != nil {
			t.Fail()
		}
		if tt.validArgs && !reflect.DeepEqual(opts, tt.opts) {
			t.Errorf("result (1) doesn't match expected (2) \n%v \n%v", opts, tt.opts)
		}
	}
}
