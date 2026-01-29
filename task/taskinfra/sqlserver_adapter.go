package taskinfra

import (
	"context"
	"time"

	"github.com/Abraxas-365/pizza/task"
	"github.com/jmoiron/sqlx"
)

type SQLServerTaskRepository struct {
	db *sqlx.DB
}

func NewSQLServerTaskRepository(db *sqlx.DB) task.Repository {
	return &SQLServerTaskRepository{db: db}
}

func (r *SQLServerTaskRepository) Save(ctx context.Context, t task.TaskType) (*task.TaskType, error) {
	// Set created_at if not already set
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}

	// SQL Server query with OUTPUT clause to get the inserted ID
	query := `
		INSERT INTO task_type (name, created_at, created_by)
		OUTPUT INSERTED.id
		VALUES (@p1, @p2, @p3)
	`

	var insertedID int
	err := r.db.QueryRowContext(ctx, query, t.Name, t.CreatedAt, t.CreatedBy).Scan(&insertedID)
	if err != nil {
		return nil, err
	}

	t.ID = insertedID
	return &t, nil
}
