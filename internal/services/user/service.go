package user

import (
	"context"
	user_domain "rai_design_pattern/internal/domain/user"
	user_repo "rai_design_pattern/internal/repository/user"
)

type UserService interface {
	GetUserById(ctx context.Context, id uint32) (user user_domain.User, err error)
	CreateUser(ctx context.Context, user user_domain.User) (err error)
	SendOTPToUser(ctx context.Context, user user_domain.User) (err error)
	VerifyUserOTP(ctx context.Context) (err error)
}

type userService struct {
	userRepo user_repo.UserRepo
}

func NewUserService(userRepo user_repo.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) GetUserById(ctx context.Context, id uint32) (user user_domain.User, err error) {
	return
}

func (u *userService) CreateUser(ctx context.Context, user user_domain.User) (err error) {
	return
}

func (u *userService) SendOTPToUser(ctx context.Context, user user_domain.User) (err error) {
	return
}

func (u *userService) VerifyUserOTP(ctx context.Context) (err error) {
	return
}