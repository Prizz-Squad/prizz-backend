package db

import (
	"context"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"log"
)

func (pg *Postgres) CreateProject(ctx context.Context, p *types.ProjectCreateRequest) (string, error) {
	query := "INSERT INTO public.projects (name, description, creator_id) VALUES ($1, $2, $3) RETURNING id"

	var projectID string
	err := pg.db.QueryRow(ctx, query, p.Name, p.Description, p.CreatorID).Scan(&projectID)
	if err != nil {
		log.Printf("Unable to insert project: %v\n", err)
		return "", err
	}
	return projectID, nil
}

func (pg *Postgres) GetProjects(ctx context.Context) ([]*types.Project, error) {
	var projects []*types.Project

	query := "SELECT * FROM public.projects"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying projects: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project types.Project
		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.CreatorID)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		projects = append(projects, &project)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v\n", err)
		return nil, err
	}
	return projects, nil
}
