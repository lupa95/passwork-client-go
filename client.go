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

func NewClient(baseURL, apiKey string) *Client {
	client := Client{
		BaseURL:      baseURL,
		apiKey:       apiKey,
		sessionToken: "",
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
	return &client
}

// Perform Login Request and set session Token in struct
func (c *Client) Login() error {
	url := fmt.Sprintf("%s/auth/login/%s", c.BaseURL, c.apiKey)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	// Do HTTP Request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	// Convert Body into byte stream
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Parse JSON into struct
	var responseObject LoginResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return err
	}

	c.sessionToken = responseObject.Data.Token
	return nil
}

// Sends HTTP request to URL with method and body
// Returns answer body
func (c *Client) sendRequest(method string, url string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Passwork-Auth", c.sessionToken)

	// Do HTTP Request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
