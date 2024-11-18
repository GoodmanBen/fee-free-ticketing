package api

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stripe/stripe-go"
)

const TicketCost = 20.0

type StripeWebhookEvent struct {
	AmountTotal     float64 `json:"amount_total"`
	Currency        string  `json:"currency"`
	CustomerDetails `json:"customer_details"`
}

type CustomerDetails struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (cf *Config) SendConfirmationEmailRequest(event *stripe.Event) error {
	webhookBody := StripeWebhookEvent{}
	err := json.Unmarshal(event.Data.Raw, &webhookBody)
	if err != nil {
		return fmt.Errorf("failed to unmarshal webhook json: %w", err)
	}

	m := mail.NewV3Mail()

	address := "hello@dragondrop.cloud"
	name := "Ben Goodman"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	m.SetTemplateID(cf.SendGridEmailTemplateID)

	p := mail.NewPersonalization()

	tos := []*mail.Email{
		mail.NewEmail(webhookBody.Name, "goodmanben@dragondrop.cloud"),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("first_name", webhookBody.Name)
	p.SetDynamicTemplateData("ticket_count", math.Floor(webhookBody.AmountTotal/TicketCost))
	p.SetDynamicTemplateData("total_cost", webhookBody.AmountTotal)

	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(
		cf.SendGridApiToken,
		"/v3/mail/send",
		"https://api.sendgrid.com",
	)

	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body

	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Printf("[sendgrid.API]%v", err)
		return fmt.Errorf("[sendgrid.API]%w", err)
	}

	fmt.Printf("Status:\n%v\n\nBody:\n%s\n", response.StatusCode, response.Body)

	return nil
}
