# Secure Auth Platform

A production-ready, full-stack authentication system built with **Go (Golang)**, **React (TypeScript)**, **PostgreSQL**, and **Docker/Kubernetes**.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.23-blue.svg)
![React Version](https://img.shields.io/badge/react-18-blue.svg)

## 🚀 Features

*   **Secure Authentication**: JWT-based stateless authentication with Argon2id password hashing.
*   **Protected Routes**: Frontend route guards ensure only authenticated users access private pages.
*   **Modern UI/UX**: Responsive card-based design with accessibility compliance (WCAG 2.1 AA).
*   **Security First**:
    *   Rate Limiting (Token Bucket algorithm)
    *   Secure Headers (HSTS, CSP, X-Frame-Options)
    *   reCAPTCHA v3 Integration
    *   Input Validation (Zod + Backend Validation)
*   **Observability**: Structured logging (`log/slog`) and health checks.
*   **DevOps Ready**: Dockerized stack, Kubernetes manifests, and a comprehensive `Makefile`.

## 🛠️ Tech Stack

*   **Backend**: Go (Gin Framework), pgx (PostgreSQL Driver), Testify/Ginkgo/Gomega (Testing)
*   **Frontend**: React, TypeScript, Vite, React Router, React Hook Form, Zod, Tailwind/CSS Modules
*   **Database**: PostgreSQL 15
*   **Infrastructure**: Docker, Docker Compose, Kubernetes (K8s)
*   **Tooling**: OpenAPI Generator, Makefile, golangci-lint, ESLint, Prettier

## 🏁 Getting Started

### Prerequisites

*   Docker & Docker Compose
*   Go 1.21+ (for local dev)
*   Node.js 18+ (for local dev)
*   Make

### ⚡️ Quick Start (Docker)

Run the entire stack (Backend, Frontend, Database) with one command:

```bash
# Create .env from example
cp .env.example .env

# Start the stack
make run
```

Access the application:
*   **Frontend**: `http://localhost` (Recommended, running via Docker)
*   **Backend API**: `http://localhost:8080`
*   **Swagger Docs**: `http://localhost:8080/swagger/index.html`

*Note: If you are running the frontend manually (`npm run dev`) outside of Docker, access it at `http://localhost:3000`.*

### 🧪 Testing

The project includes a comprehensive test suite covering Unit, Integration, and End-to-End (E2E) tests.

#### Run All Tests
```bash
make test-all
```

#### Run Specific Suites
*   **Backend Unit & Integration**: `make test-backend`
*   **Frontend Unit**: `make test-frontend`
*   **E2E (Playwright in Docker)**: `make test-e2e-docker`

### 🔧 Development Workflow

1.  **Backend Dev**:
    ```bash
    make dev-backend
    ```
2.  **Frontend Dev**:
    ```bash
    make dev-frontend
    ```
3.  **Linting**:
    ```bash
    make lint-backend
    make lint-frontend
    ```

### ☸️ Kubernetes Deployment

Deploy the stack to a local Kubernetes cluster (e.g., Minikube, Docker Desktop K8s):

```bash
# Apply K8s manifests
kubectl apply -f k8s/

# Verify pods are running
kubectl get pods
```

## 📂 Project Structure

```
.
├── backend/                # Go Backend Service
│   ├── cmd/server/         # Entry point
│   ├── internal/           # Private application code
│   │   ├── api/            # Handlers & Middleware
│   │   ├── models/         # Domain models
│   │   ├── repository/     # Database access
│   │   └── service/        # Business logic
│   └── pkg/                # Public libraries (JWT, Hash, etc.)
├── frontend/               # React Frontend Application
│   ├── src/
│   │   ├── api/            # Generated OpenAPI client
│   │   ├── components/     # UI Components
│   │   ├── contexts/       # React Contexts (Auth)
│   │   └── e2e/            # Playwright E2E tests
├── k8s/                    # Kubernetes Manifests
└── Makefile                # Automation scripts
```

## 🔒 Security Implementation

*   **Password Hashing**: Argon2id is used for password hashing, providing resistance against GPU-based brute-force attacks.
*   **JWT**: Short-lived access tokens (configurable duration) signed with HMAC-SHA256.
*   **CSP**: Strict Content Security Policy to prevent XSS.
*   **Rate Limiting**: In-memory rate limiter to mitigate DDoS and brute-force login attempts.

## 📄 License

This project is licensed under the MIT License.
