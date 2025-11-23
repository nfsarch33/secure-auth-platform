<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# ```

Wendy Low [Wendy.Low@newbridgealliance.com](mailto:Wendy.Low@newbridgealliance.com)
17:03 (50 minutes ago)
to Jason
 
Dear Jason
Hope this email finds you well.
Congratulations on being selected for the next round with Rakuten Symphony  – Full Stack Engineer in Australia.
For this stage, the client is requesting your participation in the assignment, with details as outlined below.
The deadline is Monday, 24 November, 23:59 AEDT.
Please review the information and feel free to reach out if you have any questions.
-----------------------
Please find your assignment through the provided link and submit it by the deadline below.
 
If you require an adjustment to the assignment submission deadline, please feel free to reply to this message with your reason for the adjustment and your proposed new deadline.
 
Assignment Link: [https://rak.box.com/s/1xr8grwehuu5ktjyxylq4bq9bp37etsq](https://rak.box.com/s/1xr8grwehuu5ktjyxylq4bq9bp37etsq)
Assignment Submission Deadline: November 24th (Mon) 23:59 AEDT
-----------------------

Best regards
Wendy Low
Director
Newbridge Alliance Pte Ltd
A: 10 Anson Road, international plaza, \#27-02, Singapore 079903
E : [wendy.low@newbridgealliance.com](mailto:Tripti.devleker@newbridgealliance.com)
M: +65 9820 8055
W: [www.newbridgealliance.com](http://www.newbridgealliance.com/)

EA license - 20S0283
EA license personnel: R1442533
Instagram: [https://www.instagram.com/newbridgesg/](https://www.instagram.com/newbridgesg/)

```

Review, analyse and research for the Rakuten Symphony interview process. 

Analyse and prepare for the interview code challenge, based on the up-to-date and best practice of software engineering, using Golang, React and a DB with docker and k8s using REST API (OpenAI with codegen) and a Makefile. Should follow strictly the instructions. with proper folder structure and documentations.

Based on the assignment requirements and your background in Golang, React, and DevOps, let me create a comprehensive analysis and implementation plan for the Rakuten Symphony Full Stack Engineer coding challenge.

# Rakuten Symphony Full Stack Engineer - Assignment Analysis & Implementation Plan

## Assignment Overview

You need to build a **minimal authentication application** with:
- **Backend**: REST API (signup, signin, protected profile endpoint) with database persistence
- **Frontend**: React SPA with signup, signin, protected profile page, and signout
- **Deadline**: Monday, November 24, 2025, 23:59 AEDT (4 days)
- **Deliverable**: Public GitHub repository + written report

## Tech Stack Selection

Based on your expertise and the requirements:

### Backend
- **Language**: Golang (preferred - aligns with your experience)
- **Framework**: Gin or Chi (lightweight, production-ready)
- **Database**: PostgreSQL with Docker
- **Authentication**: JWT tokens (industry standard for stateless REST APIs)
- **API Specification**: OpenAPI 3.0 with code generation

### Frontend
- **Framework**: React with TypeScript
- **State Management**: React Context + hooks (simple, no Redux needed for this scope)
- **Form Validation**: React Hook Form + Zod
- **HTTP Client**: Axios with interceptors for auth headers
- **Routing**: React Router v6

### DevOps & Infrastructure
- **Containerization**: Docker + Docker Compose
- **Orchestration**: Kubernetes manifests (basic deployment)
- **Build Tool**: Makefile for common tasks
- **CI/CD**: GitHub Actions (optional but recommended)

***

## Project Structure

```

auth-app/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── api/
│   │   │   ├── handlers/
│   │   │   │   ├── auth.go
│   │   │   │   └── user.go
│   │   │   ├── middleware/
│   │   │   │   ├── auth.go
│   │   │   │   └── cors.go
│   │   │   └── router.go
│   │   ├── models/
│   │   │   └── user.go
│   │   ├── repository/
│   │   │   ├── user_repository.go
│   │   │   └── postgres.go
│   │   ├── service/
│   │   │   ├── auth_service.go
│   │   │   └── user_service.go
│   │   └── config/
│   │       └── config.go
│   ├── pkg/
│   │   ├── jwt/
│   │   │   └── jwt.go
│   │   └── password/
│   │       └── hash.go
│   ├── api/
│   │   └── openapi.yaml
│   ├── migrations/
│   │   └── 001_create_users_table.sql
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── src/
│   │   ├── api/
│   │   │   ├── client.ts
│   │   │   └── auth.ts
│   │   ├── components/
│   │   │   ├── SignUp.tsx
│   │   │   ├── SignIn.tsx
│   │   │   ├── Profile.tsx
│   │   │   └── PrivateRoute.tsx
│   │   ├── contexts/
│   │   │   └── AuthContext.tsx
│   │   ├── types/
│   │   │   └── auth.ts
│   │   ├── App.tsx
│   │   └── main.tsx
│   ├── package.json
│   ├── tsconfig.json
│   ├── vite.config.ts
│   └── Dockerfile
├── k8s/
│   ├── backend-deployment.yaml
│   ├── frontend-deployment.yaml
│   ├── postgres-deployment.yaml
│   └── ingress.yaml
├── docker-compose.yml
├── Makefile
└── README.md

```

***

## Report Questions - Detailed Answers

### 1. **Setup and Architectural Choices**

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

***

### 2. **Potential Weaknesses \& Production Improvements**

**Current Weaknesses**:

1. **No Rate Limiting**: Vulnerable to brute-force attacks
    - **Fix**: Implement rate limiting middleware (e.g., `go-rate-limiter`, Redis-based)

```go
limiter := rate.NewLimiter(rate.Every(time.Minute), 10) // 10 req/min
```

2. **No Input Sanitization**: XSS vulnerabilities
    - **Fix**: Use `bluemonday` for HTML sanitization, strict validation
3. **JWT Token Storage**: Stored in localStorage (XSS risk)
    - **Fix**: Use httpOnly cookies with SameSite=Strict
4. **No Refresh Tokens**: Long-lived access tokens
    - **Fix**: Implement refresh token rotation pattern
5. **No Database Connection Pooling**: May exhaust connections under load
    - **Fix**: Use `pgxpool` with max connections configured
6. **No HTTPS in Production**: Traffic not encrypted
    - **Fix**: Deploy with TLS certificates (Let's Encrypt), force HTTPS redirect
7. **No Observability**: Hard to debug production issues
    - **Fix**: Add structured logging (zap/zerolog), metrics (Prometheus), tracing (Jaeger)
8. **Single Database**: No high availability
    - **Fix**: PostgreSQL replication (primary-replica), read-write splitting
9. **No CSRF Protection**: If using cookies
    - **Fix**: Implement CSRF tokens for cookie-based auth
10. **Password Reset Not Implemented**: Users locked out if forgotten
    - **Fix**: Add email-based password reset flow

**Production Checklist**:

- [ ] Rate limiting per IP/user
- [ ] HTTPS with TLS 1.3
- [ ] httpOnly, Secure, SameSite cookies
- [ ] Refresh token rotation
- [ ] Database connection pooling
- [ ] Health check endpoints (`/health`, `/ready`)
- [ ] Graceful shutdown
- [ ] Structured logging with correlation IDs
- [ ] Metrics and alerting
- [ ] CORS properly configured
- [ ] Content Security Policy headers
- [ ] Database migrations automated (goose/migrate)
- [ ] Secrets management (Vault, AWS Secrets Manager)

***

### 3. **What I'd Improve Next (More Time)**

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

**Priority 4 - Features**:

1. **Email Verification**: Send verification email on signup
2. **Password Reset**: Email-based password recovery
3. **Profile Update**: Change email/password endpoints
4. **Role-Based Access Control (RBAC)**: Admin vs. regular user permissions
5. **Session Management**: View active sessions, revoke tokens

***

### 4. **Frontend State Management \& Validation**

**Why React Context API**:

- **Simple Requirements**: Only auth state needed (user, token, isAuthenticated)
- **Avoids Over-Engineering**: Redux would be overkill for 3 routes
- **Performance**: Context updates don't trigger unnecessary re-renders with proper memoization
- **Co-location**: Auth logic lives with auth context

**State Structure**:

```typescript
interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
}

interface AuthContextValue extends AuthState {
  signIn: (email: string, password: string) => Promise<void>;
  signUp: (email: string, password: string) => Promise<void>;
  signOut: () => void;
}
```

**Why React Hook Form + Zod**:

- **React Hook Form**: Performance (uncontrolled inputs), minimal re-renders
- **Zod**: Type-safe schema validation, works seamlessly with RHF
- **DRY**: Define validation schema once, reuse for frontend + backend

**Validation Example**:

```typescript
const signUpSchema = z.object({
  email: z.string().email("Invalid email"),
  password: z.string().min(8, "Min 8 characters")
    .regex(/[A-Z]/, "Must contain uppercase")
    .regex(/[0-9]/, "Must contain number"),
});
```

**Data Flow**:

1. User submits form → React Hook Form validates
2. If valid → API call via Axios client
3. Success → Update AuthContext, store token, redirect
4. Error → Display error message (toast/alert)

***

### 5. **Types \& Contracts: Frontend ↔ Backend Sync**

**Approach**: OpenAPI-First Design with Code Generation

**Process**:

1. **Define OpenAPI Spec** (`api/openapi.yaml`):

```yaml
paths:
  /api/auth/signup:
    post:
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/SignUpRequest'
      responses:
        '201':
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/AuthResponse'

components:
  schemas:
    SignUpRequest:
      type: object
      required: [email, password]
      properties:
        email: { type: string, format: email }
        password: { type: string, minLength: 8 }
```

2. **Generate Backend Code** (Golang):

```bash
# Using oapi-codegen
oapi-codegen -package api -generate types,server api/openapi.yaml > internal/api/types.go
```

3. **Generate Frontend Client** (TypeScript):

```bash
# Using openapi-typescript-codegen
openapi-typescript-codegen --input api/openapi.yaml --output src/api/generated
```


**Benefits**:

- **Single Source of Truth**: OpenAPI spec defines contract
- **Type Safety**: Both sides have type-checked interfaces
- **Automatic Updates**: Regenerate when spec changes
- **Documentation**: OpenAPI spec serves as API docs
- **Validation**: Both sides validate against same schema

**Alternative (Manual Approach)**:

- **Shared Types Repository**: Separate npm package with TypeScript types
- **Generate from Go**: Use tools like `tygo` to convert Go structs → TypeScript
- **Runtime Validation**: Use `ajv` on frontend, `go-playground/validator` on backend

**Why OpenAPI**:

- Industry standard
- Tooling ecosystem (Swagger UI, Postman import)
- Contract testing support
- Easier onboarding for new devs

***

### 6. **Scenario 1: Brute-Force Attack on Logins**

**Attack Vector**: Attacker tries millions of password combinations on `/api/auth/signin`

**Mitigation Strategies**:

**1. Rate Limiting (Primary Defense)**:

```go
// Per-IP rate limiting
import "golang.org/x/time/rate"

var limiters = make(map[string]*rate.Limiter)

func getRateLimiter(ip string) *rate.Limiter {
    limiter, exists := limiters[ip]
    if !exists {
        limiter = rate.NewLimiter(rate.Every(time.Minute), 5) // 5 attempts/min
        limiters[ip] = limiter
    }
    return limiter
}

func RateLimitMiddleware(c *gin.Context) {
    ip := c.ClientIP()
    limiter := getRateLimiter(ip)
    
    if !limiter.Allow() {
        c.JSON(http.StatusTooManyRequests, gin.H{
            "error": "Too many requests. Try again later.",
        })
        c.Abort()
        return
    }
    c.Next()
}
```

**2. Account Lockout (Defense in Depth)**:

```go
// After 5 failed attempts, lock account for 15 minutes
type User struct {
    // ... existing fields
    FailedLoginAttempts int
    LockedUntil         *time.Time
}

func (s *AuthService) SignIn(email, password string) (*User, string, error) {
    user, err := s.repo.GetUserByEmail(email)
    if err != nil {
        return nil, "", err
    }
    
    // Check if account is locked
    if user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
        return nil, "", errors.New("account locked due to failed attempts")
    }
    
    // Verify password
    if !password.Verify(user.PasswordHash, password) {
        user.FailedLoginAttempts++
        if user.FailedLoginAttempts >= 5 {
            lockUntil := time.Now().Add(15 * time.Minute)
            user.LockedUntil = &lockUntil
        }
        s.repo.UpdateUser(user)
        return nil, "", errors.New("invalid credentials")
    }
    
    // Reset failed attempts on successful login
    user.FailedLoginAttempts = 0
    user.LockedUntil = nil
    s.repo.UpdateUser(user)
    
    token, _ := s.jwt.GenerateToken(user.ID)
    return user, token, nil
}
```

**3. CAPTCHA (Human Verification)**:

- Add Google reCAPTCHA v3 to login form
- Verify token on backend before processing login
- Trigger after 3 failed attempts

**4. Slow Hashing (Computational Defense)**:

```go
import "golang.org/x/crypto/bcrypt"

// bcrypt automatically adds salt, configurable cost
const bcryptCost = 12 // Higher = slower (good for defense)

func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
    return string(hash), err
}
```

**5. IP Blacklisting (Automated)**:

```go
// Track failed attempts per IP in Redis
// Ban IP after 50 failed attempts across all accounts
if failedAttemptsFromIP > 50 {
    redis.Set("blacklist:"+ip, "true", 24*time.Hour)
}
```

**6. Monitoring \& Alerting**:

```go
// Alert security team if:
// - >100 failed logins in 5 minutes (automated attack)
// - Same IP hits multiple accounts
// - Unusual geographic patterns
```

**7. Multi-Factor Authentication (Best Long-Term Solution)**:

- Require second factor (TOTP, SMS, email code)
- Even if password is compromised, attacker can't access account

**Architecture Changes**:

- **Firewall**: Use AWS WAF / Cloudflare to block malicious IPs before reaching app
- **Distributed Rate Limiting**: Use Redis for rate limiting across multiple backend instances
- **Audit Logging**: Log all auth attempts to separate service for forensic analysis

***

### 7. **Scenario 2: Handle Millions of Requests/Sec, Fault-Tolerant**

**Current Bottlenecks**:

1. Single server instance
2. Single database (SPOF - Single Point of Failure)
3. No caching
4. Synchronous request processing

**Scalability Architecture**:

**1. Horizontal Scaling (Kubernetes)**:

```yaml
# k8s/backend-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth-backend
spec:
  replicas: 10  # Start with 10 pods
  selector:
    matchLabels:
      app: auth-backend
  template:
    metadata:
      labels:
        app: auth-backend
    spec:
      containers:
      - name: backend
        image: auth-backend:latest
        resources:
          requests:
            cpu: "500m"
            memory: "512Mi"
          limits:
            cpu: "1000m"
            memory: "1Gi"
