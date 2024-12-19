package main

import (
	"context"
	"fmt"
	user_domain "rai_design_pattern/internal/domain/user"
	respo_factory "rai_design_pattern/internal/repository/factory"
	user_repo "rai_design_pattern/internal/repository/user"
	"time"

	// "rai_design_pattern/internal/services/finance"
	"rai_design_pattern/internal/services/user"
	"rai_design_pattern/pkg/logger"
)

func main() {
	fmt.Println("-> main function invoked")

	logInst := logger.GetLogger()

	// TODO : initalize user repo using [Factory pattern]
	var userRepo user_repo.UserRepo
	{
		userRepo = respo_factory.GetUserRepository("inmemory")
	}

	// var financeService finance.FinanceService
	// {
	// 	financeService = finance.NewFinanceService()
	// 	financeService = finance.NewLoggingMiddleware(logInst, financeService) // [Decorator Pattern]
	// 	// TODO : add caching wrapper
	// }

	var userService user.UserService
	{
		userService = user.NewUserService(userRepo)
		userService = user.NewLoggingMiddleware(logInst, userService) // [Decorator Pattern]
		// TODO : add caching wrapper
	}

	// fmt.Printf("add : %d\n", financeService.Add(int(1), int(2)))
	// fmt.Printf("sub : %d\n", financeService.Sub(int(1), int(2)))

	err := userService.CreateUser(context.TODO(), user_domain.User{
		Id:        1,
		Name:      "Roshan",
		Age:       uint32(24),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		fmt.Printf("err : %+n\n", err)
	}
	_, _ = userService.GetUserById(context.TODO(), uint32(1))
	_, _ = userService.GetUserById(context.TODO(), uint32(2))
}
