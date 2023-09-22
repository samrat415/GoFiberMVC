package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// CurlService provides methods for making GET and POST requests
type CurlService struct{}

// Get makes a GET request to the specified URL with optional headers
func (s *CurlService) Get(url string, headers map[string]string) (map[string]interface{}, error) {
	// Create a custom HTTP client with insecure TLS configuration
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	// Set headers if provided
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var parsedResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&parsedResponse)
	if err != nil {
		return nil, err
	}
	return parsedResponse, nil
}

// Post Request
func (s *CurlService) Post(url string, fields url.Values, headers map[string]string) (map[string]interface{}, error) {
	reqBody := bytes.NewBufferString(fields.Encode())
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create a custom HTTP client with insecure TLS configuration
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var parsedResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&parsedResponse)
	if err != nil {
		return nil, err
	}

	return parsedResponse, nil
}

func (s *CurlService) PostJSON(url string, payload interface{}, headers map[string]string) (map[string]interface{}, error) {

	// Marshal the JSON payload
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	//if resp.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("unexpected response status code: %v", resp.StatusCode)
	//}
	var parsedResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}
	return parsedResponse, nil
}
