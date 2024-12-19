package factory

import (
	user_repo "rai_design_pattern/internal/repository/user"
	user_mysql_repo "rai_design_pattern/internal/repository/mysql"
	user_inmemory_repo "rai_design_pattern/internal/repository/inmemory"
)

func GetUserRepositoryFactory(repoMode string) user_repo.UserRepo {
	userMysqlRepo := user_mysql_repo.NewUserMysqlRepo()
	userInMemoryRepo := user_inmemory_repo.NewUserInMemoRepo()

	switch repoMode {
		case "mysql":
			return userMysqlRepo
		default:
			return userInMemoryRepo
	}
}