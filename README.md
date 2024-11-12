# PureGym API Client

This project is an unofficial Go client library for interacting with the PureGym API. It provides methods to access various endpoints of the PureGym API, such as retrieving gym sessions, member information, and membership details.

## Features

- Retrieve gym sessions for a specific gym or member
- Get detailed member information
- Access membership details
- Login and authentication handling

## Installation

To install the package, use the following command:

```sh
go get github.com/varantha/puregym-api
```

## Usage

Here is an example of how to use the client:

### New Client and Login

```go
client := puregymapi.NewClient("your-email@example.com", "your-pin")

err := client.Login()
if err != nil {
    log.Fatalf("Login failed: %v", err)
}
```

### Get Membership Details

```go
member, err := client.GetMember()
if err != nil {
    log.Fatalf("Failed to get member details: %v", err)
}
fmt.Printf("Member Details: %+v\n", member)
```

### Get Number of People in a Specific Gym

```go
gymId := "22"
sessions, err := client.GetGymSessions(gymId)
if err != nil {
    log.Fatalf("Failed to get gym sessions: %v", err)
}
fmt.Printf("Current Number of Members at Gym ID %s: %+v\n", gymId, sessions.TotalPeopleInGym)
```

### Get Your Recent Gym Sessions 

```go
sessions, err := client.GetMemberGymSessions()
if err != nil {
    log.Fatalf("Failed to get recent gym sessions: %v", err)
}
fmt.Printf("Recent Gym Sessions: %+v\n", sessions)
```

### Full Example Usage

```go
package main

import (
    "fmt"
    "log"
    "time"

    puregymapi "github.com/varantha/puregym-api"
)

func main() {
    client := puregymapi.NewClient("your-email@example.com", "your-pin")

    err := client.Login()
    if err != nil {
        log.Fatalf("Login failed: %v", err)
    }

    fromDate := time.Now().AddDate(0, -3, 0)
    toDate := time.Now()
    sessions, err := client.GetMemberGymSessionsBetween(&fromDate, &toDate)
    if err != nil {
        log.Fatalf("Failed to get gym sessions: %v", err)
    }

    fmt.Printf("Gym Sessions: %+v\n", sessions)
}
```
