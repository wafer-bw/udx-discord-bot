package common

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Request makes an HTTP request, responding with `statusCode int`, `body []bytes`, `error`
func Request(method string, url string, headers map[string]string, body io.Reader) (int, []byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return 0, nil, err
	}

	for key, val := range headers {
		request.Header.Set(key, val)
	}
	response, err := client.Do(request)
	if err != nil {
		return 0, nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, nil, err
	}

	return response.StatusCode, data, nil
}
