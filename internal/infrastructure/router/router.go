package router

import (
	"real-time-quiz/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(ih *handler.InteractionHandler, sh *handler.ScoreHandler, wsHandler gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	handler.RegisterRoutes(r, ih, sh)
	r.GET("/ws/leaderboard", wsHandler)
	return r
}
