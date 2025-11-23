# Changelog

All notable changes to this project will be documented in this file.

## [1.0.0] - 2025-11-23

### Added
- **Backend**:
  - Implemented Clean Architecture structure (handlers, service, repository, models).
  - Added `PostgresUserRepository` with `pgxmock` tests.
  - Added `AuthService` with Argon2id hashing and JWT generation.
  - Added `AuthHandler` for Sign Up and Sign In endpoints.
  - configured `golangci-lint` for strict linting.
  - Dockerized backend with multi-stage build (Alpine based).
- **Frontend**:
  - Scaffolded React + TypeScript + Vite project.
  - Integrated `openapi-typescript-codegen` for type-safe API client.
  - Implemented `SignUpForm` and `SignInForm` with React Hook Form and Zod validation.
  - Added Unit tests with Vitest and React Testing Library (JSDOM environment).
  - Added E2E tests with Cypress.
  - Dockerized frontend with Nginx.
- **Infrastructure**:
  - Created `docker-compose.yml` for full stack local development.
  - Added Kubernetes manifests (`k8s/`) for Deployment, Service, and Ingress.
  - Added `Makefile` for common automation tasks.
- **Documentation**:
  - Created comprehensive `README.md`.
  - Created technical `REPORT.md`.
  - Created `docs/progress.md` and `docs/ADR.md`.
- **CI/CD**:
  - Configured `husky` and `commitlint` for Conventional Commits.
  - Added pre-push hooks for linting and testing.

### Fixed
- Resolved Go 1.23 compatibility issues with `go.uber.org/mock`.
- Fixed `document is not defined` error in frontend tests by using `jsdom`.
- Fixed Docker build issues for frontend (API client generation, missing CSS).
- Fixed backend test flakiness (time matching in SQL mocks).

