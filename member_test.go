package puregymapi

import (
	"testing"
	"time"

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
		HomeGym: Gym{
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

	GetMemberQRCodeSuccessResponse = GetMemberQRCodeResponse{
		QrCode:    "exerp:checkin:xxxyyyzzz",
		RefreshAt: parseTime("2024-05-20T18:34:10.1930835Z"),
		ExpiresAt: parseTime("2024-05-27T18:28:10.1930835Z"),
		RefreshIn: "0:01:00",
		ExpiresIn: "167:55:00",
	}

	GetMembershipSuccessResponse = GetMembershipResponse{
		Name:              "PremiumMultiAccess",
		Level:             "PremiumMultiAccess",
		StartDate:         time.Time{},
		EndDate:           time.Time{},
		PaymentDayOfMonth: 9,
		HoursOfAccess:     time.Time{},
		IncludedGyms: []Gym{
			{
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
			}},
		FreezeDetails: "",
	}
)

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339Nano, timeStr)
	return t
}

func TestGetMember(t *testing.T) {
	defer gock.Off()

	setupMockLogin()

	setupDefaultMockRoutes("/api/v2/member", GetMemberSuccessResponse)

	client := NewClient(ValidEmail, ValidPin)

	t.Run("unauthenticated request fails", func(t *testing.T) {
		_, err := client.GetMember()
		assert.Error(t, err)
	})

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

func TestGetMemberQRCode(t *testing.T) {
	defer gock.Off()

	setupMockLogin()
	setupDefaultMockRoutes("/api/v2/member/qrcode", GetMemberQRCodeSuccessResponse)

	client := NewClient(ValidEmail, ValidPin)

	t.Run("unauthenticated request fails", func(t *testing.T) {
		_, err := client.GetMemberQRCode()
		assert.Error(t, err)
	})

	client.Login()

	t.Run("returns member QR Code", func(t *testing.T) {
		// Arrange
		expectedResponse := &GetMemberQRCodeSuccessResponse
		// Act
		response, err := client.GetMemberQRCode()
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

}

func TestGetMembership(t *testing.T) {
	defer gock.Off()

	setupMockLogin()

	setupDefaultMockRoutes("/api/v2/member/membership", GetMembershipSuccessResponse)

	client := NewClient(ValidEmail, ValidPin)

	t.Run("unauthenticated request fails", func(t *testing.T) {
		_, err := client.GetMembership()
		assert.Error(t, err)
	})

	client.Login()

	t.Run("returns membership details", func(t *testing.T) {
		// Arrange
		expectedResponse := &GetMembershipSuccessResponse
		// Act
		response, err := client.GetMembership()
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

}
