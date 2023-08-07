package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

// GinContextKey is the key to get gin context from context
type GinContextKey string

var ginContextKey = GinContextKey("ginContext")

// GinContextToContextMiddleware is a middleware to add gin context in context
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
