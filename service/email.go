package service

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"phase_example/dto"
)

type MailingService interface {
	// Send メールを送信する
	Send(ctx context.Context, info dto.EmailInfo) error
}

type mailingService struct {
	client *sesv2.Client
}

// Send メールを送信する
func (s *mailingService) Send(ctx context.Context, info dto.EmailInfo) error {
	content := &types.Content{
		Data: &info.Body,
	}
	var body *types.Body
	switch info.Type {
	case dto.HTML:
		body = &types.Body{
			Html: content,
		}
	case dto.TEXT:
		body = &types.Body{
			Text: content,
		}
	}

	input := &sesv2.SendEmailInput{
		FromEmailAddress: &info.From,
		Destination: &types.Destination{
			ToAddresses:  info.To,
			BccAddresses: info.Bcc,
			CcAddresses:  info.Cc,
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: body,
				Subject: &types.Content{
					Data: &info.Subject,
				},
			},
		},
	}

	_, err := s.client.SendEmail(ctx, input)
	return err
}

// NewMailingService　service.MailingServiceを変換 (AWS SES-v2の実装)
func NewMailingService(client *sesv2.Client) MailingService {
	return &mailingService{
		client: client,
	}
}
