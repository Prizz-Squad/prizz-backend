package types

import (
	"time"
)

var (
	LowPriority  = 0
	MidPriority  = 1
	HighPriority = 2
)

var (
	Todo       = 0
	InProgress = 1
	Done       = 2
)

var (
	Design     = 0
	Caption    = 1
	Scheduling = 2
)

type Ticket struct {
	ID          string    `json:"id"`
	Priority    int       `json:"priority"`
	Status      int       `json:"status"`
	Department  int       `json:"department"`
	UserID      string    `json:"userID"`
	ProjectID   string    `json:"projectID"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	PostDate    time.Time `json:"postDate"`
}

type CreateTicketRequest struct {
	Priority    int       `json:"priority"`
	Status      int       `json:"status"`
	Department  int       `json:"department"`
	UserID      string    `json:"userID"`
	ProjectID   string    `json:"projectID"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	PostDate    string    `json:"postDate"`
}

type UpdateDepartment struct {
	Department int `json:"department"`
}

type UpdateStatus struct {
	Status int `json:"status"`
}

type UpdateDescription struct {
	Description string `json:"description"`
}

type UpdatePostDate struct {
	PostDate string `json:"postDate"`
}
