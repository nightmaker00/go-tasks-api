.PHONY: run build test tidy lint vet fmt clean deps\
	docker-up docker-down migrate-up migrate-down

APP_NAME := go-tasks-api
BIN_DIR := bin
CMD_DIR := ./cmd/app
MIGRATIONS_DIR := ./migrations

GO ?= go


run:
	@echo "Запуск приложения..."
	$(GO) run $(CMD_DIR)

build: $(BIN_DIR)/$(APP_NAME)

$(BIN_DIR)/$(APP_NAME):
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

fmt:
	$(GO) fmt ./...
	$(GO) vet ./...

lint:
	$(GO) vet ./...
	gofmt -l .

vet:
	$(GO) vet ./...

clean:
	rm -rf $(BIN_DIR)

deps:
	@echo "Установка зависимостей..."
	$(GO) mod download
	$(GO) mod tidy

docker-up:
	docker compose -f deployments/docker-compose.yml up -d

docker-down:
	docker compose -f deployments/docker-compose.yml down

migrate-up:
	psql "$$DATABASE_URL" -f $(MIGRATIONS_DIR)/000001_task.up.sql

migrate-down:
	psql "$$DATABASE_URL" -f $(MIGRATIONS_DIR)/000001_task.down.sql
