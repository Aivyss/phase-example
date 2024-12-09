package usecase

import (
	"context"
	"phase_example/dto/request"
	"phase_example/dto/response"
)

type AccountUsecase interface {
	Signup(ctx context.Context, req request.PostAccountSignup) (*response.PostAccountSignup, error)
}

type accountUsecase struct {
}

func (u *accountUsecase) Signup(ctx context.Context, req request.PostAccountSignup) (*response.PostAccountSignup, error) {
	// dummy data
	return response.NewPostAccountSignup(1), nil
}

func NewAccountUsecase() AccountUsecase {
	return &accountUsecase{}
}