---
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: auth-backend-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: auth-backend
  minReplicas: 10
  maxReplicas: 100
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

**2. Database Scaling**:

**Read Replicas**:

```go
// Separate read/write database connections
type DBConfig struct {
    Primary  *sql.DB  // For writes (INSERT, UPDATE)
    Replicas []*sql.DB  // For reads (SELECT)
}

func (db *DBConfig) GetUser(id string) (*User, error) {
    // Route reads to replica
    replica := db.Replicas[rand.Intn(len(db.Replicas))]
    return queryUser(replica, id)
}

func (db *DBConfig) CreateUser(user *User) error {
    // Route writes to primary
    return insertUser(db.Primary, user)
}
```

**Connection Pooling**:

```go
import "github.com/jackc/pgx/v5/pgxpool"

pool, err := pgxpool.New(ctx, "postgres://...")
// Configure pool
pool.Config().MaxConns = 100
pool.Config().MinConns = 10
pool.Config().MaxConnLifetime = time.Hour
```

**3. Caching Layer (Redis)**:

```go
// Cache user profiles after first load
func (s *UserService) GetProfile(userID string) (*User, error) {
    // Check cache first
    cached, err := s.redis.Get("user:"+userID).Result()
    if err == nil {
        return deserializeUser(cached), nil
    }
    
    // Cache miss - load from DB
    user, err := s.repo.GetUser(userID)
    if err != nil {
        return nil, err
    }
    
    // Store in cache (TTL: 5 minutes)
    s.redis.Set("user:"+userID, serialize(user), 5*time.Minute)
    return user, nil
}
```

