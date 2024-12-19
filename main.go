package main

import (
	"fmt"
	"rai_design_pattern/logger"
	"rai_design_pattern/services/finance"
	"rai_design_pattern/services/user"
	user_repo "rai_design_pattern/repository/user"
	user_domain "rai_design_pattern/domain/user"
)

func main() {
	fmt.Println("-> main function invoked")

	logInst := logger.GetLogger()

	// TODO : initalize user repo using [Factory pattern]
	var userRepo user_repo.UserRepo

	var financeService finance.FinanceService
	{
		financeService = finance.NewFinanceService()
		financeService = finance.NewLoggingMiddleware(logInst, financeService) // [Decorator Pattern]
		// TODO : add caching wrapper
	}

	var userService user.UserService
	{
		userService = user.NewUserService(userRepo)
		userService = user.NewLoggingMiddleware(logInst, userService) // [Decorator Pattern]
		// TODO : add caching wrapper
	}

	fmt.Printf("add : %d\n", financeService.Add(int(1), int(2)))
	fmt.Printf("sub : %d\n", financeService.Sub(int(1), int(2)))

	userService.CreateUser(nil, user_domain.User{})
	userService.GetUserById(nil, uint32(1))

}
