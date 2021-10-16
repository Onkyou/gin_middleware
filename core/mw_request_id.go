package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, _ := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}
