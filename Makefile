PROJECT_NAME = music-player-api
PROJECT_PATH = cmd/$(PROJECT_NAME).go

.PHONY:run
run:
	go run $(PROJECT_PATH)

.PHONY:build-вум
build:
	docker-compose --env-file .env.dev up -d

.PHONY:test
test:
	go test ./...

.PHONY:lint
lint:
	golangci-lint run

