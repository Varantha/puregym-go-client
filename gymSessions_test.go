package puregymapi

import (
	"testing"

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
)

func TestGetGymSessions(t *testing.T) {
	defer gock.Off()

	setupMockLogin()

	setupDefaultMockRoutes("/api/v2/gymSessions/gym", GetGymSessionsSuccessResponse)

	client := NewClient(ValidEmail, ValidPin)

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

}
