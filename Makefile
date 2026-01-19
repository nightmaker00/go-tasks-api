.PHONY: run build test tidy lint lint-install vet fmt clean deps\
	docker-up docker-down migrate-up migrate-down

APP_NAME := go-tasks-api
BIN_DIR := bin
CMD_DIR := ./cmd/app
MIGRATIONS_DIR := ./migrations

GO ?= go
GOLANGCI_LINT ?= golangci-lint


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
	$(GOLANGCI_LINT) run ./...

lint-install:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

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
	@for f in $(MIGRATIONS_DIR)/*.up.sql; do \
		echo "apply $$f"; \
		psql "$$DATABASE_URL" -f $$f; \
	done

migrate-down:
	@for f in $$(ls -r $(MIGRATIONS_DIR)/*.down.sql); do \
		echo "rollback $$f"; \
		psql "$$DATABASE_URL" -f $$f; \
	done
