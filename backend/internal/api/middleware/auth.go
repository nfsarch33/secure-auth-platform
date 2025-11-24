package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nfsarch33/secure-auth-platform/backend/internal/api"
	"github.com/nfsarch33/secure-auth-platform/backend/pkg/jwt"
)

const UserIDKey = "userID"

func AuthMiddleware(tokenService *jwt.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{Error: "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{Error: "Invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		claims, err := tokenService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{Error: "Invalid or expired token"})
			return
		}

		c.Set(UserIDKey, claims.UserID)
		c.Next()
	}
}
