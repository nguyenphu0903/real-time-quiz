package registry

import (
	"real-time-quiz/internal/api/handler"
	"real-time-quiz/internal/infrastructure/repository"
	"real-time-quiz/internal/usecase/interaction"
	"real-time-quiz/internal/usecase/score"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	InteractionHandler *handler.InteractionHandler
	ScoreHandler       *handler.ScoreHandler
	WSHandler          gin.HandlerFunc
}

func NewRegistry() *Registry {
	ir := repository.NewInteractionRepository()
	sr := repository.NewScoreRepository()
	interactionUC := interaction.NewInteractionUsecase(ir, sr)
	scoreUC := score.NewScoreUsecase(sr)
	wsHandler := handler.NewWSHandler(scoreUC)
	return &Registry{
		InteractionHandler: handler.NewInteractionHandler(interactionUC),
		ScoreHandler:       handler.NewScoreHandler(scoreUC),
		WSHandler:          wsHandler,
	}
}
