cat > Makefile << 'EOL'
.PHONY: build run test clean install docker-build docker-run

build:
	go build -o bin/feast-friends-api cmd/server/main.go

run:
	go run cmd/server/main.go

install:
	go mod tidy
	go mod download

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	rm -rf bin/
	rm -rf dist/
	go clean

fmt:
	go fmt ./...

docker-build:
	docker build -t feast-friends-api -f docker/Dockerfile .

docker-up:
	docker-compose -f docker/docker-compose.yml up --build

docker-down:
	docker-compose -f docker/docker-compose.yml down

migrate-up:
	@echo "Run database migrations here"

migrate-down:
	@echo "Rollback database migrations here"
EOL