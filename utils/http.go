package utils

import (
	"net/http"
)

var EMPTY_HTTP_RESPONSE = &http.Response{}

func HttpGetHeaders(link string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return EMPTY_HTTP_RESPONSE, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return EMPTY_HTTP_RESPONSE, err
	}

	return response, nil
}
