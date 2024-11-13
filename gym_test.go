package puregymapi

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var (
	GetGymSuccessResponse = Gym{
		ID:     1,
		Name:   "Manchester Spinningfields",
		Status: "Open",
		Location: Location{
			Address: Address{
				Line1:    "3 Hardman Street",
				Line2:    "Spinningfields",
				Line3:    nil,
				Town:     "MANCHESTER",
				County:   "GREATER MANCHESTER",
				Province: nil,
				Postcode: "M3 3HF",
				Country:  "GB",
			},
			GeoLocation: GeoLocation{
				Longitude: -2.2504,
				Latitude:  53.4798,
			},
		},
		GymAccess: GymAccess{
			AccessOptions: AccessOptions{
				PinAccess:    true,
				QrCodeAccess: true,
			},
			OpeningHours: OpeningHours{
				IsAlwaysOpen: true,
				OpeningHours: []interface{}{},
			},
			StandardOpeningTimes: []StandardOpeningTime{
				{"Sunday", "0:00:00", "23:59:00", false},
				{"Monday", "0:00:00", "23:59:00", false},
				{"Tuesday", "0:00:00", "23:59:00", false},
				{"Wednesday", "0:00:00", "23:59:00", false},
				{"Thursday", "0:00:00", "23:59:00", false},
				{"Friday", "0:00:00", "23:59:00", false},
				{"Saturday", "0:00:00", "23:59:00", false},
			},
			ReopenDate: "2021-04-12T00:00:00+01 Europe/London",
		},
		ContactInfo: ContactDetails{
			PhoneNumber:  "+44 3444770005",
			EmailAddress: "info-manchester@puregym.com",
		},
		TimeZone: "Europe/London",
	}

	GetGymsSuccessResponse = []Gym{
		GetGymSuccessResponse,
		Gym{
			ID:     2,
			Name:   "Wolverhampton Bentley Bridge",
			Status: "Open",
			Location: Location{
				Address: Address{
					Line1:    "Bentley Bridge Retail Park",
					Line2:    "Bentley Bridge Way",
					Line3:    "Wednesfield",
					Town:     "WOLVERHAMPTON",
					County:   "WEST MIDLANDS",
					Province: "Wednesfield",
					Postcode: "WV11 1BP",
					Country:  "GB",
				},
				GeoLocation: GeoLocation{
					Longitude: -2.0879,
					Latitude:  52.596,
				},
			},
			GymAccess: GymAccess{
				AccessOptions: AccessOptions{
					PinAccess:    true,
					QrCodeAccess: true,
				},
				OpeningHours: OpeningHours{
					IsAlwaysOpen: true,
					OpeningHours: []interface{}{},
				},
				StandardOpeningTimes: []StandardOpeningTime{
					{"Sunday", "0:00:00", "23:59:00", false},
					{"Monday", "0:00:00", "23:59:00", false},
					{"Tuesday", "0:00:00", "23:59:00", false},
					{"Wednesday", "0:00:00", "23:59:00", false},
					{"Thursday", "0:00:00", "23:59:00", false},
					{"Friday", "0:00:00", "23:59:00", false},
					{"Saturday", "0:00:00", "23:59:00", false},
				},
				ReopenDate: "2021-04-12T00:00:00+01 Europe/London",
			},
			ContactInfo: ContactDetails{
				PhoneNumber:  "+44 3444770005",
				EmailAddress: "Info-Wolverhampton@puregym.com",
			},
			TimeZone: "Europe/London",
		},
	}
)

func TestGetGyms(t *testing.T) {
	defer gock.Off()

	route := "/api/v2/gyms"
	response := GetGymsSuccessResponse

	setupMockLogin()

	gock.New("https://capi.puregym.com").
		Get(route).
		MatchHeader("Authorization", "^Bearer\\s.+").
		Reply(200).
		JSON(response)

	gock.New("https://capi.puregym.com").
		Get(route).
		Reply(401)

	client := NewClient(validEmail, validPin)

	t.Run("unauthenticated request fails", func(t *testing.T) {
		_, err := client.GetGyms()
		assert.Error(t, err)
	})

	client.Login()

	t.Run("returns gym details", func(t *testing.T) {
		// Arrange
		expectedResponse := &GetGymsSuccessResponse
		// Act
		response, err := client.GetGyms()
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})
}

func TestGetGym(t *testing.T) {
	defer gock.Off()

	route := "/api/v2/gyms"
	response := GetGymsSuccessResponse

	setupMockLogin()

	gock.New("https://capi.puregym.com").
		Get(route).
		MatchHeader("Authorization", "^Bearer\\s.+").
		Reply(200).
		JSON(response)

	gock.New("https://capi.puregym.com").
		Get(route).
		Reply(401)

	client := NewClient(validEmail, validPin)

	t.Run("unauthenticated request fails", func(t *testing.T) {
		_, err := client.GetGym(1)
		assert.Error(t, err)
	})

	client.Login()

	t.Run("returns gym details", func(t *testing.T) {
		// Arrange
		expectedResponse := &GetGymSuccessResponse
		// Act
		response, err := client.GetGym(1)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})
}
