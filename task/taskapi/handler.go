package taskapi

import (
	"github.com/Abraxas-365/pizza/task"
	"github.com/Abraxas-365/pizza/task/tasksrv"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskSrv tasksrv.TaskService
}

func NewTaskHandler(taskSrv tasksrv.TaskService) *TaskHandler {
	return &TaskHandler{taskSrv}
}

func (th *TaskHandler) CreateTaskType(c *fiber.Ctx) error {

	var req task.CreateTaskTypeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	t, err := th.taskSrv.CreateTaskType(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create task type",
		})
	}

	return c.JSON(t)
}
