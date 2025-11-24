.PHONY: dev test lint build gen-api install clean test-backend lint-backend test-frontend lint-frontend test-e2e test-e2e-docker security test-all run swagger

install:
	cd backend && go mod download && go install github.com/onsi/ginkgo/v2/ginkgo@v2.13.0 && go install go.uber.org/mock/mockgen@latest && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cd frontend && npm install

dev:
	docker-compose up --build

run: dev

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
	# Ensure stack is up with RECAPTCHA disabled for tests
	RECAPTCHA_DISABLED=true docker-compose up -d
	# Wait for backend to be ready
	@echo "Waiting for backend to be ready..."
	@sleep 10
	# Run Playwright
	docker run --rm --init --network secure-auth-platform_auth-net \
		-v "$(PWD)/frontend:/app" \
		-w /app \
		-e BASE_URL=http://secure-auth-platform-frontend-1 \
		-e BACKEND_URL=http://secure-auth-platform-backend-1:8080 \
		-e CI=true \
		mcr.microsoft.com/playwright:v1.56.1-jammy \
		/bin/bash -c "npm ci && npx playwright test"

lint: lint-backend lint-frontend

lint-backend:
	@echo "Running Backend Lint..."
	docker run --rm -v "$(PWD)/backend:/app" -v golangci-lint-cache:/root/.cache -v go-mod-cache:/go/pkg/mod -w /app golangci/golangci-lint:v1.61.0 golangci-lint run -v

lint-frontend:
	cd frontend && npm run lint

build:
	docker build -t auth-backend:latest ./backend
	docker build -t auth-frontend:latest ./frontend

gen-api:
	cd backend && go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -package api -generate types api/openapi.yaml > internal/api/types.go
	# Post-processing to suppress gosec G101 for BearerAuthScopes
	@if [ "$$(uname)" = "Darwin" ]; then \
		sed -i '' 's/BearerAuthScopes = "BearerAuth.Scopes"/BearerAuthScopes = "BearerAuth.Scopes" \/\/ #nosec G101/g' backend/internal/api/types.go; \
	else \
		sed -i 's/BearerAuthScopes = "BearerAuth.Scopes"/BearerAuthScopes = "BearerAuth.Scopes" \/\/ #nosec G101/g' backend/internal/api/types.go; \
	fi
	# Using docker to avoid local shell issues for frontend generation if needed, but keeping npm run for now as it works in CI/Docker
	cd frontend && npm run generate-api

swagger:
	@echo "Generating Swagger docs..."
	docker run --rm -v "$(PWD)/backend:/app" -w /app golang:1.23-alpine sh -c "apk add --no-cache git && go install github.com/swaggo/swag/cmd/swag@latest && swag init -g cmd/server/main.go -o docs"

clean:
	docker-compose down -v

security:
	@echo "Running Gosec..."
	docker run --rm -v "$(PWD)/backend:/app" -w /app securego/gosec:latest ./...

test-all: test-backend test-frontend test-e2e-docker security
