package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Abraxas-365/pizza/task/taskapi"
	"github.com/Abraxas-365/pizza/task/taskinfra"
	"github.com/Abraxas-365/pizza/task/tasksrv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb" // MS SQL Server driver
)

func main() {
	// Initialize database connection
	db, err := initDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("✅ Database connected successfully")

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Pizza API v1.0",
		// ErrorHandler: GlobalErrorHandler, // Uncomment when implemented
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// Initialize layers: Repository -> Service -> Handler
	taskRepo := taskinfra.NewSQLServerTaskRepository(db)
	taskService := tasksrv.NewTaskService(taskRepo)
	taskHandler := taskapi.NewTaskHandler(*taskService)

	// Register routes
	taskHandler.RegisterRoutes(app)

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"app":    "Pizza API v1.0",
		})
	})

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}

// initDB initializes the MS SQL Server database connection
func initDB() (*sqlx.DB, error) {
	host := os.Getenv("MSSQL_HOST")
	port := os.Getenv("MSSQL_PORT")
	user := os.Getenv("MSSQL_USER")
	password := os.Getenv("MSSQL_PASSWORD")
	database := os.Getenv("MSSQL_DATABASE")

	// Validate required environment variables
	if host == "" || port == "" || user == "" || password == "" || database == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}

	// Build connection string for MS SQL Server
	// Format: sqlserver://username:password@host:port?database=dbname
	connStr := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=30",
		user, password, host, port, database,
	)

	// Open database connection
	db, err := sqlx.Connect("sqlserver", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

