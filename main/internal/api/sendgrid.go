package api

import "github.com/stripe/stripe-go"

type SendGridRequest struct{}

func (cf *Config) SendConfirmationEmailRequest(event *stripe.Event) error {
	return nil
}
