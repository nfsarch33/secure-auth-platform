# Rakuten Symphony Full Stack Engineer - Technical Report

## Candidate Information
- **Name**: Jason Lian
- **Email**: jaslian@gmail.com
- **GitHub**: https://github.com/nfsarch33/secure-auth-platform
- **Date**: November 23, 2025

## 1. Setup and Architectural Choices

**Architecture Pattern**: Clean Architecture with layered separation

**Backend Architecture**:
- **Presentation Layer** (`api/handlers`): HTTP handlers, request/response DTOs
- **Service Layer** (`service`): Business logic, JWT generation, password hashing
- **Repository Layer** (`repository`): Database operations, data persistence
- **Domain Layer** (`models`): Core entities (User model)

**Why This Approach**:
- **Separation of Concerns**: Each layer has single responsibility (KISS principle)
- **Testability**: Easy to mock dependencies at each layer
- **Maintainability**: Changes in DB don't affect business logic
- **Scalability**: Can swap PostgreSQL for another DB without changing business logic

**Frontend Architecture**:
- **Component-Based**: Reusable, composable UI components
- **Context API**: Centralized auth state management (user, token, isAuthenticated)
- **Protected Routes**: HOC pattern for route guards
- **API Client**: Centralized Axios instance with interceptors

**Database Schema**:
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_users_email ON users(email);
```

**API Design**:
- RESTful conventions (POST for mutations, GET for reads)
- JSON request/response bodies
- HTTP status codes (201 Created, 200 OK, 401 Unauthorized, 400 Bad Request)
- JWT in Authorization header: `Bearer <token>`

## 2. Potential Weaknesses & Production Improvements

**Current Weaknesses & Mitigations**:

1. **Rate Limiting**: Implemented basic in-memory rate limiting (per IP) using `golang.org/x/time/rate`.
   - **Production Fix**: Move to Redis-based rate limiting for distributed scaling.
2. **No Input Sanitization**: XSS vulnerabilities potential.
   - **Fix**: Use `bluemonday` for HTML sanitization, strict validation.
3. **JWT Token Storage**: Stored in localStorage (XSS risk).
   - **Fix**: Use httpOnly cookies with SameSite=Strict.
4. **No Refresh Tokens**: Long-lived access tokens.
   - **Fix**: Implement refresh token rotation pattern.
5. **Database Connection Pooling**: Implemented using `pgxpool`.
   - **Production Fix**: Tune max connections based on load testing.
6. **Secure Headers**: Implemented basic security headers (HSTS, CSP, etc.) via middleware.
   - **Fix**: Fine-tune CSP policy for production assets.
7. **No HTTPS in Production**: Traffic not encrypted locally.
   - **Fix**: Deploy with TLS certificates (Let's Encrypt), force HTTPS redirect via Ingress.
8. **No Observability**: Hard to debug production issues.
   - **Fix**: Add structured logging (zap/zerolog), metrics (Prometheus), tracing (Jaeger).
9. **Single Database**: No high availability.
   - **Fix**: PostgreSQL replication (primary-replica), read-write splitting.
10. **Password Reset Not Implemented**: Users locked out if forgotten.
    - **Fix**: Add email-based password reset flow.

## 3. Future Improvements (More Time)

**Priority 1 - Security Enhancements**:
1. **OAuth2/OIDC Integration**: Let users sign in with Google/GitHub
2. **Multi-Factor Authentication (MFA)**: TOTP-based (e.g., Google Authenticator)
3. **Audit Logging**: Track all auth events (login, failed attempts, password changes)
4. **Account Lockout**: After N failed login attempts

**Priority 2 - Testing**:
1. **Unit Tests**: 80%+ coverage for service/repository layers
2. **Integration Tests**: Test API endpoints with test database
3. **E2E Tests**: Cypress/Playwright for frontend flows
4. **Load Testing**: k6 scripts to test under 1000 req/sec

**Priority 3 - DevOps**:
1. **CI/CD Pipeline**: GitHub Actions for test → build → deploy
2. **Blue-Green Deployment**: Zero-downtime releases
3. **Auto-scaling**: HPA in Kubernetes based on CPU/memory
4. **Monitoring Dashboard**: Grafana + Prometheus for real-time metrics

## 4. Frontend State Management & Validation

**Why React Context API**:
- **Simple Requirements**: Only auth state needed (user, token, isAuthenticated)
- **Avoids Over-Engineering**: Redux would be overkill for 3 routes
- **Performance**: Context updates don't trigger unnecessary re-renders with proper memoization
- **Co-location**: Auth logic lives with auth context

**Why React Hook Form + Zod**:
- **React Hook Form**: Performance (uncontrolled inputs), minimal re-renders
- **Zod**: Type-safe schema validation, works seamlessly with RHF
- **DRY**: Define validation schema once, reuse for frontend + backend

## 5. Types & Contracts: Frontend ↔ Backend Sync

**Approach**: OpenAPI-First Design with Code Generation
- **Single Source of Truth**: OpenAPI spec defines contract
- **Type Safety**: Both sides have type-checked interfaces
- **Automatic Updates**: Regenerate when spec changes
- **Documentation**: OpenAPI spec serves as API docs

## 6. Scenario 1: Brute-Force Attack Mitigation

**Strategy**:
1. **Rate Limiting**: Per-IP limiting using Redis/memory store.
2. **Account Lockout**: Lock account for 15 mins after 5 failed attempts.
3. **Slow Hashing**: Use Argon2id which is memory-hard, making brute-force computationally expensive.
4. **Monitoring**: Alert on high failure rates.

## 7. Scenario 2: Scale to Millions of Requests/Sec

**Architecture for Scale**:
1. **Horizontal Scaling**: Kubernetes Deployment with HPA (Horizontal Pod Autoscaler) to scale pods based on CPU/Memory.
2. **Database Scaling**: Read replicas for GET requests, connection pooling (pgxpool), and potentially sharding for write-heavy loads.
3. **Caching**: Redis cluster for session/profile caching to reduce DB hits.
4. **Load Balancing**: Cloud load balancer (AWS ALB / K8s Ingress) to distribute traffic.
5. **Asynchronous Processing**: Use message queues (RabbitMQ/Kafka) for non-critical tasks (emails, analytics).

## Appendix
- API Documentation (OpenAPI spec)
- Database Schema
- Deployment Instructions

