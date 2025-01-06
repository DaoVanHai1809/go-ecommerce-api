package routers

import (
	"go-ecommerce-backend-api/internal/routers/admin"
	"go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)