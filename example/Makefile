export
POSTGRES_DB := postgres
POSTGRES_USER := postgres
POSTGRES_PASS := postgres
POSTGRES_PORT := 15432
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

.PHONY: run
run:
	@go run main.go

.PHONY: migrate
migrate:
	@go run cmd/main.go
