package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	interactionsMocks "github.com/wafer-bw/discobottest/app/generatedmocks/app/interactions"
	"github.com/wafer-bw/discobottest/app/mocks"
)

var url = "http://localhost/api"
var interactionsMock = &interactionsMocks.Interactions{}
var handlerImpl = New(&Deps{Interactions: interactionsMock}, mocks.Conf)
var handler = http.HandlerFunc(handlerImpl.Handle)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestNew(t *testing.T) {
	require.NotNil(t, handlerImpl)
	require.IsType(t, &impl{}, handlerImpl)
}

func TestHandle(t *testing.T) {
	headers := map[string]string{"Accept": "application/json"}
	interactionsMock.On("Verify", mock.Anything, mock.Anything, mock.Anything).Return(true).Times(1)
	body, resp, err := httpRequest(http.MethodGet, url, headers, mocks.PingRequestBody)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode, string(body))
}

func httpRequest(method string, url string, headers map[string]string, body string) ([]byte, *http.Response, error) {
	request := httptest.NewRequest(method, url, strings.NewReader(body))
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody, err := ioutil.ReadAll(recorder.Body)
	return responseBody, response, err
}
