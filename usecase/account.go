package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"phase_example/dto"
	"phase_example/dto/request"
	"phase_example/dto/response"
	"phase_example/repository"
	"phase_example/service"
)

type AccountUsecase interface {
	Signup(ctx context.Context, req request.PostAccountSignup) (*response.PostAccountSignup, error)
}

type accountUsecase struct {
	accountRepository repository.AccountRepository
	cacheRepository   repository.CacheRepository
	mailingService    service.MailingService
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

	if err := u.mailingService.Send(ctx, dto.EmailInfo{
		To:      []string{req.UserID},
		From:    "mTq4y@example.com",
		Subject: "認証メール",
		Body:    "認証キー: " + key,
		Type:    dto.TEXT,
	}); err != nil {
		return nil, err
	}

	user, err := u.accountRepository.FindByUserID(req.UserID)
	if err != nil {
		return nil, err
	}

	return response.NewPostAccountSignup(user.ID), nil
}

func NewAccountUsecase(
	accountRepository repository.AccountRepository,
	cacheRepository repository.CacheRepository,
	mailingService service.MailingService,
) AccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
		cacheRepository:   cacheRepository,
		mailingService:    mailingService,
	}
}
