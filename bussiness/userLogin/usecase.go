package userLogin

import (
	"time"
)

type UserUsecase struct {
	Repo           Repository
	ContextTimeOut time.Duration
}

// func NewUserLoginUsecase(repo Repository, timeout time.Duration) Usecase {
// 	return &UserUsecase{
// 		Repo:           repo,
// 		ContextTimeOut: timeout,
// 	}
// }

// func (ulc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {

// }
