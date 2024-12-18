package main

import (
	"fmt"
	"rai_design_pattern/logger"
	"rai_design_pattern/services/finance"
)

func main() {
	fmt.Println("-> main function invoked")

	logInst := logger.GetLogger()

	var financeService finance.FinanceService
	{
		financeService = finance.NewFinanceService()
		financeService = finance.NewLoggingMiddleware(logInst, financeService) // wrap logger
	}

	fmt.Printf("add : %d\n", financeService.Add(int(1), int(2)))
	fmt.Printf("sub : %d\n", financeService.Sub(int(1), int(2)))

}