**4. Load Balancer**:

```yaml
# k8s/ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: auth-ingress
  annotations:
    nginx.ingress.kubernetes.io/limit-rps: "1000"  # Rate limit per IP
spec:
  rules:
  - host: api.auth-app.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: auth-backend
            port:
              number: 8080
```

**5. Asynchronous Processing**:

```go
// Send welcome email asynchronously
func (s *AuthService) SignUp(email, password string) error {
    user, err := s.createUser(email, password)
    if err != nil {
        return err
    }
    
    // Queue email job instead of blocking
    s.queue.Publish("emails", WelcomeEmailJob{
        UserID: user.ID,
        Email:  user.Email,
    })
    
    return nil
}
```

**6. Circuit Breaker Pattern**:

```go
import "github.com/sony/gobreaker"

var cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
    Name:        "DB",
    MaxRequests: 3,
    Timeout:     10 * time.Second,
})

func (r *UserRepository) GetUser(id string) (*User, error) {
    result, err := cb.Execute(func() (interface{}, error) {
        return r.db.QueryUser(id)
    })
    
    if err != nil {
        // Circuit open - return cached data or error
        return nil, err
    }
    
    return result.(*User), nil
}
```

**7. Database Sharding (Long-Term)**:

```go
// Shard users by ID hash
func (db *ShardedDB) GetShardForUser(userID string) *sql.DB {
    hash := crc32.ChecksumIEEE([]byte(userID))
    shardIndex := hash % uint32(len(db.Shards))
    return db.Shards[shardIndex]
}
```

**8. CDN for Frontend**:

- Serve React build from CloudFront/Cloudflare
- Reduce latency globally

**9. Observability**:

```go
import "go.opentelemetry.io/otel"

// Distributed tracing
ctx, span := tracer.Start(ctx, "SignIn")
defer span.End()

// Metrics
signInAttempts.Inc()
signInLatency.Observe(duration.Seconds())
```

**10. Chaos Engineering**:

- Regularly test failure scenarios (pod crashes, DB failures)
- Use tools like Chaos Mesh to inject faults

**Performance Targets**:

- **Throughput**: 1M req/sec (1000 pods × 1000 req/sec each)
- **Latency**: p99 < 100ms
- **Availability**: 99.99% uptime (52 minutes downtime/year)

***

## Implementation Timeline (4 Days)

### Day 1 (Nov 21): Backend Core

