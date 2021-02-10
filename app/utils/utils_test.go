package utils

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type testobj struct {
	Name   string   `json:"name"`
	Rename bool     `json:"something_else"`
	Arr    []string `json:"strings"`
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestFormatJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		obj := &testobj{
			Name:   "TEST",
			Rename: false,
			Arr:    []string{"a", "b", "c"},
		}
		expect := "{\n    \"name\": \"TEST\",\n    \"something_else\": false,\n    \"strings\": [\n        \"a\",\n        \"b\",\n        \"c\"\n    ]\n}"
		res := FormatJSON(obj)
		require.Equal(t, expect, res)
	})

	t.Run("failure/panics", func(t *testing.T) {
		require.Panics(t, func() { FormatJSON(func() {}) })
	})
}
