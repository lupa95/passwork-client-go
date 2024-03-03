package passwork

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetVault(vaultId string) (VaultResponse, error) {
	url := fmt.Sprintf("%s/vaults/%s", c.BaseURL, vaultId)
	method := http.MethodGet
	var responseObject VaultResponse
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
		return responseObject, err
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

	if responseObject.Status != "success" && responseObject.Code != "vaultCreated" {
		return responseObject, err
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
		return responseObject, err
	}

	return responseObject, nil
}

func (c *Client) DeleteVault(vaultId string) (DeleteResponse, error) {
	url := fmt.Sprintf("%s/vaults/%s", c.BaseURL, vaultId)
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
		return responseObject, err
	}

	return responseObject, nil
}
