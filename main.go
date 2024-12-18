package main

import (
	"fmt"
	"rai_design_pattern/logger"
	"time"
)

type FinanceService interface {
	Add(a int, b int) (c int)
	Sub(a int, b int) (c int)
}

type financeService struct {

}

func (fs *financeService) Add(a int, b int) (c int) {
	return a + b
}

func (fs *financeService) Sub(a int, b int) (c int) {
	return a - b
}

func NewFinanceService() FinanceService {
	return &financeService{}
}


type loggingMiddleware struct {
	next FinanceService // main component of decorator design pattern
	log logger.Logger
}

func (lm *loggingMiddleware) Add(a int, b int) (c int) {
	defer func(instance time.Time) {
		lm.log.Info(fmt.Sprintf("took: %d || add: %d + %d = %d",  time.Since(instance), a,b,c))
	}(time.Now())
	return lm.next.Add(a, b) // closure behaviour
}

func (lm *loggingMiddleware) Sub(a int, b int) (c int) {
	defer func(instance time.Time) {
		lm.log.Info(fmt.Sprintf("took: %d || sub: %d - %d = %d",  time.Since(instance).Milliseconds(), a,b,c))
	}(time.Now())
	return lm.next.Sub(a, b)
}

func NewLoggingMiddleware( log logger.Logger, next FinanceService ) ( FinanceService ) {
	return &loggingMiddleware{
		next: next,
		log: log,
	} 
}

func main() {
	fmt.Println("-> main function invoked")
	
	logInst := logger.GetLogger()

	finServ := NewFinanceService()
	finLogger := NewLoggingMiddleware( logInst, finServ) // wrap logger

	fmt.Printf("add : %d\n", finLogger.Add(int(1), int(2)) )
	fmt.Printf("sub : %d\n", finLogger.Sub(int(1), int(2)) )

}