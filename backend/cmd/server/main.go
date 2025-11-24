package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/nfsarch33/secure-auth-platform/backend/docs" // Swagger docs
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api/handlers"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api/middleware"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/repository/postgres"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/service"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/jwt"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/recaptcha"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Rakuten Symphony Auth API
// @version 1.0
// @description Authentication API for Rakuten Symphony assignment
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := run(); err != nil {
		slog.Error("Application exited with error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		slog.Warn("Error loading .env file, using system environment")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return os.ErrInvalid // Or a custom error
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		if os.Getenv("GIN_MODE") == "release" {
			slog.Error("JWT_SECRET must be set in release mode")
			return os.ErrInvalid
		}
		slog.Warn("JWT_SECRET is not set, using default for development")
		secretKey = "dev-secret-key" // #nosec G101
	}

	userRepo := postgres.NewPostgresUserRepository(pool)
	tokenService := jwt.NewTokenService(secretKey, "secure-auth-platform", 24*time.Hour)
	authService := service.NewAuthService(userRepo, tokenService)

	recaptchaSecret := os.Getenv("RECAPTCHA_SECRET_KEY")
	if recaptchaSecret == "" {
		slog.Warn("RECAPTCHA_SECRET_KEY is not set, using dummy value for development")
		recaptchaSecret = "dummy-secret"
	}
	recaptchaDisabled := os.Getenv("RECAPTCHA_DISABLED") == "true"
	recaptchaVerifier := recaptcha.NewVerifier(recaptchaSecret, recaptchaDisabled)

	authHandler := handlers.NewAuthHandler(authService, recaptchaVerifier)

	r := gin.Default()

	// Middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RateLimitMiddleware(60)) // 60 requests per minute
	r.Use(middleware.SecureHeadersMiddleware())

	// Health Check
	r.GET("/health", handlers.HealthCheck)

	// Swagger
	r.StaticFile("/openapi.yaml", "api/openapi.yaml")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/openapi.yaml")))

	// Routes
	apiGroup := r.Group("/api")
	{
		auth := apiGroup.Group("/auth")
		{
			auth.POST("/signup", authHandler.SignUp)
			auth.POST("/signin", authHandler.SignIn)
			auth.POST("/signout", authHandler.SignOut)
		}

		// Protected routes
		protected := apiGroup.Group("/")
		protected.Use(middleware.AuthMiddleware(tokenService))
		{
			protected.GET("/me", authHandler.GetMe)
		}
	}

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
		return err
	}

	slog.Info("Server exiting")
	return nil
}
