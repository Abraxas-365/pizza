package task

import "context"

type Repository interface {
	Save(ctx context.Context, t TaskType) (*TaskType, error)
}
