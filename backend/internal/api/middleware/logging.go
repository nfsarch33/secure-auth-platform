package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LoggerMiddleware logs incoming requests with structured logging
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		method := c.Request.Method

		// Generate or get Request ID
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Header("X-Request-ID", requestID)

		c.Next()

		// Calculate duration
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		if raw != "" {
			path = path + "?" + raw
		}

		logger := slog.With(
			slog.String("request_id", requestID),
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status", statusCode),
			slog.Duration("latency", duration),
			slog.String("client_ip", clientIP),
		)

		switch {
		case statusCode >= 500:
			logger.Error("Request failed")
		case statusCode >= 400:
			logger.Warn("Request client error")
		default:
			logger.Info("Request processed")
		}
	}
}
