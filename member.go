package puregymapi

type ID struct {
	ExternalID string `json:"ExternalId"`
	CompoundID string `json:"CompoundId"`
}

type ContactDetails struct {
	PhoneNumber  string `json:"PhoneNumber"`
	EmailAddress string `json:"EmailAddress"`
}

type Address struct {
	Line1    interface{} `json:"Line1"`
	Line2    interface{} `json:"Line2"`
	Line3    interface{} `json:"Line3"`
	Town     string      `json:"Town"`
	County   interface{} `json:"County"`
	Province interface{} `json:"Province"`
	Postcode string      `json:"Postcode"`
	Country  string      `json:"Country"`
}

type PersonalDetails struct {
	FirstName      string         `json:"FirstName"`
	LastName       string         `json:"LastName"`
	DateOfBirth    string         `json:"DateOfBirth"`
	ContactDetails ContactDetails `json:"ContactDetails"`
	Address        Address        `json:"Address"`
}

type GeoLocation struct {
	Longitude float64 `json:"Longitude"`
	Latitude  float64 `json:"Latitude"`
}

type Location struct {
	Address     Address     `json:"Address"`
	GeoLocation GeoLocation `json:"GeoLocation"`
}

type OpeningHours struct {
	IsAlwaysOpen bool          `json:"IsAlwaysOpen"`
	OpeningHours []interface{} `json:"OpeningHours"`
}

type StandardOpeningTime struct {
	DayOfWeek string `json:"DayOfWeek"`
	StartTime string `json:"StartTime"`
	EndTime   string `json:"EndTime"`
	IsHoliday bool   `json:"IsHoliday"`
}

type GymAccess struct {
	AccessOptions        interface{}           `json:"AccessOptions"`
	OpeningHours         OpeningHours          `json:"OpeningHours"`
	StandardOpeningTimes []StandardOpeningTime `json:"StandardOpeningTimes"`
	ReopenDate           string                `json:"ReopenDate"`
}

type HomeGym struct {
	ID          int            `json:"Id"`
	Name        string         `json:"Name"`
	Status      string         `json:"Status"`
	Location    Location       `json:"Location"`
	GymAccess   GymAccess      `json:"GymAccess"`
	ContactInfo ContactDetails `json:"ContactInfo"`
	TimeZone    string         `json:"TimeZone"`
}

type GetMemberResponse struct {
	ID              ID              `json:"Id"`
	PersonalDetails PersonalDetails `json:"PersonalDetails"`
	HomeGym         HomeGym         `json:"HomeGym"`
	GymAccessPin    string          `json:"GymAccessPin"`
	SuspendedReason string          `json:"SuspendedReason"`
	MemberStatus    string          `json:"MemberStatus"`
}

// GetMember gets information for the currently logged in member including Personal details, Home Gym details, and membership status
func (c *Client) GetMember() (*GetMemberResponse, error) {
	route := "member"

	var member GetMemberResponse
	err := c.sendRequest("GET", route, nil, &member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}
