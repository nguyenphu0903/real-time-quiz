package handler

import (
	"net/http"
	"real-time-quiz/internal/entity"
	"real-time-quiz/internal/usecase/interaction"

	"github.com/gin-gonic/gin"
)

type InteractionHandler struct {
	uc interaction.InteractionUsecase
}

func NewInteractionHandler(uc interaction.InteractionUsecase) *InteractionHandler {
	return &InteractionHandler{uc: uc}
}

func (h *InteractionHandler) JoinQuiz(c *gin.Context) {
	var req entity.JoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.uc.JoinQuiz(c.Request.Context(), req.QuizID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "joined"})
}

func (h *InteractionHandler) SubmitAnswer(c *gin.Context) {
	var req entity.AnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	correct, score, totalScore, _ := h.uc.SubmitAnswer(c.Request.Context(), req.QuizID, req.UserID, req.Answer, req.QuestionID)
	c.JSON(http.StatusOK, gin.H{"correct": correct, "score": score, "total_score": totalScore})
}
