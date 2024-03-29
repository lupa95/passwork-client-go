package passwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Get a password by ID
func (c *Client) GetPassword(pwId string) (PasswordResponse, error) {
	url := fmt.Sprintf("%s/passwords/%s", c.BaseURL, pwId)
	method := http.MethodGet
	var responseObject PasswordResponse
	var err error

	// HTTP request
	resp, err := c.sendRequest(method, url, nil)
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	PasswordResponseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	err = json.Unmarshal(PasswordResponseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}

// Search for password by name
func (c *Client) SearchPassword(request PasswordSearchRequest) (PasswordSearchResponse, error) {
	url := fmt.Sprintf("%s/passwords/search", c.BaseURL)
	method := http.MethodPost
	var responseObject PasswordSearchResponse
	var err error

	body, err := json.Marshal(request)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	resp, err := c.sendRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Parse JSON into struct (this returns a list of results)
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}

func (c *Client) AddPassword(pwRequest PasswordRequest) (PasswordResponse, error) {
	url := fmt.Sprintf("%s/passwords", c.BaseURL)
	method := http.MethodPost
	var responseObject PasswordResponse
	var err error

	body, err := json.Marshal(pwRequest)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	resp, err := c.sendRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}

func (c *Client) EditPassword(pwId string, request PasswordRequest) (PasswordResponse, error) {
	url := fmt.Sprintf("%s/passwords/%s", c.BaseURL, pwId)
	method := http.MethodPut
	var responseObject PasswordResponse

	body, err := json.Marshal(request)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	resp, err := c.sendRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}

func (c *Client) DeletePassword(pwId string) (DeleteResponse, error) {
	url := fmt.Sprintf("%s/passwords/%s", c.BaseURL, pwId)
	method := http.MethodDelete
	var responseObject DeleteResponse

	// HTTP request
	resp, err := c.sendRequest(method, url, nil)
	if err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}
