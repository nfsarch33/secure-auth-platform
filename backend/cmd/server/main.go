package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api/handlers"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api/middleware"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/repository/postgres"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/service"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/jwt"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Database Connection
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	// Initialize Dependencies
	userRepo := postgres.NewPostgresUserRepository(dbPool)
	
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}
	tokenService := jwt.NewTokenService(jwtSecret, "auth-service", 24*time.Hour)
	
	authService := service.NewAuthService(userRepo, tokenService)
	authHandler := handlers.NewAuthHandler(authService)

	// Router Setup
	r := gin.Default()
	
	// CORS Config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // In production, replace with specific origins
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Security Middlewares
	r.Use(middleware.SecureHeadersMiddleware())
	r.Use(middleware.RateLimitMiddleware(60)) // 60 requests per minute

	// Routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/signin", authHandler.SignIn)
	}

	// Server Setup
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Graceful Shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

