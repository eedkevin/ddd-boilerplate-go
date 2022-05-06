package main

import (
	"ddd-boilerplate/cmd/app"
	"ddd-boilerplate/internal/cfg"
)

func main() {
	cfg.Init()
	app.Bootstrap()
}
