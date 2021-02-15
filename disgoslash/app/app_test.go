package app

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/mocks"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestLoadConf(t *testing.T) {
	t.Run("failure/panics", func(t *testing.T) {
		require.Panics(t, func() { LoadConf(nil) })
	})
	t.Run("success", func(t *testing.T) {
		require.NotPanics(t, func() { LoadConf(mocks.Conf) })
	})
}

func TestNewHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		LoadConf(mocks.Conf)
		require.NotPanics(t, func() { NewHandler(slashcommands.Map{}) })
	})
	t.Run("failure/panics", func(t *testing.T) {
		conf = nil
		require.Panics(t, func() { NewHandler(slashcommands.Map{}) })
	})

}

func TestNewClient(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		LoadConf(mocks.Conf)
		require.NotPanics(t, func() { NewClient() })
	})
	t.Run("failure/panics", func(t *testing.T) {
		conf = nil
		require.Panics(t, func() { NewClient() })
	})
}

func TestNewSyncer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		LoadConf(mocks.Conf)
		require.NotPanics(t, func() { NewSyncer() })
	})
	t.Run("failure/panics", func(t *testing.T) {
		conf = nil
		require.Panics(t, func() { NewSyncer() })
	})
}
