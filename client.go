package puregymapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BaseURLV2 = "https://capi.puregym.com/api/v2/"
	AuthUrl   = "https://auth.puregym.com/connect/token"
)

type (
	Client struct {
		baseURL    string
		authHost   string
		httpClient *http.Client
		username   string
		password   string
		token      string
	}
	TokenResponse struct {
		AccessToken string `json:"access_token"`
		Expires_in  int    `json:"expires_in"`
		Token_type  string `json:"token_type"`
		Scope       string `json:"scope"`
	}
)

func NewClient(emailAddress string, pin string) *Client {
	return &Client{
		baseURL:  BaseURLV2,
		authHost: AuthUrl,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		username: emailAddress,
		password: pin,
		token:    "",
	}
}

func (c *Client) Login() error {
	// Create the form data
	formData := url.Values{}
	formData.Set("client_id", "ro.client")
	formData.Set("scope", "pgcapi")
	formData.Set("grant_type", "password")
	formData.Set("username", c.username)
	formData.Set("password", c.password)

	// Convert form data to a reader
	body := strings.NewReader(formData.Encode())

	req, err := http.NewRequest(http.MethodPost, c.authHost, body)

	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "PureGym/1523 CFNetwork/1312 Darwin/21.0.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error calling request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("unexpected status: %s, Your username or PIN may be incorrect", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var tokenResponse TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return fmt.Errorf("can't decode response: %v", err)
	}

	c.token = tokenResponse.AccessToken
	return nil
}
