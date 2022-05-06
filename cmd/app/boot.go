package app

import (
	"ddd-boilerplate/internal/app"
	"ddd-boilerplate/internal/cfg"
	"fmt"
)

func Bootstrap() {
	server := app.NewServer(&cfg.Cfg)
	server.Run(fmt.Sprintf(":%v", server.Cfg.App.ServerPort))
}
