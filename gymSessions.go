package puregymapi

import (
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

// HistoryOptions holds the optional parameters for fetching user history
type HistoryOptions struct {
	StartDate *time.Time
	EndDate   *time.Time
}

// Option is a function that sets some option on the HistoryOptions
type Option func(*HistoryOptions)

// WithStartDate sets the start date option
func WithStartDate(date time.Time) Option {
	return func(opts *HistoryOptions) {
		opts.StartDate = &date
	}
}

// WithEndDate sets the end date option
func WithEndDate(date time.Time) Option {
	return func(opts *HistoryOptions) {
		opts.EndDate = &date
	}
}

func (c *Client) GetMemberGymSessions(opts ...Option) (*GetMemberGymSessionsResponse, error) {
	route := "gymSessions/member"

	// Default options
	options := &HistoryOptions{}

	// Apply each option to the HistoryOptions
	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	if options.StartDate != nil {
		params.Add("fromDate", options.StartDate.Format("2006-01-02T15:04:05"))
	}
	if options.EndDate != nil {
		params.Add("toDate", options.EndDate.Format("2006-01-02T15:04:05"))
	}

	var gymSessions GetMemberGymSessionsResponse
	err := c.sendRequest("GET", route, nil, params, &gymSessions)
	if err != nil {
		return nil, err
	}

	return &gymSessions, nil
}
