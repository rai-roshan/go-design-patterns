package inmemory

import (
	user_domain "rai_design_pattern/domain/user"
	user_repo "rai_design_pattern/repository/user"
)

type userInMemo struct {
	idToUser map[uint32]user_domain.User	
}

func NewUserInMemoRepo() user_repo.UserRepo {
	return &userInMemo{
		idToUser:  make(map[uint32]user_domain.User),
	} 
}

func (uIm *userInMemo) GetUserById() {}
func (uIm *userInMemo) GetSubscribedUsers() {}
func (uIm *userInMemo) UnSubscribeUser() {}
func (uIm *userInMemo) ActivateUser() {}
func (uIm *userInMemo) CreateUser() {}