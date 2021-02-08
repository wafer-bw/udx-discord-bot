package models

// https://discord.com/developers/docs/interactions/slash-commands#interaction

// InteractionRequest - The base request model sent when a user invokes a command
type InteractionRequest struct {
	ID        string                             `json:"id"`
	Type      InteractionType                    `json:"type"`
	Data      *ApplicationCommandInteractionData `json:"data"`
	GuildID   string                             `json:"guild_id"`
	ChannelID string                             `json:"channel_id"`
	Member    interface{}                        `json:"member"` // todo
	Token     string                             `json:"token"`
	Version   int                                `json:"version"`
}

// InteractionType - The type of the interaction
type InteractionType int

// InteractionTypeEnum - Acts as an enum struct of all `IntaractionType`s
type InteractionTypeEnum struct {
	Ping               InteractionType
	ApplicationCommand InteractionType
}

// InteractionTypes - `InteractionTypeEnum`
var InteractionTypes = &InteractionTypeEnum{
	Ping:               1,
	ApplicationCommand: 2,
}

// ApplicationCommandInteractionData - The command data payload
type ApplicationCommandInteractionData struct {
	ID      string                                     `json:"id"`
	Name    string                                     `json:"name"`
	Options []*ApplicationCommandInteractionDataOption `json:"options"`
}

// ApplicationCommandInteractionDataOption - The params + values from the user
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Value   string                                     `json:"value"`
	Options []*ApplicationCommandInteractionDataOption `json:"options"`
}
