package puregymapi

import (
	"net/url"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var (
	UnsuccessfulTokenResponse = map[string]string{
		"error":             "invalid_grant",
		"error_description": "invalid_username_or_password",
	}
)

func TestNewClient(t *testing.T) {
	defer gock.Off()

	setupMockLogin()

	client := NewClient(validEmail, validPin)

	t.Run("username set correctly", func(t *testing.T) {
		assert.Equal(t, validEmail, client.username)
	})

	t.Run("pin set correctly", func(t *testing.T) {
		assert.Equal(t, validPin, client.password)
	})
}

func TestLogin(t *testing.T) {
	defer gock.Off()

	validBodyValues := url.Values{
		"client_id":  {"ro.client"},
		"scope":      {"pgcapi"},
		"grant_type": {"password"},
		"username":   {validEmail},
		"password":   {validPin},
	}

	validBody := strings.NewReader(validBodyValues.Encode())

	gock.New("https://auth.puregym.com").
		Post("/connect/token").
		MatchType("url").
		Body(validBody).
		Reply(200).
		JSON(successfulTokenResponse)

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
		{"valid login", false, validEmail, validPin},
		{"invalid login", true, invalidEmail, invalidPin},
		{"invalid pin", true, validEmail, invalidPin},
		{"invalid email", true, invalidEmail, validPin},
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
