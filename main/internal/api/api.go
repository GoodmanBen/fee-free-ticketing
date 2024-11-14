package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiPrefix = "/api/v1"

func (cf *Config) NewAPI() *gin.Engine {
	r := gin.New()

	apiGroup := r.Group(apiPrefix)

	apiGroup.Handle(http.MethodPost, "/", cf.HandleStripeCheckoutSessionComplete)

	return r
}

// HandleStripeCheckoutSessionComplete handles a Stripe webhook event from session.checkout.completed
// format it, and then forward to the appropriate SendGrid email template.
func (cf *Config) HandleStripeCheckoutSessionComplete(c *gin.Context) {
	sessionComplete, err := cf.VerifyAndParseRequest(c)
	if err != nil {

		fmt.Printf("[VerifyAndParseRequest]%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error"})
	}

	err = cf.SendConfirmationEmailRequest(sessionComplete)
	if err != nil {

		fmt.Printf("[SendConfirmationEmailRequest]%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
