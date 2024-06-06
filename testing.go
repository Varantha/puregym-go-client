package puregymapi

import "github.com/h2non/gock"

var (
	SuccessfulTokenResponse = TokenResponse{
		AccessToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2F1dGgucHVyZWd5bS5jb20iLCJpYXQiOjE3MTY2NDA3NTksImV4cCI6MTc0ODE3Njc1OSwiYXVkIjoiW1wiaHR0cHM6Ly9hdXRoLnB1cmVneW0uY29tL3Jlc291cmNlc1wiLCBcInBnY2FwaVwiXSIsInN1YiI6Im1lbWJlckBleGFtcGxlLmNvbSIsImNsaWVudF9pZCI6InJvLmNsaWVudCIsImV4dGVybmFsX21lbWJlcl9pZCI6IjEyMzQ1Njc4OSIsImVtYWlsIjoibWVtYmVyQGV4YW1wbGUuY29tIiwibmFtZSI6Ik1yIE1lbWJlciIsIm5vdGUiOiJEaWQgeW91IHJlYWxseSB0YWtlIHRoZSB0aW1lIHRvIGRlY29kZSB0aGlzIGFjY2VzcyB0b2tlbj8_In0.rCPgy-LNefkq4dx653IWCv80oPRnT-kErLq9T_y2eQI",
		Expires_in:  5184000,
		Token_type:  "Bearer",
		Scope:       "pgcapi",
	}
)

func setupMockLogin() {
	gock.New("https://auth.puregym.com").
		Post("/connect/token").
		Reply(200).
		JSON(SuccessfulTokenResponse)
}

func setupDefaultMockRoutes(route string, response interface{}) {
	gock.New("https://capi.puregym.com").
		Get(route).
		MatchHeader("Authorization", "^Bearer\\s.+").
		Reply(200).
		JSON(response)

	gock.New("https://capi.puregym.com").
		Get(route).
		Reply(401)

}
