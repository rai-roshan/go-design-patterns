package mysql

import (
	user_domain "rai_design_pattern/internal/domain/user"
	user_repo "rai_design_pattern/internal/repository/user"
)

type userMysqlRepo struct {
	idToUser map[uint32]user_domain.User	
}

func NewUserMysqlRepo() user_repo.UserRepo {
	return &userMysqlRepo{
		idToUser:  make(map[uint32]user_domain.User),
	} 
}

func (uIm *userMysqlRepo) GetUserById() {}
func (uIm *userMysqlRepo) GetSubscribedUsers() {}
func (uIm *userMysqlRepo) UnSubscribeUser() {}
func (uIm *userMysqlRepo) ActivateUser() {}
func (uIm *userMysqlRepo) CreateUser() {}