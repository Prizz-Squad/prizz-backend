package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"log"
)

//create, getAll, getAllByUserId, getAllByStart&EndDateByUserId

// Create
func (pg *Postgres) Create(ctx context.Context, tr *types.TaskRequest) (string, error) {
	query := "INSERT INTO public.task_history (user_id, task_id, start_date, end_date) VALUES ($1, $2, $3, $4) RETURNING id"
	var taskHistoryId string
	err := pg.db.QueryRow(ctx, query, tr.UserId, tr.TaskId, tr.StartDate, tr.EndDate).Scan(&taskHistoryId)
	if err != nil {
		log.Printf("Unable to insert task: %v\n", err)
		return "", err
	}
	return taskHistoryId, nil
}

// GetAll
func (pg *Postgres) GetAll(ctx context.Context) ([]*types.TaskHistory, error) {
	var tasksHistory []*types.TaskHistory
	query := "SELECT * FROM public.task_history"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying messages: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var taskHistory types.TaskHistory
		err := rows.Scan(&taskHistory.ID, &taskHistory.UserId, &taskHistory.TaskId, &taskHistory.StartDate, &taskHistory.EndDate, &taskHistory.Hours)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		tasksHistory = append(tasksHistory, &taskHistory)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over tasks history rows: %v\n", err)
		return nil, err
	}
	return tasksHistory, nil
}

// filter by startDate & endDate &userId
func (pg *Postgres) GetTaskHistory(ctx context.Context, startTime, endTime string, userID string) ([]*types.TaskHistory, error) {
	var tasksHistory []*types.TaskHistory
	query := `SELECT id, user_id, task_id, start_date, end_date, hours 
              FROM task_history 
              WHERE start_date >= $1 AND end_date <= $2 AND user_id = $3`
	rows, err := pg.db.Query(ctx, query, startTime, endTime, userID)
	if err != nil {
		log.Printf("Error querying task history: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var taskHistory types.TaskHistory
		err := rows.Scan(&taskHistory.ID, &taskHistory.UserId, &taskHistory.TaskId, &taskHistory.StartDate, &taskHistory.EndDate, &taskHistory.Hours)
		if err != nil {
			log.Printf("Error scanning task history row: %v\n", err)
			continue
		}
		tasksHistory = append(tasksHistory, &taskHistory)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over task history rows: %v\n", err)
		return nil, err
	}
	return tasksHistory, nil
}

func (pg *Postgres) GetAllHoursByUserId(ctx context.Context, userId string) (*types.TotalCountResponse, error) {
	var totalCountResponse types.TotalCountResponse
	fmt.Println(userId)
	query := `
    SELECT user_id, SUM(hours) AS total_hours,
    COUNT(task_id) AS count_task
    FROM public.task_history
    WHERE user_id = $1
    GROUP BY user_id;
`

	row := pg.db.QueryRow(ctx, query, userId)
	err := row.Scan(&totalCountResponse.UserId, &totalCountResponse.TotalHours, &totalCountResponse.CountTask)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("report with ID %s not found", userId)
		}
		log.Printf("Error scanning report data: %v\n", err)
		return nil, err
	}
	return &totalCountResponse, nil
}
