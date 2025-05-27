package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, ih *InteractionHandler, sh *ScoreHandler) {
	r.POST("/api/join", ih.JoinQuiz)
	r.POST("/api/answer", ih.SubmitAnswer)
	r.GET("/api/leaderboard", sh.GetLeaderboard)
}
