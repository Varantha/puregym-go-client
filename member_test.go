package puregymapi

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var (
	GetMemberSuccessResponse = GetMemberResponse{
		ID: ID{
			ExternalID: "dummyExternalID",
			CompoundID: "dummyCompoundID",
		},
		PersonalDetails: PersonalDetails{
			FirstName:   "DummyFirstName",
			LastName:    "DummyLastName",
			DateOfBirth: "2000-01-01",
			ContactDetails: ContactDetails{
				PhoneNumber:  "1234567890",
				EmailAddress: "dummy@example.com",
			},
			Address: Address{
				Line1:    "Dummy Line1",
				Line2:    "Dummy Line2",
				Line3:    "Dummy Line3",
				Town:     "Dummy Town",
				County:   "Dummy County",
				Province: "Dummy Province",
				Postcode: "Dummy Postcode",
				Country:  "Dummy Country",
			},
		},
		HomeGym: HomeGym{
			ID:     1,
			Name:   "Dummy Gym Name",
			Status: "Open",
			Location: Location{
				Address: Address{
					Line1:    "Dummy Gym Line1",
					Line2:    "Dummy Gym Line2",
					Line3:    "Dummy Gym Line3",
					Town:     "Dummy Gym Town",
					County:   "Dummy Gym County",
					Province: "Dummy Gym Province",
					Postcode: "Dummy Gym Postcode",
					Country:  "Dummy Gym Country",
				},
				GeoLocation: GeoLocation{
					Longitude: 123.456,
					Latitude:  78.910,
				},
			},
			GymAccess: GymAccess{
				AccessOptions: "Dummy Access Options",
				OpeningHours: OpeningHours{
					IsAlwaysOpen: true,
					OpeningHours: []interface{}{"Dummy Opening Hours"},
				},
				StandardOpeningTimes: []StandardOpeningTime{
					{
						DayOfWeek: "Monday",
						StartTime: "09:00",
						EndTime:   "17:00",
						IsHoliday: false,
					},
				},
				ReopenDate: "2022-01-01",
			},
			ContactInfo: ContactDetails{
				PhoneNumber:  "1234567890",
				EmailAddress: "dummygym@example.com",
			},
			TimeZone: "Dummy TimeZone",
		},
		GymAccessPin:    "123456",
		SuspendedReason: "None",
		MemberStatus:    "Active",
	}
)

func TestGetMember(t *testing.T) {
	defer gock.Off()

	gock.New("https://capi.puregym.com").
		Get("/api/v2/member").
		Reply(200).
		JSON(GetMemberSuccessResponse)

	client := NewClient(ValidEmail, ValidPin)
	client.Login()

	t.Run("returns member information", func(t *testing.T) {
		// Arrange
		expectedResponse := &GetMemberSuccessResponse
		// Act
		response, err := client.GetMember()
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})
}
