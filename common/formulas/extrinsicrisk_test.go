package formulas

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

func TestExtrinsicValue(t *testing.T) {
	expect := 4.616895874263264
	er := GetExtrinsicValue(101.80, 87.50, 19)
	require.Equal(t, expect, er)
}
