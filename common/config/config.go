package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/wafer-bw/disgoslash/discord"
)

// EnvVars defines expected & required environment variables
type EnvVars struct {
	DiscordPublicKey string `envconfig:"DISCORD_PUBLIC_KEY" required:"true" split_words:"true"`
	DiscordClientID  string `envconfig:"DISCORD_CLIENT_ID" required:"true" split_words:"true"`
	DiscordToken     string `envconfig:"DISCORD_TOKEN" required:"true" split_words:"true"`
	TradierEndpoint  string `envconfig:"TRADIER_ENDPOINT" required:"true" split_words:"true"`
	TradierToken     string `envconfig:"TRADIER_TOKEN" required:"true" split_words:"true"`
}

// TradierConfig data
type TradierConfig struct {
	Endpoint string
	Token    string
}

// Config data
type Config struct {
	Discord *discord.Credentials
	Tradier *TradierConfig
}

// New returns a new `Config` struct; panics if unable
func New() *Config {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("Warning: could not load .env file")
	// }
	env := getEnvVars()
	ensureNoBlankEnvVars(env)
	return &Config{
		Discord: &discord.Credentials{
			PublicKey: env.DiscordPublicKey,
			ClientID:  env.DiscordClientID,
			Token:     env.DiscordToken,
		},
		Tradier: &TradierConfig{
			Endpoint: env.TradierEndpoint,
			Token:    env.TradierToken,
		},
	}
}

func getEnvVars() EnvVars {
	var env EnvVars
	err := envconfig.Process("", &env)
	if err != nil {
		panic(err)
	}
	return env
}

func ensureNoBlankEnvVars(env EnvVars) {
	blanks := findBlankEnvVars(env)
	if len(blanks) > 0 {
		panic(fmt.Errorf("the following environment variables are blank: %s", strings.Join(blanks, ", ")))
	}
}

func findBlankEnvVars(env EnvVars) []string {
	var blanks []string
	valueOfStruct := reflect.ValueOf(env)
	typeOfStruct := valueOfStruct.Type()
	for i := 0; i < valueOfStruct.NumField(); i++ {
		if valueOfStruct.Field(i).Interface() == "" {
			blanks = append(blanks, typeOfStruct.Field(i).Name)
		}
	}
	return blanks
}
