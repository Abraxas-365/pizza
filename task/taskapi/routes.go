package taskapi

import "github.com/gofiber/fiber/v2"

// RegisterRoutes registra las rutas de autenticación en Fiber
func (th *TaskHandler) RegisterRoutes(app *fiber.App) {
	auth := app.Group("/task")

	auth.Post("/", th.CreateTaskType)

}