- [x] Setup Go project structure
- [x] Implement user model, repository
- [x] JWT authentication service
- [x] Signup/signin endpoints
- [x] Protected `/me` endpoint
- [x] PostgreSQL integration
- [x] Basic unit tests


### Day 2 (Nov 22): Frontend Core

- [x] Setup React + TypeScript + Vite
- [x] Auth context \& state management
- [x] SignUp/SignIn forms with validation
- [x] Protected routes
- [x] Profile page
- [x] API client with interceptors


### Day 3 (Nov 23): DevOps \& Integration

- [x] Docker Compose for local dev
- [x] Kubernetes manifests
- [x] Makefile with common commands
- [x] OpenAPI spec
- [x] README with setup instructions
- [x] Integration testing


### Day 4 (Nov 24): Polish \& Report

- [x] Code cleanup, comments
- [x] Final testing
- [x] Write comprehensive report
- [x] Deploy to public GitHub
- [x] Record demo video (optional)
- [x] Submit by 23:59 AEDT

***

## Makefile Commands

```makefile
.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $1, $2}'

.PHONY: install
install: ## Install dependencies
	cd backend && go mod download
	cd frontend && npm install

.PHONY: dev
dev: ## Run development environment
	docker-compose up --build

.PHONY: test
test: ## Run tests
	cd backend && go test ./... -v
	cd frontend && npm test

.PHONY: build
build: ## Build Docker images
	docker build -t auth-backend:latest ./backend
	docker build -t auth-frontend:latest ./frontend

.PHONY: k8s-deploy
k8s-deploy: ## Deploy to Kubernetes
	kubectl apply -f k8s/

.PHONY: migrate
migrate: ## Run database migrations
	docker exec -i postgres psql -U auth -d authdb < backend/migrations/001_create_users_table.sql

.PHONY: gen-api
gen-api: ## Generate API code from OpenAPI spec
	oapi-codegen -package api -generate types,server backend/api/openapi.yaml > backend/internal/api/types.go
	cd frontend && npm run generate-api

.PHONY: clean
clean: ## Clean up containers and volumes
	docker-compose down -v
```


***

## Key Golang Packages

```go
// go.mod
module github.com/yourusername/auth-app

go 1.21

require (
    github.com/gin-gonic/gin v1.10.0           // HTTP framework
    github.com/golang-jwt/jwt/v5 v5.2.0        // JWT tokens
    github.com/jackc/pgx/v5 v5.5.0             // PostgreSQL driver
    github.com/google/uuid v1.6.0              // UUID generation
    golang.org/x/crypto v0.18.0                // bcrypt password hashing
    github.com/go-playground/validator/v10     // Input validation
    github.com/joho/godotenv v1.5.1            // Environment variables
    github.com/rs/cors v1.10.1                 // CORS middleware
)
```


***

## Security Best Practices Implemented

