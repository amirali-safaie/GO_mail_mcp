package usecase

import (
	"github.com/amirali-safaie/GO_mail_mcp/internal/domain"
)

type SendEmailUseCase struct {
	repo domain.EmailRepository
}

func NewSendEmailUseCase(repo domain.EmailRepository) *SendEmailUseCase {
	return &SendEmailUseCase{repo: repo}
}

func (uc *SendEmailUseCase) Execute(to, subject, body string) error {
	//business logic befor the sending email
	return uc.repo.Send(domain.Email{
		To:      []string{to},
		Subject: subject,
		Body:    body,
	})
}
