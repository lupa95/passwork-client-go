package passwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) GetFolder(folderId string) (FolderResponse, error) {
	url := fmt.Sprintf("%s/folders/%s", c.BaseURL, folderId)
	method := http.MethodGet
	var responseObject FolderResponse
	var err error

	response, responseCode, err := c.sendRequest(method, url, nil)
	if responseCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(response, &responseObject)
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
	response, responseCode, err := c.sendRequest(method, url, bytes.NewReader(body))
	if responseCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Parse JSON into struct (this returns a list of results)
	err = json.Unmarshal(response, &responseObject)
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
	response, responseCode, err := c.sendRequest(method, url, bytes.NewReader(body))
	if responseCode != http.StatusCreated || err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	err = json.Unmarshal(response, &responseObject)
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
	response, responseCode, err := c.sendRequest(method, url, bytes.NewReader(body))
	if responseCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	err = json.Unmarshal(response, &responseObject)
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
	response, responseCode, err := c.sendRequest(method, url, nil)
	if responseCode != http.StatusOK || err != nil {
		return responseObject, err
	}

	// Parse JSON into struct
	err = json.Unmarshal(response, &responseObject)
	if err != nil {
		return responseObject, err
	}

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}
