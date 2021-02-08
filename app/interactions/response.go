package interactions

// https://discord.com/developers/docs/interactions/slash-commands#interaction-response

// InteractionResponse - The base model of a response to an interaction request
type InteractionResponse struct {
	Type InteractionResponseType                    `json:"type"`
	Data *InteractionApplicationCommandCallbackData `json:"data"`
}

// InteractionResponseType - The type of the response
type InteractionResponseType int

// InteractionResponseType enums
const (
	Pong                     InteractionResponseType = 1
	Acknowledge              InteractionResponseType = 2
	ChannelMessage           InteractionResponseType = 3
	ChannelMessageWithSource InteractionResponseType = 4
	AcknowledgeWithSource    InteractionResponseType = 5
)

// InteractionApplicationCommandCallbackData - Optional response message payload
type InteractionApplicationCommandCallbackData struct {
	TTS             bool           `json:"tts"`
	Content         string         `json:"content"`
	Embeds          []*interface{} `json:"embeds"`           // todo
	AllowedMentions interface{}    `json:"allowed_mentions"` // todo
}
