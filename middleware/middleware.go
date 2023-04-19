package middleware

import (
	"main/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the header starts with "Bearer"
		if !strings.HasPrefix(authHeader, "Bearer") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get the token from the header
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

		// Verify token
		jwtData, err := helper.VerifyJWT(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Add JWT data to the context object
		c.Set(helper.JWTIssuer, jwtData[helper.JWTIssuer])
		c.Next()
	}
}
