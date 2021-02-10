package models

// Action is executed to resolve an interaction request
type Action func(request *InteractionRequest) (*InteractionResponse, error)

// SlashCommand - All the data required for a slash command
type SlashCommand struct {
	Name    string
	Action  Action
	Command *ApplicationCommand
}
