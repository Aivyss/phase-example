package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"phase_example/dto/request"
	"phase_example/dto/response"
	"phase_example/repository"
)

type AccountUsecase interface {
	Signup(ctx context.Context, req request.PostAccountSignup) (*response.PostAccountSignup, error)
}

type accountUsecase struct {
	accountRepository repository.AccountRepository
	cacheRepository   repository.CacheRepository
}

func (u *accountUsecase) Signup(ctx context.Context, req request.PostAccountSignup) (*response.PostAccountSignup, error) {
	if err := u.accountRepository.Insert(req.UserID, req.Password); err != nil {
		return nil, err
	}

	if err := u.accountRepository.Insert(req.UserID, req.Password); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s_%s", req.UserID, uuid.New().String())
	if err := u.cacheRepository.Insert(key); err != nil {
		return nil, err
	}

	// dummy data
	return response.NewPostAccountSignup(1), nil
}

func NewAccountUsecase(
	accountRepository repository.AccountRepository,
	cacheRepository repository.CacheRepository,
) AccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
		cacheRepository:   cacheRepository,
	}
}
