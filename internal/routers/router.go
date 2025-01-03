package routers

import (
	"go-ecommerce-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	var r = gin.Default()
	var v1 = r.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user", controller.NewUserController().GetUser)
	}
	return r
}