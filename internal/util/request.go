package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/gofrs/uuid"
)

func SetApiKey(req *http.Request, apiKey string) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
}

func SetRequestID(req *http.Request, requestID string) {
	req.Header.Add("X-Request-Id", requestID)
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
	if apiKey != "" {
		SetApiKey(req, apiKey)
	}
	return req, nil
}

func GenerateGetRequest(url string, queryData any, apiKey string) (*http.Request, error) {
	queryString := buildQuery(queryData)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", url, queryString), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	if apiKey != "" {
		SetApiKey(req, apiKey)
	}
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

	if string(body) == "" {
		return *data, nil
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return *data, err
	}
	return *data, nil
}

func GenerateRequestID() string {
	id, _ := uuid.NewV1()
	return id.String()
}

func buildQuery(params interface{}) string {
	v := reflect.ValueOf(params)
	if v.Kind() != reflect.Struct {
		return ""
	}

	var queryParts []string
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("url")
		if tag != "" {
			value := v.Field(i).Interface()
			queryParts = append(queryParts, fmt.Sprintf("%s=%v", tag, value))
		}
	}

	return strings.Join(queryParts, "&")
}

type CircleRequest struct {
	*http.Request
	requestID      string
	idempotencyKey string
}

func (r *CircleRequest) GetRequestID() string {
	return r.requestID
}

func (r *CircleRequest) GetIdempotencyKey() string {
	return r.idempotencyKey
}

func NewCircleRequest(req *http.Request) *CircleRequest {
	return &CircleRequest{
		requestID:      GenerateRequestID(),
		idempotencyKey: uuid.Must(uuid.NewV4()).String(),
		Request:        req,
	}
}
