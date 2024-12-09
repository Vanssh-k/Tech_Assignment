package main

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func extractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func handleJWTError(c *gin.Context, err error) {
	switch err.Error() {
	case "token is expired":
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
	case "token is revoked":
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been revoked"})
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}