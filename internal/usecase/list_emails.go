package usecase

import (
	"github.com/amirali-safaie/GO_mail_mcp/internal/domain"
)

type ListEmailsUseCase struct {
	repo domain.EmailRepository
}

func NewListEmailsUseCase(repo domain.EmailRepository) *ListEmailsUseCase {
	return &ListEmailsUseCase{repo: repo}
}

func (uc *ListEmailsUseCase) Execute(n int) ([]domain.Email, error) {
	if n <= 0 {
		n = 10
	}
	return uc.repo.ListInbox(n)
}
