package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api/handlers"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api/middleware"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/repository/postgres"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/service"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/jwt"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/recaptcha"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/nfsarch33/secure-auth-platform/backend/docs" // Swagger docs
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
	if err := godotenv.Load(); err != nil {
		slog.Warn("Error loading .env file, using system environment")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		slog.Error("DATABASE_URL is required")
		os.Exit(1)
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		slog.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "super-secret-key-change-in-production"
	}

	userRepo := postgres.NewPostgresUserRepository(pool)
	tokenService := jwt.NewTokenService(secretKey, "secure-auth-platform", 24*time.Hour)
	authService := service.NewAuthService(userRepo, tokenService)

	recaptchaSecret := os.Getenv("RECAPTCHA_SECRET_KEY")
	if recaptchaSecret == "" {
		recaptchaSecret = "6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe" // Google test secret
	}
	recaptchaDisabled := os.Getenv("RECAPTCHA_DISABLED") == "true"
	recaptchaVerifier := recaptcha.NewVerifier(recaptchaSecret, recaptchaDisabled)

	authHandler := handlers.NewAuthHandler(authService, recaptchaVerifier)

	r := gin.Default()

	// Middleware
	r.Use(middleware.RateLimitMiddleware(60)) // 60 requests per minute
	r.Use(middleware.SecureHeadersMiddleware())

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	auth := r.Group("/auth")
	{
		auth.POST("/signup", authHandler.SignUp)
		auth.POST("/signin", authHandler.SignIn)
	}

	if err := r.Run(":8080"); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
