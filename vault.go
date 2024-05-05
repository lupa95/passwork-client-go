package passwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) GetVault(vaultId string) (VaultResponse, error) {
	url := fmt.Sprintf("%s/vaults/%s", c.BaseURL, vaultId)
	method := http.MethodGet
	var responseObject VaultResponse
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

func (c *Client) AddVault(vaultRequest VaultAddRequest) (VaultOperationResponse, error) {
	url := fmt.Sprintf("%s/vaults", c.BaseURL)
	method := http.MethodPost
	var responseObject VaultOperationResponse
	var err error

	body, err := json.Marshal(vaultRequest)
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

	if responseObject.Status != "success" && responseObject.Code != "vaultCreated" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}

func (c *Client) EditVault(vaultId string, request VaultEditRequest) (VaultOperationResponse, error) {
	url := fmt.Sprintf("%s/vaults/%s", c.BaseURL, vaultId)
	method := http.MethodPut
	var responseObject VaultOperationResponse

	body, err := json.Marshal(request)
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

	if responseObject.Status != "success" {
		return responseObject, errors.New(responseObject.Status)
	}

	return responseObject, nil
}

func (c *Client) DeleteVault(vaultId string) (DeleteResponse, error) {
	url := fmt.Sprintf("%s/vaults/%s", c.BaseURL, vaultId)
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
