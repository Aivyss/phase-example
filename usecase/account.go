package usecase

import (
	"context"
	"phase_example/dto/request"
	"phase_example/dto/response"
	"phase_example/repository"
)

type AccountUsecase interface {
	Signup(ctx context.Context, req request.PostAccountSignup) (*response.PostAccountSignup, error)
}

type accountUsecase struct {
	accountRepository repository.AccountRepository
}

func (u *accountUsecase) Signup(ctx context.Context, req request.PostAccountSignup) (*response.PostAccountSignup, error) {
	if err := u.accountRepository.Insert(req.UserID, req.Password); err != nil {
		return nil, err
	}

	// dummy data
	return response.NewPostAccountSignup(1), nil
}

func NewAccountUsecase(accountRepository repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
	}
}