1. **Password Hashing**: bcrypt with cost factor 12
2. **JWT Tokens**: HS256 algorithm, 24-hour expiry
3. **CORS**: Restricted to frontend origin only
4. **Input Validation**: Email format, password strength
5. **SQL Injection Prevention**: Parameterized queries (pgx)
6. **Error Messages**: Generic messages (don't leak "user exists")
7. **HTTPS**: Enforced in production (k8s ingress with TLS)
8. **Environment Variables**: Secrets never committed to Git

***

## Report Structure (To Submit)

```markdown
# Rakuten Symphony Full Stack Engineer - Technical Report

## Candidate Information
- Name: Jason [Your Last Name]
- Email: [Your Email]
- GitHub: https://github.com/yourusername/auth-app
- Date: November 24, 2025

## 1. Setup and Architectural Choices
[Detailed explanation with diagrams]

## 2. Potential Weaknesses & Production Improvements
[10 weaknesses with concrete solutions]

## 3. Future Improvements (More Time)
[Priority-ordered feature list]

## 4. Frontend State Management & Validation
[Why React Context + React Hook Form + Zod]

## 5. Types & Contracts: Frontend ↔ Backend Sync
[OpenAPI-first design approach]

## 6. Scenario 1: Brute-Force Attack Mitigation
[Rate limiting + account lockout + MFA strategy]

## 7. Scenario 2: Scale to Millions of Requests/Sec
[Kubernetes HPA + DB replicas + Redis caching + load balancing]

## Appendix
- API Documentation (OpenAPI spec)
- Database Schema
- Deployment Instructions
- Test Coverage Report
```


***

## Next Steps

1. **Clone Starter Template**: I can generate the complete codebase structure
2. **Implement Backend**: Follow clean architecture pattern
3. **Implement Frontend**: React + TypeScript with proper type safety
4. **Docker Setup**: docker-compose for local dev
5. **Kubernetes**: Basic deployment manifests
6. **Testing**: Unit tests for critical paths
7. **Documentation**: Comprehensive README
8. **Report**: Answer all 7 questions thoroughly

Would you like me to:

1. **Generate the complete codebase** with all files?
2. **Create a detailed step-by-step implementation guide**?
3. **Generate the report template** with your specific answers?

Let me know which part you'd like to start with, and I'll create the actual code!
<span style="display:none">[^1_1][^1_10][^1_11][^1_12][^1_13][^1_14][^1_15][^1_16][^1_17][^1_18][^1_19][^1_2][^1_20][^1_21][^1_22][^1_23][^1_24][^1_3][^1_4][^1_5][^1_6][^1_7][^1_8][^1_9]</span>

<div align="center">⁂</div>

[^1_1]: image.jpg

[^1_2]: image.jpg

[^1_3]: image.jpg

[^1_4]: image.jpg

[^1_5]: https://www.ntia.gov/sites/default/files/publications/rakuten_symphony.pdf

[^1_6]: https://www.linkedin.com/posts/vinita-718767224_my-interview-experience-at-rakuten-symphony-activity-7225573193929940992-8jCw

[^1_7]: https://www.browserstack.com/guide/rest-api-in-golang

[^1_8]: https://leadiq.com/c/rakuten-symphony/5eea78162fd07506cb0c4b4d

[^1_9]: https://www.youtube.com/watch?v=Py4ClrZ1g-4

[^1_10]: https://stackoverflow.blog/2021/10/06/best-practices-for-authentication-and-authorization-for-rest-apis/

[^1_11]: https://symphony.rakuten.com/blog/rakuten-symphony-marketing-transformation-for-an-industry-in-transition

[^1_12]: https://www.reddit.com/r/developersIndia/comments/1fb2pbe/rejected_after_final_round_interview_feeling/

[^1_13]: https://zuplo.com/learning-center/api-authentication

[^1_14]: https://symphony.rakuten.com

[^1_15]: https://nodeflair.com/companies/rakuten-symphony/interviews

[^1_16]: https://developer.auth0.com/resources/guides/api/standard-library/basic-authorization

[^1_17]: https://technologymagazine.com/articles/solving-the-puzzle-of-scaled-cloud-capabilities-at-the-edge

[^1_18]: https://blog.stackademic.com/golang-use-restful-api-like-a-pro-best-practices-for-restful-apis-in-golang-39d808fbbdb6

[^1_19]: https://symphony.rakuten.com/blog/from-cloud-to-cognition-the-rise-of-the-ai-native-telco

[^1_20]: https://dev.to/githaiga22/mastering-modern-apis-with-go-from-basics-to-best-practices-2obn

[^1_21]: https://omdia.tech.informa.com/om022142/rakuten-symphony-reimagineering-telecoms

[^1_22]: https://refold.ai/blog/9-best-practices-for-api-authentication-and-authorization-2/

[^1_23]: https://symphony.rakuten.com/blog/rakuten-symphony-and-ciq-bring-back-open-source-and-open-community-with-rocky-linux

[^1_24]: https://cristiancurteanu.com/5-api-design-patterns-in-go-that-solve-your-biggest-problems-2025/


---

# based on our chat evolution and progression, can we research for the best encription solution for DB in particular using it in our case?

Based on the research and our authentication app context, here's a comprehensive analysis of the best encryption solutions for your database:

# Database Encryption Strategy for Authentication App

## Three-Layer Encryption Approach

For your Rakuten Symphony authentication application, I recommend a **layered security approach** combining multiple encryption strategies:

***

## 1. Password Hashing (CRITICAL - Application Layer)

### **Recommendation: Argon2id > bcrypt**

**Why Argon2id over bcrypt**:

- **OWASP Standard**: As of 2024, OWASP recommends Argon2id as the **primary** choice[^2_1]
- **Memory-hard**: Resistant to GPU/ASIC attacks (bcrypt is CPU-only)
- **Configurable**: Memory, time, and parallelism parameters
- **Future-proof**: Winner of Password Hashing Competition

**Implementation in Golang**:

```go
package password

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "strings"
    
    "golang.org/x/crypto/argon2"
)

// Argon2id parameters (recommended for 2024)
const (
    // Memory: 64 MB (65536 KB)
    Memory      = 64 * 1024
    // Iterations: 3 (balance between security and UX)
    Iterations  = 3
    // Parallelism: 2 threads
    Parallelism = 2
    // Salt length: 16 bytes
    SaltLength  = 16
    // Key length: 32 bytes
    KeyLength   = 32
)

// GenerateFromPassword creates an Argon2id hash of the password
func GenerateFromPassword(password string) (string, error) {
    // Generate cryptographically secure salt
    salt := make([]byte, SaltLength)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    
    // Generate hash
    hash := argon2.IDKey([]byte(password), salt, Iterations, Memory, Parallelism, KeyLength)
    
    // Encode as: $argon2id$v=19$m=65536,t=3,p=2$salt$hash
    b64Salt := base64.RawStdEncoding.EncodeToString(salt)
    b64Hash := base64.RawStdEncoding.EncodeToString(hash)
    
    encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
        argon2.Version, Memory, Iterations, Parallelism, b64Salt, b64Hash)
    
    return encodedHash, nil
}

// ComparePasswordAndHash verifies a password against an Argon2id hash
func ComparePasswordAndHash(password, encodedHash string) (bool, error) {
    // Parse encoded hash
    parts := strings.Split(encodedHash, "$")
    if len(parts) != 6 {
        return false, errors.New("invalid hash format")
    }
    
    var memory, iterations uint32
    var parallelism uint8
    _, err := fmt.Sscanf(parts[^2_3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
    if err != nil {
        return false, err
    }
    
    salt, err := base64.RawStdEncoding.DecodeString(parts[^2_4])
    if err != nil {
        return false, err
    }
    
    hash, err := base64.RawStdEncoding.DecodeString(parts[^2_5])
    if err != nil {
        return false, err
    }
    
    // Generate hash with same parameters
    compareHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, uint32(len(hash)))
    
    // Constant-time comparison to prevent timing attacks
    return subtle.ConstantTimeCompare(hash, compareHash) == 1, nil
}
```

**Performance Characteristics**:[^2_2]

- **Hash time**: ~50-100ms per password (tunable)
- **Memory usage**: 64MB per hash operation
- **CPU**: 100% during hashing (expected behavior)
- **Mitigation**: Use goroutines, handle concurrently

**Why NOT bcrypt for new projects**:

- OWASP: "bcrypt should **only** be used for password storage in legacy systems"[^2_1]
- 72-byte input limit (truncates longer passwords)
- Not memory-hard (vulnerable to GPU attacks)

***

## 2. Sensitive Field Encryption (Application Layer)

### **Recommendation: pgsodium > pgcrypto**

For encrypting specific database columns (e.g., email, PII), use **pgsodium** instead of pgcrypto.

**Why pgsodium**:[^2_3][^2_4][^2_5]


| Feature | pgcrypto | pgsodium |
| :-- | :-- | :-- |
| **Maintenance** | Outdated, "serious red flags"[^2_5] | Modern, actively maintained |
| **Library** | Custom implementations | libsodium (industry standard) |
| **Key Management** | Manual | Server Key Management built-in |
| **Performance** | Slower | Faster (libsodium optimized) |
| **Features** | Basic | Public key signing, sealed boxes, key derivation |
| **Vault Integration** | No | Yes (remote key management) |

**pgsodium Architecture**:

```sql
-- 1. Install extension
CREATE EXTENSION pgsodium;

-- 2. Create encrypted table with Transparent Column Encryption
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT,  -- Encrypted column
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Create security barrier view
CREATE VIEW users_decrypted WITH (security_barrier) AS
SELECT 
    id,
    pgsodium.decrypt(
        decode(email, 'base64'),
        pgsodium.crypto_secretbox_open_key()
    )::TEXT AS email,
    password_hash,
    created_at
FROM users;

-- 4. Use triggers for automatic encryption
CREATE OR REPLACE FUNCTION encrypt_user_email()
RETURNS TRIGGER AS $$
BEGIN
    NEW.email := encode(
        pgsodium.crypto_secretbox(
            NEW.email::bytea,
            pgsodium.crypto_secretbox_new_key()
        ),
        'base64'
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER encrypt_email_trigger
BEFORE INSERT OR UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION encrypt_user_email();
```

**Key Management with pgsodium + Vault**:[^2_4]

```bash
# 1. Generate master key and store in Vault
vault kv put secret/pgsodium/master-key key=$(openssl rand -base64 32)

# 2. Configure pgsodium to fetch key from Vault
# In postgresql.conf
pgsodium.getkey_script = '/usr/local/bin/get_pgsodium_key.sh'
```

```bash
#!/bin/bash
# /usr/local/bin/get_pgsodium_key.sh
vault kv get -field=key secret/pgsodium/master-key
```

**Benefits**:

- Encryption key NEVER stored in database or SQL
- Application code unchanged (transparent encryption)
- Key rotation supported without re-encrypting all data
- Meets compliance requirements (GDPR, HIPAA)

***

## 3. Full Database Encryption (Infrastructure Layer)

### **Recommendation: Transparent Data Encryption (TDE) + Disk Encryption**

**For Production: Multi-Layer Defense**

#### Option A: PostgreSQL TDE (pg_tde)[^2_6][^2_7][^2_8]

**Percona pg_tde** (open-source extension):

```sql
-- Install extension
CREATE EXTENSION pg_tde;

-- Configure keyring (Vault integration)
SELECT pg_tde_add_vault_key_provider(
    'vault-provider',
    'https://vault.example.com:8200',
    'vault-token',
    '/v1/transit/keys/pg-tde-key'
);

-- Create encrypted database
CREATE DATABASE authdb WITH ENCRYPTION USING VAULT KEY;
```

**Performance Impact**:[^2_9]

- **Read-only workloads**: ~5% overhead (negligible)
- **Read-write workloads**: 30-50% overhead (significant)
- **Mitigation**: Increase `shared_buffers` to reduce disk I/O

**When to use**:

- Compliance requirements (PCI DSS, HIPAA)
- Protecting backups and dumps
- Defense against physical disk theft


#### Option B: Disk-Level Encryption (Simpler)

```bash
# LUKS encryption (Linux)
cryptsetup luksFormat /dev/sdb
cryptsetup luksOpen /dev/sdb pgdata_encrypted
mkfs.ext4 /dev/mapper/pgdata_encrypted
mount /dev/mapper/pgdata_encrypted /var/lib/postgresql/data
```

**Benefits**:

- Zero application changes
- No performance impact on queries
- OS-level protection

**Limitations**:

- Doesn't protect against SQL injection
- Doesn't encrypt in-memory data
- DBA with OS access can read data

***

## Recommended Architecture for Your Auth App

### **Layered Encryption Strategy**

```
┌─────────────────────────────────────────────────────────────┐
│                     Application Layer                        │
│  • Argon2id password hashing (Golang)                       │
│  • Input validation & sanitization                          │
│  • JWT token encryption (HS256)                             │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                   Database Column Layer                      │
│  • pgsodium for email encryption                            │
│  • Server Key Management (keys in Vault)                    │
│  • Transparent encryption/decryption via views              │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                    Transport Layer (TLS)                     │
│  • PostgreSQL SSL/TLS (SCRAM-SHA-256 auth)                 │
│  • Backend ↔ DB: TLS 1.3 certificate-based                 │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                  Storage Layer (Optional)                    │
│  • Disk encryption (LUKS/dm-crypt)                          │
│  • Cloud provider encryption (AWS KMS, GCP CMEK)            │
└─────────────────────────────────────────────────────────────┘
```


***

## Implementation for Your Assignment

### **Phase 1: Minimum Viable Security (4-day deadline)**

```go
// models/user.go
type User struct {
    ID           uuid.UUID  `json:"id" db:"id"`
    Email        string     `json:"email" db:"email"`
    PasswordHash string     `json:"-" db:"password_hash"`  // Never expose in JSON
    CreatedAt    time.Time  `json:"created_at" db:"created_at"`
}

// service/auth_service.go
import (
    "golang.org/x/crypto/argon2"
)

type AuthService struct {
    repo repository.UserRepository
}

func (s *AuthService) SignUp(email, password string) (*User, error) {
    // Hash password with Argon2id
    passwordHash, err := hashPassword(password)
    if err != nil {
        return nil, err
    }
    
    user := &User{
        Email:        email,
        PasswordHash: passwordHash,
    }
    
    return s.repo.Create(user)
}

func (s *AuthService) SignIn(email, password string) (*User, string, error) {
    user, err := s.repo.GetByEmail(email)
    if err != nil {
        return nil, "", errors.New("invalid credentials")
    }
    
    // Verify password
    if !verifyPassword(password, user.PasswordHash) {
        return nil, "", errors.New("invalid credentials")
    }
    
    // Generate JWT token
    token, err := s.jwt.Generate(user.ID)
    if err != nil {
        return nil, "", err
    }
    
    return user, token, nil
}

func hashPassword(password string) (string, error) {
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }
    
    hash := argon2.IDKey([]byte(password), salt, 3, 64*1024, 2, 32)
    
    // Encode as $argon2id$...
    return encodeHash(salt, hash), nil
}
```

**PostgreSQL Setup**:

```sql
-- Minimal secure setup
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Force SSL connections
-- In postgresql.conf
ssl = on
ssl_cert_file = '/etc/ssl/certs/server.crt'
ssl_key_file = '/etc/ssl/private/server.key'

-- In pg_hba.conf (force SSL + SCRAM-SHA-256)
hostssl all all 0.0.0.0/0 scram-sha-256
```


### **Phase 2: Production-Ready (Post-Interview)**

Add pgsodium + Vault integration:

```bash
# docker-compose.yml additions
services:
  vault:
    image: vault:latest
    environment:
      VAULT_DEV_ROOT_TOKEN_ID: dev-token
    ports:
      - "8200:8200"
  
  postgres:
    image: postgres:16
    environment:
      POSTGRES_DB: authdb
      POSTGRES_USER: auth
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
    command: |
      -c ssl=on
      -c ssl_cert_file=/etc/ssl/certs/server.crt
      -c ssl_key_file=/etc/ssl/private/server.key
```


***

## Security Comparison Matrix

| Threat | Argon2id | pgsodium | TDE | Disk Encryption |
| :-- | :-- | :-- | :-- | :-- |
| **SQL Injection** | ✅ Yes | ✅ Yes | ❌ No | ❌ No |
| **Database Dump Theft** | ✅ Yes | ✅ Yes | ✅ Yes | ✅ Yes |
| **Memory Dump** | ✅ Yes | ⚠️ Partial | ❌ No | ❌ No |
| **Disk Theft** | ✅ Yes | ✅ Yes | ✅ Yes | ✅ Yes |
| **Malicious DBA** | ✅ Yes | ✅ Yes* | ❌ No | ❌ No |
| **Brute Force** | ✅ Yes | N/A | N/A | N/A |

*With Vault integration (key not in DB)

***

## Report Answer: Database Encryption

**For your report, answer the encryption question as follows**:

> **Database Encryption Strategy**
>
> We implement a three-layer encryption approach:
>
> 1. **Password Hashing** (Application Layer): Argon2id instead of bcrypt, following OWASP 2024 recommendations. Parameters: 64MB memory, 3 iterations, 2 threads. Provides ~50ms hash time, resistant to GPU attacks.
>
> 2. **Transport Encryption**: PostgreSQL SSL/TLS with SCRAM-SHA-256 authentication. Prevents man-in-the-middle attacks and password sniffing.
>
> 3. **At-Rest Encryption** (Production): Disk-level encryption (LUKS/dm-crypt) for basic protection. For higher compliance requirements, pg_tde extension with Vault key management would be added.
>
> **Why Argon2id over bcrypt**: Memory-hard algorithm resistant to specialized hardware attacks, configurable security parameters, and recommended as primary choice by OWASP. bcrypt limited to 72-byte passwords and only CPU-hard.
>
> **Production Enhancement**: pgsodium extension for transparent column encryption of sensitive fields (email, PII) with keys stored in HashiCorp Vault, never accessible via SQL. Supports key rotation without re-encrypting data.

***

## Golang Package Dependencies

```go
// go.mod additions
require (
    golang.org/x/crypto v0.18.0  // Argon2id
    github.com/jackc/pgx/v5 v5.5.0  // PostgreSQL with SSL support
    github.com/hashicorp/vault/api v1.10.0  // Vault integration (optional)
)
```


***

## Key Takeaways

1. **Use Argon2id** for password hashing (NOT bcrypt in 2024)
2. **Add pgsodium** if encrypting sensitive fields beyond passwords
3. **Always use TLS** for PostgreSQL connections
4. **Disk encryption** is table-stakes for production
5. **Vault integration** for enterprise-grade key management

This approach balances security, performance, and implementation complexity perfectly for your 4-day deadline while showing understanding of production-grade security.
<span style="display:none">[^2_10][^2_11][^2_12][^2_13][^2_14][^2_15][^2_16][^2_17][^2_18][^2_19][^2_20][^2_21][^2_22][^2_23][^2_24][^2_25][^2_26][^2_27][^2_28][^2_29][^2_30][^2_31][^2_32][^2_33][^2_34][^2_35][^2_36][^2_37]</span>

<div align="center">⁂</div>

[^2_1]: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html

[^2_2]: https://www.reddit.com/r/golang/comments/1falf4v/argonbcrypt_takes_100_cpu_while_crypt_user/

[^2_3]: https://www.slideshare.net/slideshow/pgsodiums-features-those-not-provided-by-pgcrypto-and-integration-with-remote-key-management-services-240380390/240380390

[^2_4]: https://github.com/michelp/pgsodium

[^2_5]: https://news.ycombinator.com/item?id=23476248

[^2_6]: https://www.percona.com/blog/transparent-data-encryption-for-postgresql-release-candidate-is-here/

[^2_7]: https://www.percona.com/blog/protect-your-postgresql-database-with-pg_tde-safe-and-secure/

[^2_8]: https://www.percona.com/blog/webinar-qa-no-more-workarounds-open-source-postgresql-tde-is-here/

[^2_9]: https://www.cybertec-postgresql.com/en/postgresql-tde-performance/

[^2_10]: https://www.crunchydata.com/blog/data-encryption-in-postgres-a-guidebook

[^2_11]: https://severalnines.com/blog/best-practices-for-securing-postgresql-in-hybrid-environments/

[^2_12]: https://dev.to/nikita_rykhlov/go-tools-password-hashing-with-argon2-instead-of-bcrypt-38aj

[^2_13]: https://dev.to/arctype/a-guide-to-transparent-data-encryption-in-postgresql-3elh

[^2_14]: https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go

[^2_15]: https://www.enterprisedb.com/blog/everything-need-know-postgres-data-encryption

[^2_16]: https://www.tigerdata.com/learn/postgres-security-best-practices

[^2_17]: https://docs.aws.amazon.com/dms/latest/sql-server-to-aurora-postgresql-migration-playbook/chap-sql-server-aurora-pg.security.transparentdataencryption.html

[^2_18]: https://www.enterprisedb.com/blog/how-to-secure-postgresql-security-hardening-best-practices-checklist-tips-encryption-authentication-vulnerabilities

[^2_19]: https://stackoverflow.com/questions/64220922/bcrypt-vs-argon2-and-their-hashing-algorithms

[^2_20]: https://stackoverflow.com/questions/4002184/database-encryption-or-application-level-encryption

[^2_21]: https://www.reddit.com/r/PostgreSQL/comments/1fhz832/customer_asks_if_the_postgresql_database_can_be/

[^2_22]: https://www.pingidentity.com/en/resources/blog/post/encryption-vs-hashing-vs-salting.html

[^2_23]: https://learn.microsoft.com/en-us/azure/azure-sql/database/transparent-data-encryption-tde-overview?view=azuresql

[^2_24]: https://blog.secureflag.com/2024/06/05/security-hardening-for-your-postgresql-database/

[^2_25]: https://snyk.io/blog/secure-password-hashing-in-go/

[^2_26]: https://ngublag.com/posts/2023-05-07-vault-with-go.html

[^2_27]: https://github.com/zalando/postgres-operator/issues/847

[^2_28]: https://developer.hashicorp.com/vault/docs/configuration/storage/postgresql

[^2_29]: https://splitmind.dev/posts/generate-creds-postgres-vault-with-golang/

[^2_30]: https://rizqimulki.com/database-encryption-at-rest-performance-vs-security-trade-offs-f66698feef45

[^2_31]: https://dev.to/breda/dynamic-postgresql-credentials-using-hashicorp-vault-with-php-symfony-go-examples-4imj

[^2_32]: https://cpl.thalesgroup.com/encryption/selecting-right-encryption-approach

[^2_33]: https://www.reddit.com/r/PostgreSQL/comments/yce0jx/pgcrypto_question/

[^2_34]: https://www.technowizardry.net/2025/09/vault-postgres-home-lab/

[^2_35]: https://www.highgo.ca/2019/09/30/the-transparent-data-encryption-in-postgresql/

[^2_36]: https://www.scalingpostgres.com/episodes/247-generate-test-data-faster-archiving-date-statistics-useless-indexes/

[^2_37]: https://www.reddit.com/r/hashicorp/comments/1cx38aw/vault_postgres_database_secrets_engine_performance/


---

# based on our chat evolution and progression, please suggest a repository name for this project properly. I'm going to create this repo as requested now.

Based on the context (Rakuten Symphony Full Stack Engineer assignment, authentication app with Golang backend + React frontend), here are my repository name recommendations:

## Primary Recommendation

```
rakuten-fullstack-auth-challenge
```

**Why this works best**:

- ✅ Company-specific context (shows it's for Rakuten Symphony)
- ✅ Describes tech stack (full-stack)
- ✅ Purpose-clear (authentication)
- ✅ Professional naming convention (kebab-case)
- ✅ GitHub-friendly (descriptive, searchable)
- ✅ Interview context clear (challenge/assignment)


## Alternative Options (Ranked)

### Option 2: More Technical Focus

```
auth-api-react-golang
```

- Emphasizes tech stack
- Shorter, technical
- Good for portfolio


### Option 3: Descriptive + Professional

```
secure-auth-platform
```

- Generic, reusable
- Production-sounding
- Could be portfolio piece beyond interview


### Option 4: Assignment-Specific

```
rakuten-symphony-assignment-2024
```

- Very explicit about context
- Easy for HR/reviewers to identify
- Date-stamped for reference


### Option 5: Feature-Focused

```
jwt-auth-fullstack-app
```

- Highlights key technology (JWT)
- Tech-forward naming
- Portfolio-friendly


## My Top Pick for You

I recommend:

```
rakuten-fullstack-auth-challenge
```

**Repository Description (for GitHub)**:

```
Full-stack authentication application with Go backend (REST API, JWT, PostgreSQL) 
and React TypeScript frontend. Built for Rakuten Symphony technical assessment.
```

**Topics/Tags** (add these in GitHub repo settings):

```
golang, react, typescript, authentication, jwt, postgresql, docker, kubernetes, rest-api, fullstack
```


## Repository Initialization Commands

```bash
# Create the repository
gh repo create rakuten-fullstack-auth-challenge --public --description "Full-stack authentication app with Go + React + PostgreSQL for Rakuten Symphony technical assessment"

# Or if you prefer using Git directly
mkdir rakuten-fullstack-auth-challenge
cd rakuten-fullstack-auth-challenge
git init
git remote add origin git@github.com:YOUR_USERNAME/rakuten-fullstack-auth-challenge.git

# Add initial README
cat > README.md << 'EOF'
# Rakuten Symphony Full-Stack Authentication Challenge

Full-stack authentication application built with modern technologies for the Rakuten Symphony technical assessment.

## Tech Stack

**Backend:**
- Go 1.21+ (Gin framework)
- PostgreSQL 16
- JWT authentication
- Argon2id password hashing
- OpenAPI 3.0 specification

**Frontend:**
- React 18 + TypeScript
- React Hook Form + Zod validation
- React Router v6
- Axios with interceptors

**DevOps:**
- Docker + Docker Compose
- Kubernetes manifests
- Makefile automation
- GitHub Actions CI/CD

## Quick Start

```


# Clone repository

git clone https://github.com/YOUR_USERNAME/rakuten-fullstack-auth-challenge.git
cd rakuten-fullstack-auth-challenge

# Start development environment

make dev

# Access application

# Frontend: http://localhost:3000

# Backend API: http://localhost:8080

# API Docs: http://localhost:8080/swagger

```

## Assignment Requirements

Built to fulfill the following requirements:
- ✅ REST API for signup/signin/protected profile endpoint
- ✅ PostgreSQL database persistence
- ✅ React SPA with authentication flow
- ✅ Docker containerization
- ✅ Kubernetes deployment manifests
- ✅ Comprehensive documentation

## Report

See [REPORT.md](./REPORT.md) for detailed architectural decisions, security considerations, and scalability analysis.

## Author

**Jason [Your Last Name]**  
Built for Rakuten Symphony Full Stack Engineer position  
November 2024

## License

MIT
EOF

# Create basic .gitignore
cat > .gitignore << 'EOF'
# Backend
backend/bin/
backend/*.exe
backend/*.log

# Frontend
frontend/node_modules/
frontend/dist/
frontend/build/
frontend/.env.local

# IDE
.vscode/
.idea/
*.swp
*.swo

# Environment
.env
.env.local
*.pem
*.key

# OS
.DS_Store
Thumbs.db

# Docker
docker-compose.override.yml
EOF

# Initial commit
git add .
git commit -m "Initial commit: Project structure and README"
git branch -M main
git push -u origin main
```


## Pro Tips

1. **Keep it public** - Shows confidence and transparency to reviewers
2. **Add LICENSE** - MIT or Apache 2.0 (shows professionalism)
3. **Professional commit messages** - Use conventional commits format:

```
feat: implement JWT authentication service
docs: add API documentation and usage examples
fix: resolve CORS issue in production build
test: add unit tests for auth handlers
```

4. **README badges** (add after CI setup):

```markdown
![Build Status](https://github.com/YOUR_USERNAME/rakuten-fullstack-auth-challenge/workflows/CI/badge.svg)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB?logo=react)
```

5. **Add topics in GitHub UI**:
    - Go to repository → Settings → About (Edit)
    - Add: `golang`, `react`, `authentication`, `jwt`, `postgresql`, `docker`, `kubernetes`

Would you like me to generate the complete initial repository structure with all files ready for you to push?
