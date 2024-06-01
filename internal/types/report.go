package types

import "time"

type Report struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Hours     int       `json:"hours"`
}

type ReportRequest struct {
	UserID    string `json:"userId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
type ReportResponse struct {
	UserID    string `json:"userId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Hours     int    `json:"hours"`
}
