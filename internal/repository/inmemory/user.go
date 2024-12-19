package inmemory

import (
	"errors"
	"fmt"
	user_domain "rai_design_pattern/internal/domain/user"
	user_repo "rai_design_pattern/internal/repository/user"
)

type userInMemo struct {
	idToUser map[uint32]user_domain.User	
}

func NewUserInMemoRepo() user_repo.UserRepo {
	return &userInMemo{
		idToUser:  make(map[uint32]user_domain.User),
	} 
}

func (uIm *userInMemo) GetUserById(id uint32) (user user_domain.User, err error) {
	if _, ok := uIm.idToUser[id]; !ok {
		err = errors.New(fmt.Sprintf("no user with id: %d exists", id))
		return
	}
 	return uIm.idToUser[id], nil
}
func (uIm *userInMemo) GetSubscribedUsers() {}
func (uIm *userInMemo) UnSubscribeUser() {}
func (uIm *userInMemo) ActivateUser() {}
func (uIm *userInMemo) CreateUser(user user_domain.User) (err error) {
	uIm.idToUser[ user.Id ] = user
	return nil
}