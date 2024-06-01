package types

import "time"

type TaskHistory struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	TaskId    string    `json:"task_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Hours     int       `json:"hours"`
}

type TotalCountResponse struct {
	UserId     string `json:"id"`
	TotalHours int    `json:"total_hours"`
	CountTask  int    `json:"count_task"`
}

type TaskRequest struct {
	UserId    string    `json:"user_id"`
	TaskId    string    `json:"task_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type TaskResponse struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	TaskId    string    `json:"task_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
