package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/xc5216/circle-console-go/internal/setting"
	"github.com/xc5216/circle-console-go/model"
)

type EndPoint struct {
	URL         string
	URLTemplate string
	Method      string
}

func SendRequest[P, Q, R any](endPoint EndPoint, apiKey string, idempotencyKey string, payload *P, query *Q, response *R) (requestID string, err error) {
	requestID = GenerateRequestID()

	queryString := ""
	if query != nil {
		queryString = buildQuery(query)
	}
	endPoint = EndPoint{
		URL:    fmt.Sprintf("%s%s?%s", setting.GetServerURL(), endPoint.URL, queryString),
		Method: endPoint.Method,
	}

	req, err := GenerateRequest(endPoint, idempotencyKey, requestID, payload)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	err = ParseResult(res.Body, response)
	if err != nil && res.StatusCode/100 != 2 {
		err = fmt.Errorf("http status code: %d\n%w", res.StatusCode, err)
		return
	}
	return
}

func GenerateRequest[P any](endPoint EndPoint, idempotencyKey string, requestID string, payload *P) (req *http.Request, err error) {
	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		if idempotencyKey == "" {
			value, err := uuid.NewV4()
			if err != nil {
				return nil, err
			}
			idempotencyKey = value.String()
		}

		jsonData = append(jsonData[0:len(jsonData)-1], []byte(`,"idempotencyKey":"`+idempotencyKey+`"}`)...)
		body = bytes.NewBuffer(jsonData)
	}

	req, err = http.NewRequest(endPoint.Method, endPoint.URL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("X-Request-Id", requestID)
	return req, nil
}

func ParseResult[R any](reader io.Reader, response *R) error {
	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	if string(body) == "" && response == nil {
		return nil
	}

	if response != nil {
		type output struct {
			model.CircleAPIError
			Data R `json:"data"`
		}

		data := output{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			return err
		}
		*response = data.Data
		if data.Code != 0 {
			return data.CircleAPIError
		}
	} else {
		data := model.CircleAPIError{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			return err
		}
		if data.Code != 0 {
			return data
		}
	}
	return nil
}
