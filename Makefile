export MSSQL_HOST=localhost
export MSSQL_PORT=1433
export MSSQL_USER=sa
export MSSQL_PASSWORD=YourStrong!Passw0rd
export MSSQL_DATABASE=pizzadb


# PostgreSQL
export POSTGRES_DB = loopdb
export POSTGRES_USER = loop
export POSTGRES_PASSWORD = supersecret
export POSTGRES_HOST = localhost
export POSTGRES_PORT = 5432
# Server
export PORT = 8081
export ENVIRONMENT = development

.PHONY: dev
dev: ## Run the development server
	go mod tidy
	go run ./cmd/server


.PHONY: up
up: ## Start all services (postgres + redis)
	docker compose up -d --remove-orphans
	@echo "Waiting for services to be ready..."
	@sleep 3
	@make health

.PHONY: down
down: ## Stop all services
	docker compose down

