package config

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestMustNewPanics(t *testing.T) {
	require.Panics(t, func() { New() })
}

func TestFindBlankEnvVars(t *testing.T) {
	blanks := findBlankEnvVars(EnvVars{DiscordToken: "test"})
	for _, b := range blanks {
		require.NotEqual(t, "DiscordToken", b)
	}
}

func TestMustGetEnvVars(t *testing.T) {
	require.Panics(t, func() { getEnvVars() })
}

func TestMustHaveNoBlankEnvVars(t *testing.T) {
	require.Panics(t, func() { ensureNoBlankEnvVars(EnvVars{}) })
}
