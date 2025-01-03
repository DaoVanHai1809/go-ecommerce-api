package service

import "go-ecommerce-backend-api/internal/repositories"

type UserService struct {
	userRepo *repositories.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepo(),
	}
}

func (us *UserService) GetUserService() string {
	return us.userRepo.GetUserRepo()
}