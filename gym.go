package puregymapi

import "fmt"

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
	AccessOptions        AccessOptions         `json:"AccessOptions"`
	OpeningHours         OpeningHours          `json:"OpeningHours"`
	StandardOpeningTimes []StandardOpeningTime `json:"StandardOpeningTimes"`
	ReopenDate           string                `json:"ReopenDate"`
}

type AccessOptions struct {
	PinAccess    bool `json:"PinAccess"`
	QrCodeAccess bool `json:"QrCodeAccess"`
}

// Gym represents a gym with its details.
type Gym struct {
	ID          int            `json:"Id"`
	Name        string         `json:"Name"`
	Status      string         `json:"Status"`
	Location    Location       `json:"Location"`
	GymAccess   GymAccess      `json:"GymAccess"`
	ContactInfo ContactDetails `json:"ContactInfo"`
	TimeZone    string         `json:"TimeZone"`
}

// GetGyms retrieves a list of gyms from the API.
// It sends a GET request to the "gyms" endpoint and returns a slice of Gym structs.
func (c *Client) GetGyms() (*[]Gym, error) {
	route := "gyms"

	var gyms []Gym
	err := c.sendRequest("GET", route, nil, nil, &gyms)
	if err != nil {
		return nil, err
	}

	return &gyms, nil
}

// GetGym retrieves a specific gym by its ID.
// It first fetches the list of gyms and then searches for the gym with the given ID.
func (c *Client) GetGym(gymId int) (*Gym, error) {
	gyms, err := c.GetGyms()
	if err != nil {
		return nil, err
	}

	for _, gym := range *gyms {
		if gym.ID == gymId {
			return &gym, nil
		}
	}

	return nil, fmt.Errorf("gym with ID %d not found", gymId)
}
