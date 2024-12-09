package service

import (
	"context"
	"phase_example/dto"
)

type MailingService interface {
	// Send メールを送信する
	Send(ctx context.Context, info dto.EmailInfo) error
}
