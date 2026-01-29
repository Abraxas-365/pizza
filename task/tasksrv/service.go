package tasksrv

import (
	"context"
	"time"

	"github.com/Abraxas-365/pizza/task"
)

type TaskService struct {
	taskRepo task.Repository
}

func NewTaskService(taskRepo task.Repository) *TaskService {
	return &TaskService{
		taskRepo,
	}
}

func (s TaskService) CreateTaskType(
	ctx context.Context,
	req task.CreateTaskTypeRequest,
) (*task.TaskType, error) {
	t := task.TaskType{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		CreatedAt: time.Now(),
	}

	return s.taskRepo.Save(ctx, t)
}
