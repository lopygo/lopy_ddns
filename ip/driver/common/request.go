package common

import (
	"io/ioutil"
	"net/http"
	"time"
)

func DoRequest(requestUrl string) ([]byte, error) {
	newRequest, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	newRequest.Header.Add("User-Agent", "curl/7.64.0")

	client := &http.Client{
		Timeout: time.Duration(3) * time.Second,
	}
	r, err := client.Do(newRequest)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(r.Body)

}
