package api

type SendGridRequest struct{}

func (cf *Config) SendConfirmationEmailRequest(purchase *SessionCheckoutCompleted) error {
	return nil
}
