ifneq ("$(wildcard .env)", "")
include .env
else ifneq "$(MAKECMDGOALS)" "setup-env"
$(error No .env file found. Please run setup-env first)
endif

.PHONY: service-up service-down run-boardgame migration-up migration-down setup-env

service-up:
	docker compose -f docker-compose.yml up -d

service-down:
	docker compose -f docker-compose.yml down

run-boardgame:
	go run main.go

migration-up:
	migrate -path ./migrations -database ${DB}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable up

migration-down:
	migrate -path ./migrations -database ${DB}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable down

setup-env:
	cp env/$(ENV).env .env