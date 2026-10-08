package domain

type EmailRepository interface {
	ListInbox(limit int) ([]Email, error)
	Send(email Email) error
}
