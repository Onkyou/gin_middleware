package auth

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const Key_Auth_Token_Header = "X-Auth-Token"

func AuthTokenMiddleware() gin.HandlerFunc {

	requiredToken := os.Getenv("API_TOKEN")

	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader(Key_Auth_Token_Header)

		if token == "" {
			respondWithError(c, 401, "API token required in HTTP headers.")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token found in HTTP headers.")
			return
		}

		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
