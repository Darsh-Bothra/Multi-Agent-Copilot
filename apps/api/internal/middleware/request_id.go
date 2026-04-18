package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// it's a requestId middle ware
// A worker (function) that will run for every request

// Every request goes through the router, so we register the RequestID middleware once during setup. This middleware returns a handler function, and that returned function is what gets executed for every incoming request to generate and attach the request ID.

// It's here to create the worker
func RequestId() gin.HandlerFunc {

	// This what returning the actual worker, func is the worker
	return func(ctx *gin.Context) {

		// Generate uuid
		reqId := uuid.New().String()

		// Store the context
		ctx.Set("request_id", reqId)

		// Add to response handler
		ctx.Writer.Header().Set("X-Request-ID", reqId)

		// Move to next
		ctx.Next()
	}
}
