package handler

import (
	"net/http"
	"real-time-quiz/internal/usecase/score"

	"github.com/gin-gonic/gin"
)

type ScoreHandler struct {
	uc score.ScoreUsecase
}

func NewScoreHandler(uc score.ScoreUsecase) *ScoreHandler {
	return &ScoreHandler{uc: uc}
}

func (h *ScoreHandler) GetLeaderboard(c *gin.Context) {
	quizID := c.Query("quiz_id")
	entries, _ := h.uc.GetLeaderboard(c.Request.Context(), quizID, 10)
	c.JSON(http.StatusOK, gin.H{"leaderboard": entries})
}
