package puregymapi

import (
	"testing"
	"time"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var (
	GetGymSessionsSuccessResponse = GetGymSessionsGymResponse{
		TotalPeopleInGym:     52,
		TotalPeopleInClasses: 0,
		AttendanceTime:       "2024-06-06T20:42:08.105643+01 Europe/London",
		LastRefreshed:        "2024-06-06T19:42:08.105643Z",
		MaximumCapacity:      0,
	}
	GetMemberGymSessionsSuccessResponse = GetMemberGymSessionsResponse{
		Summary: Summary{
			Total: MemberActivity{
				Activities: 0,
				Visits:     2,
				Duration:   61,
			},
			ThisWeek: MemberActivity{
				Activities: 0,
				Visits:     0,
				Duration:   0,
			},
		},
		Visits: []Visit{
			{
				IsDurationEstimated: false,
				Gym: Gym{
					ID:          236,
					Name:        "London Camden",
					Status:      "Blocked",
					Location:    Location{},
					GymAccess:   GymAccess{},
					ContactInfo: ContactDetails{},
					TimeZone:    "",
				},
				StartTime: "2024-05-12T09:16:00",
				Duration:  21,
				Name:      nil,
			},
			{
				IsDurationEstimated: false,
				Gym: Gym{
					ID:          236,
					Name:        "London Camden",
					Status:      "Blocked",
					Location:    Location{},
					GymAccess:   GymAccess{},
					ContactInfo: ContactDetails{},
					TimeZone:    "",
				},
				StartTime: "2024-03-10T10:26:00",
				Duration:  40,
				Name:      nil,
			},
		},
		Activities: []Activity{},
	}
)

func TestGetGymSessions(t *testing.T) {
	defer gock.Off()

	route := "/api/v2/gymSessions/gym"
	response := GetGymSessionsSuccessResponse

	setupMockLogin()

	gock.New("https://capi.puregym.com").
		Get(route).
		MatchHeader("Authorization", "^Bearer\\s.+").
		MatchParam("gymId", "^[0-9]+$").
		Reply(200).
		JSON(response)

	gock.New("https://capi.puregym.com").
		Get(route).
		Reply(401)

	client := NewClient(validEmail, validPin)

	t.Run("unauthenticated request fails", func(t *testing.T) {
		_, err := client.GetGymSessions("236")
		assert.Error(t, err)
	})

	client.Login()

	t.Run("returns membership details", func(t *testing.T) {
		// Arrange
		expectedResponse := &GetGymSessionsSuccessResponse
		// Act
		response, err := client.GetGymSessions("236")
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("fail with no gymId", func(t *testing.T) {
		_, err := client.GetGymSessions("")
		assert.Error(t, err)
	})

}

func TestGetMemberGymSessionsBetween(t *testing.T) {
	defer gock.Off()

	route := "/api/v2/gymSessions/member"
	response := GetMemberGymSessionsSuccessResponse

	setupMockLogin()

	gock.New("https://capi.puregym.com").
		Get(route).
		MatchHeader("Authorization", "^Bearer\\s.+").
		MatchParam("fromDate", "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$").
		MatchParam("toDate", "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$").
		Reply(200).
		JSON(response)

	gock.New("https://capi.puregym.com").
		Get(route).
		Reply(401)

	client := NewClient(validEmail, validPin)

	fromDate, _ := time.Parse("2006-01-02T15:04:05", "2024-01-01T00:00:00")
	toDate, _ := time.Parse("2006-01-02T15:04:05", "2024-12-31T23:59:59")

	t.Run("unauthenticated request fails", func(t *testing.T) {
		_, err := client.GetMemberGymSessionsBetween(&fromDate, &toDate)
		assert.Error(t, err)
	})

	client.Login()

	t.Run("returns membership details", func(t *testing.T) {
		// Arrange
		expectedResponse := &GetMemberGymSessionsSuccessResponse
		// Act
		response, err := client.GetMemberGymSessionsBetween(&fromDate, &toDate)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("fail with no gymId", func(t *testing.T) {
		_, err := client.GetGymSessions("")
		assert.Error(t, err)
	})
}

func TestGetMemberGymSessions(t *testing.T) {
	defer gock.Off()

	route := "/api/v2/gymSessions/member"
	response := GetMemberGymSessionsSuccessResponse

	setupMockLogin()

	gock.New("https://capi.puregym.com").
		Get(route).
		MatchHeader("Authorization", "^Bearer\\s.+").
		MatchParam("fromDate", "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$").
		MatchParam("toDate", "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$").
		Reply(200).
		JSON(response)

	gock.New("https://capi.puregym.com").
		Get(route).
		Reply(401)

	client := NewClient(validEmail, validPin)

	t.Run("unauthenticated request fails", func(t *testing.T) {
		_, err := client.GetMemberGymSessions()
		assert.Error(t, err)
	})

	client.Login()

	t.Run("returns membership details", func(t *testing.T) {
		// Arrange
		expectedResponse := &GetMemberGymSessionsSuccessResponse
		// Act
		response, err := client.GetMemberGymSessions()
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})
}
