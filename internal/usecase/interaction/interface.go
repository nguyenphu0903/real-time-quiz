package interaction

import (
	"context"
)

type InteractionUsecase interface {
	JoinQuiz(ctx context.Context, quizID, userID string) error
	SubmitAnswer(ctx context.Context, quizID, userID, answer string, questionID string) (bool, int, int, error)
}
