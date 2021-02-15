package slashcommands

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/models"
)

func TestNew(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		name := "HelloWorld"
		command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
		slashCommand := New(name, command, nil, true, []string{"12345"})
		require.Equal(t, strings.ToLower(name), slashCommand.Name)
		require.Equal(t, name, slashCommand.AppCommand.Name)
		require.Equal(t, 2, len(slashCommand.GuildIDs))
	})
	t.Run("success/global only", func(t *testing.T) {
		name := "HelloWorld"
		command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
		slashCommand := New(name, command, nil, true, []string{})
		require.Equal(t, strings.ToLower(name), slashCommand.Name)
		require.Equal(t, name, slashCommand.AppCommand.Name)
		require.Equal(t, 1, len(slashCommand.GuildIDs))
	})
	t.Run("success/guild only", func(t *testing.T) {
		name := "HelloWorld"
		command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
		slashCommand := New(name, command, nil, false, []string{"12345"})
		require.Equal(t, strings.ToLower(name), slashCommand.Name)
		require.Equal(t, name, slashCommand.AppCommand.Name)
		require.Equal(t, 1, len(slashCommand.GuildIDs))
	})
	t.Run("success/accepts nil guildIDs slice", func(t *testing.T) {
		name := "HelloWorld"
		command := &models.ApplicationCommand{Name: name, Description: "Says hello world!"}
		slashCommand := New(name, command, nil, false, nil)
		require.Equal(t, strings.ToLower(name), slashCommand.Name)
		require.Equal(t, name, slashCommand.AppCommand.Name)
		require.Equal(t, 0, len(slashCommand.GuildIDs))
	})
}
