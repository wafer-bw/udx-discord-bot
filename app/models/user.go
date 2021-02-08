package models

// User - A discord user
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
	System        bool   `json:"system"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	Locale        string `json:"locale"`
	Verified      bool   `json:"verified"`
	Email         string `json:"email"`
	Flags         int    `json:"flags"`        // todo enum https://discord.com/developers/docs/resources/user#user-object-user-flags
	PremiumType   int    `json:"premium_type"` // todo enum https://discord.com/developers/docs/resources/user#user-object-premium-types
	PublicFlags   int    `json:"public_flags"` // todo enum https://discord.com/developers/docs/resources/user#user-object-user-flags
}
