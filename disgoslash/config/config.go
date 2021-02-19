package config

// Deps defines `Config` dependencies
type Deps struct{}

// Config holds all config data
type Config struct {
	Credentials *Credentials
	DiscordAPI  *DiscordAPI
}

// DiscordAPI config data
type DiscordAPI struct {
	BaseURL     string
	APIVersion  string
	ContentType string
}

// Credentials config data
type Credentials struct {
	PublicKey string
	ClientID  string
	Token     string
}

// New returns a new `Config` struct; panics if unable
func New(creds *Credentials) *Config {
	return &Config{
		Credentials: creds,
		DiscordAPI: &DiscordAPI{
			BaseURL:     "https://discord.com/api",
			APIVersion:  "v8",
			ContentType: "application/json",
		},
	}
}
