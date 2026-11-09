# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Changed
- **Validation**: Zod 3 -> Zod 4 in the frontend. The sign-in and sign-up schemas
  replace `z.string().email()` with `z.email()`; the schema tests are unchanged and
  still assert the same messages.
- **Dependencies**: the Zod 4 breaking changes that affect this project:
  - `ZodError.errors` was removed; the issue list is `ZodError.issues`.
  - String format failures are reported with the code `invalid_format` instead of
    `invalid_string`.
  - `@hookform/resolvers` 3.x reads `ZodError.errors` and therefore cannot
    validate a Zod 4 schema; it moves to 7.x, which supports Zod 3 and Zod 4.
- **Dependencies**: ESLint 10 requires `eslint-plugin-react-hooks` 7.x, because
  5.x declares a peer range that stops at ESLint 9 and `npm ci` failed with
  ERESOLVE. `@eslint/js` moves to 10.x so the recommended rule set matches the
  ESLint major in use.
- **Dependencies**: routine updates - `react-router-dom` 6.30.6, `axios` 1.21.0,
  `vite` 8.3.6, `vitest` and `@vitest/coverage-v8` 4.1.12, `react-hook-form`
  7.66.1, `@humanfs/node` 0.16.8, `js-yaml` 4.3.2, `handlebars` 4.7.10,
  `source-map-js` 1.2.2, `follow-redirects` 1.16.3.
- **Dependencies**: backend `kin-openapi` 0.149.0 and `jsonschema/v6` 6.0.3.

## [1.1.0] - 2025-11-23

### Added
- **E2E Testing**: Added Playwright tests for authentication flow and security regression.
- **Security**: Implemented Rate Limiting middleware (60 req/min).
- **Security**: Implemented Secure Headers middleware (HSTS, CSP, X-Frame-Options, etc.).
- **CI/CD**: Added GitHub Actions workflows for CI (test/lint) and Release (build/push/tag).

### Fixed
- **Testing**: Fixed frontend unit tests hanging in CI environment by adding `CI=true`.
- **DevOps**: Resolved issue where backend Docker container was running stale code by forcing rebuild.

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
