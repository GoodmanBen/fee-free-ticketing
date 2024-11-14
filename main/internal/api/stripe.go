package api

import "github.com/gin-gonic/gin"

type SessionCheckoutCompleted struct{}

func (cf *Config) VerifyAndParseRequest(c *gin.Context) (*SessionCheckoutCompleted, error) {
	return nil, nil
}
