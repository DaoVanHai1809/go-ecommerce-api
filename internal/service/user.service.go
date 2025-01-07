package service

import (
	"go-ecommerce-backend-api/internal/repositories"
	"go-ecommerce-backend-api/packages/response"
)

// type UserService struct {
// 	userRepo *repositories.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repositories.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetUserService() string {
// 	return us.userRepo.GetUserRepo()
// }

// interface version
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repositories.IUserRepo
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repositories.IUserRepo) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
