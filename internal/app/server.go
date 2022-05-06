package app

import (
	"ddd-boilerplate/internal/app/adapter/ctrl/index"
	"ddd-boilerplate/internal/app/adapter/ctrl/session"
	"ddd-boilerplate/internal/cfg"

	"github.com/gin-gonic/gin"
)

//Server is the server
type Server struct {
	*gin.Engine
	Cfg *cfg.Config
}

// NewServer creates a go-gin server
func NewServer(cfg *cfg.Config) *Server {
	server := &Server{
		Engine: gin.Default(),
		Cfg:    cfg,
	}
	setup(server)
	return server
}

func setup(server *Server) {
	indexGroup := server.Group("/")
	{
		healthzCtrl := index.Controller{}
		indexGroup.Any("/healthz", healthzCtrl.Healthz)
	}

	sessionGroup := server.Group("/session")
	{
		sessionCtrl := session.Controller{}
		sessionGroup.POST("/create", session.CreateSessionValidator, sessionCtrl.Create)
	}
}
