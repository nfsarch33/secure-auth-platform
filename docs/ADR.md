# Architectural Decision Records (ADR)

## ADR 001: Project Structure & Tech Stack
**Date**: 2025-11-23
**Status**: Accepted

### Context
We need to build a scalable, secure authentication application with a tight deadline (4 days).

### Decision
- **Backend**: Go + Gin (performance, simplicity, strong typing).
- **Frontend**: React + Vite + TypeScript (modern standard, type safety).
- **Database**: PostgreSQL (ACID compliance, robust).
- **Auth**: JWT (stateless, scalable).
- **Password Hashing**: Argon2id (OWASP recommended, memory-hard).
- **Architecture**: Clean Architecture (separation of concerns, testability).
- **API Contract**: OpenAPI 3.0 (single source of truth, code generation).

### Consequences
- **Positive**: High maintainability, clear boundaries, type safety across stack.
- **Negative**: Slight boilerplate overhead for Clean Architecture (acceptable for long-term quality).

