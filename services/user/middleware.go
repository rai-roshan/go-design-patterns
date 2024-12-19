package user

import (
	"context"
	"fmt"
	user_domain "rai_design_pattern/domain/user"
	"rai_design_pattern/logger"
	"time"
)

type loggingMiddleware struct {
	next UserService // main component of decorator design pattern
	log  logger.Logger
}

func NewLoggingMiddleware(log logger.Logger, next UserService) UserService {
	return &loggingMiddleware{
		next: next,
		log:  log,
	}
}

func (lm *loggingMiddleware) GetUserById(ctx context.Context, id uint32) (user user_domain.User, err error) {
	defer func(instance time.Time) {
		lm.log.Info(fmt.Sprintf("took: %d", time.Since(instance)))
	}(time.Now())
	return lm.next.GetUserById(ctx, id)
}

func (lm *loggingMiddleware) CreateUser(ctx context.Context, user user_domain.User) (err error) {
	defer func(instance time.Time) {
		lm.log.Info(fmt.Sprintf("took: %d", time.Since(instance)))
	}(time.Now())
	return lm.next.CreateUser(ctx, user)
}

func (lm *loggingMiddleware) SendOTPToUser(ctx context.Context, user user_domain.User) (err error) {
	defer func(instance time.Time) {
		lm.log.Info(fmt.Sprintf("took: %d", time.Since(instance)))
	}(time.Now())
	return lm.next.SendOTPToUser(ctx, user)
}

func (lm *loggingMiddleware) VerifyUserOTP(ctx context.Context) (err error) {
	defer func(instance time.Time) {
		lm.log.Info(fmt.Sprintf("took: %d", time.Since(instance)))
	}(time.Now())
	return lm.next.VerifyUserOTP(ctx)
}
