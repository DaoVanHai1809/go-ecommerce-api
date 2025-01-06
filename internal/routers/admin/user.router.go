package admin

import "github.com/gin-gonic/gin"

type AdminUserRouter struct{}

func (aur *AdminRouter) InitAdminUserRouter(Router *gin.RouterGroup) {

	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active")
	}
}