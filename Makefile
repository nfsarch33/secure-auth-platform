.PHONY: dev test lint build gen-api install clean test-backend lint-backend test-frontend lint-frontend

install:
	cd backend && go mod download && go install github.com/onsi/ginkgo/v2/ginkgo@v2.13.0 && go install go.uber.org/mock/mockgen@latest && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd frontend && npm install

dev:
	docker-compose up --build

test: test-backend test-frontend

test-backend:
	cd backend && ginkgo -r -v --cover ./...

test-frontend:
	cd frontend && npm test

lint: lint-backend lint-frontend

lint-backend:
	cd backend && golangci-lint run

lint-frontend:
	cd frontend && npm run lint

build:
	docker build -t auth-backend:latest ./backend
	docker build -t auth-frontend:latest ./frontend

gen-api:
	cd backend && go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -package api -generate types,server api/openapi.yaml > internal/api/types.go
	cd frontend && npm run generate-api

clean:
	docker-compose down -v
