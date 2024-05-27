package puregymapi

import (
	"net/url"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var (
	SuccessfulTokenResponse = TokenResponse{
		AccessToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F1dGgucHVyZWd5bS5jb20iLCJpYXQiOjE3MTY2NDA3NTksImV4cCI6MTc0ODE3Njc1OSwiYXVkIjoiW1wiaHR0cHM6Ly9hdXRoLnB1cmVneW0uY29tL3Jlc291cmNlc1wiLCBcInBnY2FwaVwiXSIsInN1YiI6Im1lbWJlckBleGFtcGxlLmNvbSIsImNsaWVudF9pZCI6InJvLmNsaWVudCIsImV4dGVybmFsX21lbWJlcl9pZCI6IjEyMzQ1Njc4OSIsImVtYWlsIjoibWVtYmVyQGV4YW1wbGUuY29tIiwibmFtZSI6Ik1yIE1lbWJlciIsIm5vdGUiOiJEaWQgeW91IHJlYWxseSB0YWtlIHRoZSB0aW1lIHRvIGRlY29kZSB0aGlzIGFjY2VzcyB0b2tlbj8_In0.rCPgy-LNefkq4dx653IWCv80oPRnT-kErLq9T_y2eQI",
		Expires_in:  5184000,
		Token_type:  "Bearer",
		Scope:       "pgcapi",
	}
	UnsuccessfulTokenResponse = map[string]string{
		"error":             "invalid_grant",
		"error_description": "invalid_username_or_password",
	}

	ValidEmail   = "member@example.com"
	ValidPin     = "12341234"
	InvalidEmail = "test@example.com"
	InvalidPin   = "43214321"
)

func setupMockLogin() {
	gock.New("https://auth.puregym.com").
		Post("/connect/token").
		Reply(200).
		JSON(SuccessfulTokenResponse)
}

func TestNewClient(t *testing.T) {
	defer gock.Off()

	setupMockLogin()

	client := NewClient(ValidEmail, ValidPin)

	t.Run("username set correctly", func(t *testing.T) {
		assert.Equal(t, ValidEmail, client.username)
	})

	t.Run("pin set correctly", func(t *testing.T) {
		assert.Equal(t, ValidPin, client.password)
	})
}

func TestLogin(t *testing.T) {
	defer gock.Off()

	validBodyValues := url.Values{
		"client_id":  {"ro.client"},
		"scope":      {"pgcapi"},
		"grant_type": {"password"},
		"username":   {ValidEmail},
		"password":   {ValidPin},
	}

	validBody := strings.NewReader(validBodyValues.Encode())

	gock.New("https://auth.puregym.com").
		Post("/connect/token").
		MatchType("url").
		Body(validBody).
		Reply(200).
		JSON(SuccessfulTokenResponse)

	gock.New("https://auth.puregym.com").
		Post("/connect/token").
		MatchType("url").
		Reply(400).
		JSON(UnsuccessfulTokenResponse)

	testCases := []struct {
		name        string
		errExpected bool
		email       string
		pin         string
	}{
		{"valid login", false, ValidEmail, ValidPin},
		{"invalid login", true, InvalidEmail, InvalidPin},
		{"invalid pin", true, ValidEmail, InvalidPin},
		{"invalid email", true, InvalidEmail, ValidPin},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewClient(tc.email, tc.pin)

			err := client.Login()

			if tc.errExpected && err == nil {
				t.Error("Expected an error, got nil")
			}
		})
	}
}
