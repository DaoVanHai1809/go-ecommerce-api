package initialize

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var router *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}
	// middlewares
	// router.Use() // logging - cross - limiter global
	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	MainGroup := router.Group("/v1")
	{
		MainGroup.GET("checkStatus") // tracking monitor
	}
	{
		adminRouter.InitAdminRouter(MainGroup)
		adminRouter.InitAdminUserRouter(MainGroup)
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}
	return router
}