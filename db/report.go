package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"log"
)

// Create
func (pg *Postgres) CreateReport(ctx context.Context, tr *types.ReportRequest) (string, error) {
	query := "INSERT INTO public.report (user_id, start_date, end_date) VALUES ($1, $2, $3) RETURNING id"
	var reportId string
	err := pg.db.QueryRow(ctx, query, tr.UserID, tr.StartDate, tr.EndDate).Scan(&reportId)
	if err != nil {
		log.Printf("Unable to insert report: %v\n", err)
		return "", err
	}
	return reportId, nil
}

// GetAll
func (pg *Postgres) GetAllReports(ctx context.Context) ([]*types.Report, error) {
	var reports []*types.Report
	query := "SELECT * FROM public.report"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying reports: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var report types.Report
		err := rows.Scan(&report.ID, &report.UserID, &report.StartDate, &report.EndDate, &report.Hours)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		reports = append(reports, &report)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over tasks reports rows: %v\n", err)
		return nil, err
	}
	return reports, nil
}

// filter by startDate & endDate &userId
func (pg *Postgres) GetReportHistory(ctx context.Context, startTime, endTime string, userID string) ([]*types.Report, error) {
	var reportHistory []*types.Report
	query := `SELECT id, user_id, start_date, end_date, hours 
              FROM report 
              WHERE start_date >= $1 AND end_date <= $2 AND user_id = $3`
	rows, err := pg.db.Query(ctx, query, startTime, endTime, userID)
	if err != nil {
		log.Printf("Error querying report history: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var report types.Report
		err := rows.Scan(&report.ID, &report.UserID, &report.StartDate, &report.EndDate, &report.Hours)
		if err != nil {
			log.Printf("Error scanning report history row: %v\n", err)
			continue
		}
		reportHistory = append(reportHistory, &report)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over report history rows: %v\n", err)
		return nil, err
	}
	return reportHistory, nil
}

func (pg *Postgres) GetAllHoursReportsByUserId(ctx context.Context, userId string) (*types.TotalCountResponse, error) {
	var totalCountResponse types.TotalCountResponse
	query := `
    SELECT user_id, SUM(hours) AS total_hours,
    COUNT(user_id) as count_task
    FROM public.report
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
