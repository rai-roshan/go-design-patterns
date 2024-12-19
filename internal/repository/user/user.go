package user

// user_inmemory_repo "rai_design_pattern/internal/repository/inmemory"

type UserRepo interface {
	GetUserById()
	GetSubscribedUsers()
	UnSubscribeUser()
	ActivateUser() // after OTP is updated
	CreateUser()
}
