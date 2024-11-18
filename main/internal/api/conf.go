package api

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	StripeWebhookSecret     string
	SendGridApiToken        string
	SendGridEmailTemplateID string
}

func LoadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("[envconfig.Process]%w", err)
	}

	return &config, nil
}
