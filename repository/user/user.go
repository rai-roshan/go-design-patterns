package user

type UserRepo interface {
	GetUserById()
	GetSubscribedUsers()
	UnSubscribeUser()
	ActivateUser() // after OTP is updated
	CreateUser()
}