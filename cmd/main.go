package main

import (
	"real-time-quiz/config"
	"real-time-quiz/internal/infrastructure/router"
	"real-time-quiz/internal/registry"
	"real-time-quiz/pkg"
)

func main() {
	cfg := config.Load()
	pkg.InitRedis()
	reg := registry.NewRegistry()
	r := router.NewRouter(reg.InteractionHandler, reg.ScoreHandler, reg.WSHandler)
	r.Run(":" + cfg.Port)
}
