package models

// https://discord.com/developers/docs/interactions/slash-commands#interaction-response

// InteractionResponse - The base model of a response to an interaction request
type InteractionResponse struct {
	Type InteractionResponseType                    `json:"type"`
	Data *InteractionApplicationCommandCallbackData `json:"data"`
}

// InteractionResponseType - The type of the response
type InteractionResponseType int

// InteractionResponseTypeEnum - Acts as an enum struct of all `InteractionResponseType`s
type InteractionResponseTypeEnum struct {
	Pong                     InteractionResponseType
	Acknowledge              InteractionResponseType
	ChannelMessage           InteractionResponseType
	ChannelMessageWithSource InteractionResponseType
	AcknowledgeWithSource    InteractionResponseType
}

// InteractionResponseTypes - `InteractionResponseTypeEnum`
var InteractionResponseTypes = &InteractionResponseTypeEnum{
	Pong:                     1,
	Acknowledge:              2,
	ChannelMessage:           3,
	ChannelMessageWithSource: 4,
	AcknowledgeWithSource:    5,
}

// InteractionApplicationCommandCallbackData - Optional response message payload
type InteractionApplicationCommandCallbackData struct {
	TTS             bool           `json:"tts"`
	Content         string         `json:"content"`
	Embeds          []*interface{} `json:"embeds"`           // todo
	AllowedMentions interface{}    `json:"allowed_mentions"` // todo
}
