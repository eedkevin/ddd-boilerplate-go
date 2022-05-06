package session

import (
	"ddd-boilerplate/internal/app/application/service"
	"ddd-boilerplate/internal/app/application/usecase"
	"ddd-boilerplate/internal/app/domain/repository"
	"ddd-boilerplate/internal/app/domain/vo"

	"github.com/gin-gonic/gin"
)

// Controller is a controller for session operations
type Controller struct {
	Telecom     service.ITelecom
	ServicePool repository.IServicePool
}

// Create handles the create session requests
func (ctrl Controller) Create(c *gin.Context) {
	result, _ := usecase.CreateSession(ctrl.Telecom, ctrl.ServicePool, vo.Paticipant{})
	c.JSON(200, result)
}
