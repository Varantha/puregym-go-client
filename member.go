package puregymapi

import "time"

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

type GetMemberResponse struct {
	ID              ID              `json:"Id"`
	PersonalDetails PersonalDetails `json:"PersonalDetails"`
	HomeGym         Gym             `json:"HomeGym"`
	GymAccessPin    string          `json:"GymAccessPin"`
	SuspendedReason string          `json:"SuspendedReason"`
	MemberStatus    string          `json:"MemberStatus"`
}

type GetMemberQRCodeResponse struct {
	QrCode    string    `json:"QrCode"`
	RefreshAt time.Time `json:"RefreshAt"`
	ExpiresAt time.Time `json:"ExpiresAt"`
	RefreshIn string    `json:"RefreshIn"`
	ExpiresIn string    `json:"ExpiresIn"`
}

type GetMembershipResponse struct {
	Name              string      `json:"Name"`
	Level             string      `json:"Level"`
	StartDate         interface{} `json:"StartDate"`
	EndDate           interface{} `json:"EndDate"`
	PaymentDayOfMonth int         `json:"PaymentDayOfMonth"`
	HoursOfAccess     interface{} `json:"MemberStatus"`
	IncludedGyms      []Gym       `json:"IncludedGyms"`
	FreezeDetails     string      `json:"FreezeDetails"`
}

// GetMember retrieves the member details from the API.
// It sends a GET request to the "member" endpoint and returns a GetMemberResponse struct.
func (c *Client) GetMember() (*GetMemberResponse, error) {
	route := "member"

	var member GetMemberResponse
	err := c.sendRequest("GET", route, nil, nil, &member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

// GetMemberQRCode retrieves the member's QR code from the API.
// It sends a GET request to the "member/qrcode" endpoint and returns a GetMemberQRCodeResponse struct.
func (c *Client) GetMemberQRCode() (*GetMemberQRCodeResponse, error) {
	route := "member/qrcode"

	var qrCode GetMemberQRCodeResponse
	err := c.sendRequest("GET", route, nil, nil, &qrCode)
	if err != nil {
		return nil, err
	}

	return &qrCode, nil
}

// GetMembership retrieves the membership details from the API.
// It sends a GET request to the "member/membership" endpoint and returns a GetMembershipResponse struct.
func (c *Client) GetMembership() (*GetMembershipResponse, error) {
	route := "member/membership"

	var membership GetMembershipResponse
	err := c.sendRequest("GET", route, nil, nil, &membership)
	if err != nil {
		return nil, err
	}

	return &membership, nil
}
