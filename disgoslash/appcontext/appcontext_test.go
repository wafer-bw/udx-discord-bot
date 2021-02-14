package appcontext

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestNewPanics(t *testing.T) {
	require.Panics(t, func() { New(slashcommands.Map{}) })
}
