package repositories

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetUserRepo() string {
// 	return "user 01"
// }

// interface version
type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct{}

// GetUserByEmail implements IUserRepo.
func (ur *userRepo) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}
