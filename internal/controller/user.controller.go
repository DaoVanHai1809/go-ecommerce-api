package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/packages/response"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	// response.SuccessResponse(c, 20001, uc.userService.GetUserService())
	response.ErrorResponse(c, 20003, "No need")
}