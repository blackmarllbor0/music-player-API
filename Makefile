PROJECT_NAME = music-player-api
PROJECT_PATH = cmd/$(PROJECT_NAME).go
ENV_MODE ?= dev

.PHONY:run
run:
	go run $(PROJECT_PATH)

.PHONY:build
build:
	docker-compose --env-file .env.${ENV_MODE} \
		--profile tools run --rm postgres-migrate up

.PHONY:down
down:
	docker-compose down

.PHONY:test
test:
	go test ./...

.PHONY:lint
lint:
	golangci-lint run