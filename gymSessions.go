package puregymapi

import "net/url"

type GetGymSessionsGymResponse struct {
	TotalPeopleInGym     int    `json:"TotalPeopleInGym"`
	TotalPeopleInClasses int    `json:"TotalPeopleInClasses"`
	AttendanceTime       string `json:"AttendanceTime"`
	LastRefreshed        string `json:"LastRefreshed"`
	MaximumCapacity      int    `json:"MaximumCapacity"`
}

func (c *Client) GetGymSessions(gymId string) (*GetGymSessionsGymResponse, error) {
	route := "gymSessions/gym"

	params := url.Values{}
	params.Add("gymId", gymId)

	var gymSessions GetGymSessionsGymResponse
	err := c.sendRequest("GET", route, nil, params, &gymSessions)
	if err != nil {
		return nil, err
	}

	return &gymSessions, nil
}
