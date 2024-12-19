package inmemory

import (
	user_domain "rai_design_pattern/domain/user"
	user_repo "rai_design_pattern/repository/user"
)

type userMysqlRepo struct {
	idToUser map[uint32]user_domain.User	
}

func NewUserInMemoRepo() user_repo.UserRepo {
	return &userMysqlRepo{
		idToUser:  make(map[uint32]user_domain.User),
	} 
}

func (uIm *userMysqlRepo) GetUserById() {}
func (uIm *userMysqlRepo) GetSubscribedUsers() {}
func (uIm *userMysqlRepo) UnSubscribeUser() {}
func (uIm *userMysqlRepo) ActivateUser() {}
func (uIm *userMysqlRepo) CreateUser() {}