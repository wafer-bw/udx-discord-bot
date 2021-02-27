package creds

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/wafer-bw/disgoslash/discord"
)

// EnvVars defines expected & required environment variables
type EnvVars struct {
	PublicKey string `envconfig:"PUBLIC_KEY" required:"true" split_words:"true"`
	ClientID  string `envconfig:"CLIENT_ID" required:"true" split_words:"true"`
	Token     string `envconfig:"TOKEN" required:"true" split_words:"true"`
}

// New returns a new `Config` struct; panics if unable
func New() *discord.Credentials {
	env := getEnvVars()
	ensureNoBlankEnvVars(env)
	return &discord.Credentials{
		PublicKey: env.PublicKey,
		ClientID:  env.ClientID,
		Token:     env.Token,
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