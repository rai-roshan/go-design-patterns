package finance

import (
	"fmt"
	"time"
	"rai_design_pattern/pkg/logger"
)

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