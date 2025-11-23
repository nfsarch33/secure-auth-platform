package middleware

import "github.com/gin-gonic/gin"

func SecureHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Content Security Policy
		// Restrict sources for scripts, styles, etc.
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self'; frame-ancestors 'none'; object-src 'none'")

		// X-Frame-Options
		// Prevent clickjacking
		c.Header("X-Frame-Options", "DENY")

		// X-Content-Type-Options
		// Prevent MIME type sniffing
		c.Header("X-Content-Type-Options", "nosniff")

		// Strict-Transport-Security (HSTS)
		// Force HTTPS (max-age=1 year)
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		// Referrer-Policy
		// Control how much referrer info is sent
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// Permissions-Policy (Feature Policy)
		// Disable sensitive features
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		c.Next()
	}
}

