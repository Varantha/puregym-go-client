package puregymapi

import (
	"errors"
	"net/url"
	"time"
)

type GetGymSessionsGymResponse struct {
	TotalPeopleInGym     int    `json:"TotalPeopleInGym"`
	TotalPeopleInClasses int    `json:"TotalPeopleInClasses"`
	AttendanceTime       string `json:"AttendanceTime"`
	LastRefreshed        string `json:"LastRefreshed"`
	MaximumCapacity      int    `json:"MaximumCapacity"`
}

type GetMemberGymSessionsResponse struct {
	Summary    Summary    `json:"Summary"`
	Visits     []Visit    `json:"Visits"`
	Activities []Activity `json:"Activities"`
}

type Summary struct {
	Total    MemberActivity `json:"Total"`
	ThisWeek MemberActivity `json:"ThisWeek"`
}

type MemberActivity struct {
	Activities int `json:"Activities"`
	Visits     int `json:"Visits"`
	Duration   int `json:"Duration"`
}

type Visit struct {
	IsDurationEstimated bool    `json:"IsDurationEstimated"`
	Gym                 Gym     `json:"Gym"`
	StartTime           string  `json:"StartTime"`
	Duration            int     `json:"Duration"`
	Name                *string `json:"Name"`
}

type Activity struct {
	// Define the fields of the Activity struct here
	// I've never done an activity so I don't know what the fields are
}

func (c *Client) GetGymSessions(gymId string) (*GetGymSessionsGymResponse, error) {
	route := "gymSessions/gym"

	if gymId == "" {
		return nil, errors.New("gymId cannot be empty")
	}

	params := url.Values{}
	params.Add("gymId", gymId)

	var gymSessions GetGymSessionsGymResponse
	err := c.sendRequest("GET", route, nil, params, &gymSessions)
	if err != nil {
		return nil, err
	}

	return &gymSessions, nil
}

func (c *Client) GetMemberGymSessions() (*GetMemberGymSessionsResponse, error) {
	now := time.Now()
	fromDate := now.AddDate(0, -3, 0)

	return c.GetMemberGymSessionsBetween(&fromDate, &now)
}

func (c *Client) GetMemberGymSessionsBetween(fromDate *time.Time, toDate *time.Time) (*GetMemberGymSessionsResponse, error) {
	route := "gymSessions/member"

	if fromDate == nil || toDate == nil {
		return nil, errors.New("fromDate and toDate cannot be nil")
	}

	params := url.Values{}

	params.Add("fromDate", fromDate.Format("2006-01-02T15:04:05"))

	params.Add("toDate", toDate.Format("2006-01-02T15:04:05"))

	var gymSessions GetMemberGymSessionsResponse
	err := c.sendRequest("GET", route, nil, params, &gymSessions)
	if err != nil {
		return nil, err
	}

	return &gymSessions, nil
}
