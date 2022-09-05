POSTGRES_VERSION = 14.5

export
POSTGRES_DB := postgres
POSTGRES_USER := postgres
POSTGRES_PASS := postgres
POSTGRES_PORT := 15432
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
CONTAINER_NAME := go-accept-interfaces-return-structs

.PHONY: db
db:
	@docker run --rm -d \
		-p $(POSTGRES_PORT):5432 \
		-e TZ=UTC \
		-e LANG=ja_JP.UTF-8 \
		-e POSTGRES_HOST_AUTH_METHOD=trust \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASS) \
		-e POSTGRES_INITDB_ARGS=--encoding=UTF-8 \
		--name $(CONTAINER_NAME) \
		postgres:$(POSTGRES_VERSION)-alpine

.PHONY: psql
psql:
	@docker exec -it $(CONTAINER_NAME) psql -U postgres

.PHONY: stop
stop:
	@docker stop $(CONTAINER_NAME)

