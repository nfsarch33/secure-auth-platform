.PHONY: dev test lint build gen-api install clean test-backend lint-backend test-frontend lint-frontend test-e2e test-e2e-docker

install:
	cd backend && go mod download && go install github.com/onsi/ginkgo/v2/ginkgo@v2.13.0 && go install go.uber.org/mock/mockgen@latest && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd frontend && npm install

dev:
	docker-compose up --build

test: test-backend test-frontend test-e2e-docker

test-backend:
	@echo "Running Backend Unit & DB Tests..."
	docker run --rm -v "$(PWD)/backend:/app" -w /app golang:1.23-alpine sh -c "apk add --no-cache git && go mod tidy && go install github.com/onsi/ginkgo/v2/ginkgo@v2.13.0 && ginkgo -r -v --cover ./..."

test-frontend:
	@echo "Running Frontend Unit Tests..."
	docker run --rm -v "$(PWD)/frontend:/app" -w /app node:20-alpine sh -c "npm install && CI=true npm test"

test-e2e:
	cd frontend && npm run test:e2e

test-e2e-docker:
	@echo "Running Playwright E2E Tests in Docker..."
	# Ensure stack is up
	docker-compose up -d
	# Run Playwright
	docker run --rm --init --network secure-auth-platform_auth-net \
		-v "$(PWD)/frontend:/app" \
		-w /app \
		-e BASE_URL=http://frontend \
		-e BACKEND_URL=http://backend:8080 \
		-e CI=true \
		mcr.microsoft.com/playwright:v1.56.1-jammy \
		/bin/bash -c "npm ci && npx playwright test"

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
