package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"net/http"

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
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetUserService(),
		"status": true,
	})
}