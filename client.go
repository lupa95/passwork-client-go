package passwork

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL      string
	apiKey       string
	sessionToken string
	HTTPClient   *http.Client
}

type LoginResponse struct {
	Status string
	Data   LoginResponseData
}

type LogoutResponse struct {
	Status string
	Data   string
}

type LoginResponseData struct {
	Token                 string
	RefreshToken          string
	ToeknTtl              int
	RefreshTokenTtl       int
	TokenExpiredAt        int
	RefreshTokenExpiredAt int
	User                  User
}

type User struct {
	Name  string
	Email string
}

func NewClient(baseURL, apiKey string, timeout time.Duration) *Client {
	client := Client{
		BaseURL:      baseURL,
		apiKey:       apiKey,
		sessionToken: "",
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
	return &client
}

// Perform Login Request and set session Token in struct
func (c *Client) Login() error {
	url := fmt.Sprintf("%s/auth/login/%s", c.BaseURL, c.apiKey)

	response, responseCode, err := c.sendRequest(http.MethodPost, url, nil)
	if responseCode != http.StatusOK || err != nil {
		return err
	}

	// Parse JSON into struct
	var responseObject LoginResponse
	err = json.Unmarshal(response, &responseObject)
	if err != nil {
		return err
	}

	if responseObject.Status == "success" {
		c.sessionToken = responseObject.Data.Token
		return nil
	}

	return err
}

func (c *Client) Logout() error {
	url := fmt.Sprintf("%s/auth/logout", c.BaseURL)

	responseData, responseCode, err := c.sendRequest(http.MethodPost, url, nil)
	if responseCode != http.StatusOK || err != nil {
		return err
	}

	// Parse JSON into struct
	var responseObject LogoutResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return err
	}

	if responseObject.Status == "success" && responseObject.Data == "loggedOut" {
		return nil
	}

	return err
}

// Sends HTTP request to URL with method and body
// Returns answer body
func (c *Client) sendRequest(method string, url string, body io.Reader) ([]byte, int, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Passwork-Auth", c.sessionToken)

	// Do HTTP Request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseData, resp.StatusCode, err
	}
	defer resp.Body.Close()

	return responseData, resp.StatusCode, nil
}
