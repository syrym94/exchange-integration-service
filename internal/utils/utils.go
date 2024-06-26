package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func MakeHTTPRequest(method, url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// generateSignature creates the signature for the Bybit API request.
func generateSignature(apiSecret, payload string) string {
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

// MakeBybitAuthenticatedRequest makes an authenticated HTTP request to the Bybit API.
func MakeBybitAuthenticatedRequest(method, endpoint, apiKey, apiSecret string, params map[string]interface{}) ([]byte, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	recvWindow := "20000"

	var payload string
	body := new(bytes.Buffer)

	if method == http.MethodGet {
		query := url.Values{}
		for key, value := range params {
			query.Add(key, fmt.Sprintf("%v", value))
		}
		payload = timestamp + apiKey + recvWindow + query.Encode()
	} else {
		jsonBody, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		payload = timestamp + apiKey + recvWindow + string(jsonBody)
		body = bytes.NewBuffer(jsonBody)
	}

	signature := generateSignature(apiSecret, payload)

	finalURL := endpoint
	if method == http.MethodGet {
		query := url.Values{}
		for key, value := range params {
			query.Add(key, fmt.Sprintf("%v", value))
		}
		finalURL = fmt.Sprintf("%s?%s", endpoint, query.Encode())
	}

	req, err := http.NewRequest(method, finalURL, body)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("x-bapi-api-key", apiKey)
	req.Header.Set("x-bapi-recv-window", recvWindow)
	req.Header.Set("x-bapi-timestamp", timestamp)
	req.Header.Set("x-bapi-sign", signature)
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
