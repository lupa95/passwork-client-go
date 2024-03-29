package passwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetFolder(folderId string) (FolderResponse, error) {
	url := fmt.Sprintf("%s/folders/%s", c.BaseURL, folderId)
	method := http.MethodGet
	var responseObject FolderResponse
	var err error

	resp, err := c.sendRequest(method, url, nil)
	if err != nil {
		return responseObject, err
	}

	FolderResponseData, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(FolderResponseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}

func (c *Client) SearchFolder(request FolderSearchRequest) (FolderSearchResponse, error) {
	url := fmt.Sprintf("%s/folders/search", c.BaseURL)
	method := http.MethodPost
	var responseObject FolderSearchResponse
	var err error

	body, err := json.Marshal(request)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	resp, err := c.sendRequest(method, url, bytes.NewReader(body))
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
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

func (c *Client) AddFolder(folderRequest FolderRequest) (FolderResponse, error) {
	url := fmt.Sprintf("%s/folders", c.BaseURL)
	method := http.MethodPost
	var responseObject FolderResponse
	var err error

	body, err := json.Marshal(folderRequest)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	resp, err := c.sendRequest(method, url, bytes.NewReader(body))
	if resp.StatusCode != http.StatusCreated || err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" && responseObject.Code != "folderCreated" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}

func (c *Client) EditFolder(folderId string, request FolderRequest) (FolderResponse, error) {
	url := fmt.Sprintf("%s/folders/%s", c.BaseURL, folderId)
	method := http.MethodPut
	var responseObject FolderResponse

	body, err := json.Marshal(request)
	if err != nil {
		return responseObject, err
	}

	// HTTP request
	resp, err := c.sendRequest(method, url, bytes.NewReader(body))
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
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

func (c *Client) DeleteFolder(folderId string) (DeleteResponse, error) {
	url := fmt.Sprintf("%s/folders/%s", c.BaseURL, folderId)
	method := http.MethodDelete
	var responseObject DeleteResponse

	// HTTP request
	resp, err := c.sendRequest(method, url, nil)
	if resp.StatusCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
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
