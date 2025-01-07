package user

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/repositories"
	"go-ecommerce-backend-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// this is non-dependency
	urepo := repositories.NewUserRepo()
	us := service.NewUserService(urepo)
	uc := controller.NewUserController(us)
	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", uc.Register)
		userRouterPublic.POST("/login")
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
	}
}