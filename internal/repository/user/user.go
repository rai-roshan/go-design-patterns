package user

import (
	user_domain "rai_design_pattern/internal/domain/user"
)

type UserRepo interface {
	GetUserById(id uint32) (user_domain.User, error)
	GetSubscribedUsers()
	UnSubscribeUser()
	ActivateUser() // after OTP is updated
	CreateUser(user_domain.User) (error)
}
