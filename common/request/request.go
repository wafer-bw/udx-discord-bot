package request

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Do makes an HTTP request, responding with `statusCode int`, `body []bytes`, `error`
func Do(method string, url string, headers map[string]string, body io.Reader) (*http.Response, []byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	for key, val := range headers {
		request.Header.Set(key, val)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	return response, data, nil
}
