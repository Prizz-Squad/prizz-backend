package db

import (
	"context"
	"log"

	"github.com/EraldCaka/prizz-backend/internal/types"
)

func (pg *Postgres) CreateFile(ctx context.Context, f *types.FileCreateRequest) (string, error) {
	query := "INSERT INTO public.files (file, task_id, file_name) VALUES ($1, $2, $3) RETURNING id"

	var fileID string
	err := pg.db.QueryRow(ctx, query, f.File, f.TaskID, f.FileName).Scan(&fileID)
	if err != nil {
		log.Printf("Unable to insert file: %v\n", err)
		return "", err
	}
	return fileID, nil
}

func (pg *Postgres) GetFiles(ctx context.Context) ([]*types.File, error) {
	var files []*types.File

	query := "SELECT * FROM public.files"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying files: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var file types.File
		err := rows.Scan(&file.ID, &file.File, &file.TaskID, &file.FileName)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		files = append(files, &file)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over file rows: %v\n", err)
		return nil, err
	}
	return files, nil
}
