package models

// SlashCommands - A map of slash command objects using the command names as keys
type SlashCommands map[string]SlashCommand

// Action is executed to resolve an interaction request
type Action func(request *InteractionRequest) (*InteractionResponse, error)

// SlashCommand - All the data required for a slash command
type SlashCommand struct {
	Name    string
	Action  Action
	Command *ApplicationCommand
}
