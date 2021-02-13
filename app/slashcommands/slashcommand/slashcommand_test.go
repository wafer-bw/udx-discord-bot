package slashcommand

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/app/models"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestIsGlobal(t *testing.T) {
	t.Run("success/is global", func(t *testing.T) {
		sc := SlashCommand{GuildIDs: []string{"", "12345"}}
		require.True(t, sc.IsGlobal())
	})
	t.Run("success/is not global", func(t *testing.T) {
		sc := SlashCommand{GuildIDs: []string{"12345", "67890"}}
		require.False(t, sc.IsGlobal())
	})
}

func TestNew(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		name := "HelloWorld"
		command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
		slashCommand := New(name, command, nil, true, []string{"12345"})
		require.Equal(t, strings.ToLower(name), slashCommand.Name)
		require.Equal(t, name, slashCommand.Command.Name)
		require.Equal(t, 2, len(slashCommand.GuildIDs))
	})
	t.Run("success/global only", func(t *testing.T) {
		name := "HelloWorld"
		command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
		slashCommand := New(name, command, nil, true, []string{})
		require.Equal(t, strings.ToLower(name), slashCommand.Name)
		require.Equal(t, name, slashCommand.Command.Name)
		require.Equal(t, 1, len(slashCommand.GuildIDs))
	})
	t.Run("success/guild only", func(t *testing.T) {
		name := "HelloWorld"
		command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
		slashCommand := New(name, command, nil, false, []string{"12345"})
		require.Equal(t, strings.ToLower(name), slashCommand.Name)
		require.Equal(t, name, slashCommand.Command.Name)
		require.Equal(t, 1, len(slashCommand.GuildIDs))
	})
	t.Run("success/accepts nil guildIDs slice", func(t *testing.T) {
		name := "HelloWorld"
		command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
		slashCommand := New(name, command, nil, false, nil)
		require.Equal(t, strings.ToLower(name), slashCommand.Name)
		require.Equal(t, name, slashCommand.Command.Name)
		require.Equal(t, 0, len(slashCommand.GuildIDs))
	})
}
