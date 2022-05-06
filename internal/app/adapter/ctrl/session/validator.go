package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

// CreateSessionValidator is the validator to CreateSession
func CreateSessionValidator(c *gin.Context) {
	var in CreateSessionArgs
	c.BindJSON(&in)

	if err := validator.New().Struct(in); err != nil {
		c.String(http.StatusBadRequest, "invalid parameter")
		c.Abort()
		return
	}
	c.Next()
}
