package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SetApiKey(req *http.Request, apiKey string) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
}

func GenerateJsonPostRequest(url string, body any, apiKey string) (*http.Request, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	SetApiKey(req, apiKey)
	return req, nil
}

func DoRequestAndParseResultAs[T any](req *http.Request) (T, error) {
	data := new(T)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return *data, err
	}

	defer res.Body.Close()
	return ParseResultAs[T](res.Body)
}

func ParseResultAs[T any](reader io.Reader) (T, error) {
	data := new(T)
	body, err := io.ReadAll(reader)
	if err != nil {
		return *data, err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return *data, err
	}
	return *data, nil
}
