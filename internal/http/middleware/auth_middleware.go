package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/kika1s1/Go-Loan-Tracker-API/internal/usecase"
	"github.com/kika1s1/Go-Loan-Tracker-API/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			respondUnauthorized(c, "Authorization header required")
			return
		}
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			respondUnauthorized(c, "Invalid authorization header")
			return
		}

		tokenString := authParts[1]
		if tokenString == "" {
			respondUnauthorized(c, "Bearer token required")
			return
		}

		// Validate the token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			respondUnauthorized(c, "Invalid token")
			return
		}
	

		// Set the claims in the context and proceed
		c.Set("claims", claims)
		c.Next()
	}
}

func respondUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}
