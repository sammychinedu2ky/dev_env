build:
	go mod tidy

run:
	docker-compose -f "docker-compose.yml" up -d

test:
	go test

final: build run test