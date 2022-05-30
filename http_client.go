package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

func getClient(t float64) http.Client {
	timeout := time.Duration(time.Duration(t) * time.Second)
	c := http.Client{
		Timeout: timeout,
	}

	return c
}

func MakeRequest(method string, url string, query map[string]string, body []map[string]string, headers map[string]string) (http.Response, error) {
	url = getUrl(url, query)
	client := getClient(30)

	request, reqError := http.NewRequest(method, url, bytes.NewBuffer(getBody(body)))

	for h, v := range headers {
		request.Header.Set(h, v)
	}

	if reqError != nil {
		log.Fatalln("Request Error:", reqError)
	}

	response, resError := client.Do(request)

	if resError != nil {
		log.Fatalln("Response Error:", resError)
	}

	return *response, nil
}

func getBody(body []map[string]string) []byte {
	var requestBody []byte
	var jsonError error

	if body != nil {
		requestBody, jsonError = json.Marshal(body)

		if jsonError != nil {
			log.Fatalln("JSON Error:", jsonError)

			return nil
		}
	}

	return requestBody
}

func getUrl(uri string, query map[string]string) string {
	u, _ := url.Parse(uri)
	q := u.Query()

	for p, v := range query {
		q.Add(p, v)

	}

	u.RawQuery = q.Encode()

	return u.String()
}
