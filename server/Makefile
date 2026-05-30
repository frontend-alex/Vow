APP_NAME := vow-api
DOCKER_COMPOSE := docker compose -f deployments/docker/docker-compose.yml

.PHONY: help dev up down restart logs ps build docker-build test vet tidy run clean

help:
	@echo "Available commands:"
	@echo "  make dev          Start API + Postgres with Air hot reload"
	@echo "  make up           Start Docker services in the background"
	@echo "  make down         Stop Docker services"
	@echo "  make restart      Restart Docker services"
	@echo "  make logs         Follow Docker service logs"
	@echo "  make ps           Show Docker service status"
	@echo "  make build        Build local Go binary"
	@echo "  make docker-build Build Docker dev image"
	@echo "  make test         Run Go tests"
	@echo "  make vet          Run go vet"
	@echo "  make tidy         Run go mod tidy"
	@echo "  make run          Run API locally"
	@echo "  make clean        Remove local build artifacts"

dev:
	$(DOCKER_COMPOSE) up --build

up:
	$(DOCKER_COMPOSE) up --build -d

down:
	$(DOCKER_COMPOSE) down

restart: down up

logs:
	$(DOCKER_COMPOSE) logs -f

ps:
	$(DOCKER_COMPOSE) ps

build:
	go build -o tmp/$(APP_NAME) ./cmd/api

docker-build:
	$(DOCKER_COMPOSE) build

test:
	go test ./...

vet:
	go vet ./...

tidy:
	go mod tidy

run:
	go run ./cmd/api

clean:
	go clean
	rm -rf tmp build-errors.log
