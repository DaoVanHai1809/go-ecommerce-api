package initialize

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var router = gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	// Use the middleware
	router.Use(middlewares.AuthenMiddleware())
	var v1 = router.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user", controller.NewUserController().GetUser)
	}
	return router
}