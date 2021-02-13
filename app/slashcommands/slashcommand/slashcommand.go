package slashcommand

import (
	"strings"

	"github.com/wafer-bw/udx-discord-bot/app/models"
)

// SlashCommand properties
type SlashCommand struct {
	Do       Do
	Name     string
	GuildIDs []string
	Command  *models.ApplicationCommand
}

// Do work
type Do func(request *models.InteractionRequest) (*models.InteractionResponse, error)

// New `SlashCommand`
func New(name string, command *models.ApplicationCommand, do Do, global bool, guildIDs []string) SlashCommand {
	if guildIDs == nil {
		guildIDs = []string{}
	}
	if global {
		guildIDs = append(guildIDs, "")
	}
	return SlashCommand{
		Name:     strings.ToLower(name),
		Command:  command,
		Do:       do,
		GuildIDs: guildIDs,
	}
}

// IsGlobal or not
func (sc SlashCommand) IsGlobal() bool {
	for _, id := range sc.GuildIDs {
		if id == "" {
			return true
		}
	}
	return false
}
