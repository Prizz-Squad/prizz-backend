package types

import (
	"time"
)

type Ticket struct {
	ID          string    `json:"id"`
	Priority    int       `json:"priority"`
	Status      int       `json:"status"`
	Department  int       `json:"department"`
	Images      []byte    `json:"images"`
	UserID      string    `json:"userID"`
	ProjectID   string    `json:"projectID"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
}

type CreateTicketRequest struct {
	Priority    int       `json:"priority"`
	Status      int       `json:"status"`
	Department  int       `json:"department"`
	Images      []byte    `json:"images"`
	UserID      string    `json:"userID"`
	ProjectID   string    `json:"projectID"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
}

type UpdateDepartment struct {
	Department int `json:"department"`
}

type UpdateStatus struct {
	Status int `json:"status"`
}

type AddNewImage struct {
	Images []byte `json:"images"`
}

type UpdateDescription struct {
	Description string `json:"description"`
}

type UpdatePostDate struct {
}
