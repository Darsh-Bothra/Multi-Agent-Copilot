package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// it's a requestId middle ware


func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Generate uuid
		reqId := uuid.New().String();

		// Store the context
		ctx.Set("request_id", reqId);

		// Add to response handler
		ctx.Writer.Header().Set("X-Request-ID", reqId);

		// Move to next
		ctx.Next();
	}    
}