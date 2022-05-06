package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is a controller for health check
type Controller struct{}

// Healthz handles the healthz requests
func (ctrl Controller) Healthz(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
