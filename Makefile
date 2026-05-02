.PHONY: install-migrate migrate-up migrate-down migrate-create migrate-version migrate-goto migrate-force run build clean test fmt lint

DB_URL := postgres://postgres:123456@localhost:5432/marketplace_db?sslmode=disable
MIGRATIONS_DIR := migrations

install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-create:
	@echo "Creating new migration..."
	@migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-up:
	@echo "Running up migrations..."
	@migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	@echo "Running down one migration..."
	@migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migrate-goto:
	@echo "Migrating to version $(version)..."
	@migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" goto $(version)

migrate-force:
	@echo "Forcing migration version to $(version)..."
	@migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(version)

migrate-version:
	@echo "Current migration version:"
	@migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

run:
	go run main.go

build:
	go build -o bin/marketplace-api main.go

clean:
	rm -rf bin/

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run