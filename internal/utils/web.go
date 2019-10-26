package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
)

//FetchPage return the body content of the url
func FetchPage(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", errors.New("Cannot fetch " + url)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", errors.New("Cannot read body of " + url)
	}

	return string(bodyBytes), nil
}
