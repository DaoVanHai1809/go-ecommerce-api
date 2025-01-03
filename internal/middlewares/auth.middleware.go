package middlewares

import (
	"fmt"
	"go-ecommerce-backend-api/packages/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token = ctx.GetHeader("Authorization")
		fmt.Println(token)
		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrCodeInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
		fmt.Println("okokokk")
	}
}