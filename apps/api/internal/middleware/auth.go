package middleware

import (
	"api/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)



func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Get the header
		header := ctx.GetHeader("Authorization");

		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			ctx.Abort()
			return
		}

		// Header Format : Bearer <Token>
		parts := strings.Split(header, "")

		if len(parts) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		userId, err := auth.ValidateToken(parts[1])

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}
		
		// Store the user context	
		ctx.Set("user_id", userId);

		ctx.Next();
	}
}