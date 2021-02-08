package models

import "time"

// Embed - an embed object
type Embed struct {
	Title       string         `json:"title"`
	Type        string         `json:"type"`
	Description string         `json:"description"`
	URL         string         `json:"url"`
	Timestamp   time.Time      `json:"timestamp"` // todo verify if this works in unmarshal or if it needs to be switched to a string
	Color       int            `json:"color"`
	Footer      interface{}    `json:"footer"`    // todo struct
	Image       interface{}    `json:"image"`     // todo struct
	Thumbnail   interface{}    `json:"thumbnail"` // todo struct
	Video       interface{}    `json:"video"`     // todo struct
	Provider    interface{}    `json:"provider"`  // todo struct
	Author      interface{}    `json:"author"`    // todo struct
	Fields      []*interface{} `json:"fields"`    // todo struct https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
}

// AllowedMentions - Used to control mentions
type AllowedMentions struct {
	Parse       []string `json:"parse"` // todo enum https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mention-types
	Roles       []string `json:"roles"`
	Users       []string `json:"users"`
	RepliedUser bool     `json:"replied_user"`
}
