package api

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"
)

type SessionCheckoutCompleted struct{}

func (cf *Config) VerifyAndParseRequest(c *gin.Context) (*stripe.Event, error) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("[io.ReadAll]%v", err)

		return nil, err
	}

	event, err := webhook.ConstructEvent(
		bodyBytes,
		c.Request.Header.Get("Stripe-Signature"),
		cf.StripeWebhookSecret,
	)
	if err != nil {
		return nil, fmt.Errorf("[webhook.ConstructEvent]%w", err)
	}

	if event.Type != "checkout.session.completed" {
		fmt.Printf("received valid event but an unsupported type: %v\n", event.Type)

		return nil, nil
	}

	return &event, nil
}
