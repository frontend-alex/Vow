run:
	go run ./cmd/api

test:
	go test ./...

migrate-up:
	./scripts/migrate-up.sh

migrate-down:
	./scripts/migrate-down.sh

docker-up:
	docker compose -f deployments/docker/docker-compose.yml up
