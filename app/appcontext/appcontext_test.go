package appcontext

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/discobottest/app/mocks"
)

var appcontextImpl *AppContext

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	appcontextImpl = New(&Deps{}, mocks.Conf)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestNew(t *testing.T) {
	require.NotNil(t, appcontextImpl)
	require.IsType(t, &AppContext{}, appcontextImpl)
}
