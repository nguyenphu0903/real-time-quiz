package interaction

import (
	"context"
	"real-time-quiz/internal/infrastructure/repository"
)

type interactionUsecase struct {
	interactionRepo repository.InteractionRepository
	scoreRepo       repository.ScoreRepository
}

func NewInteractionUsecase(ir repository.InteractionRepository, sr repository.ScoreRepository) InteractionUsecase {
	return &interactionUsecase{interactionRepo: ir, scoreRepo: sr}
}

func (u *interactionUsecase) JoinQuiz(ctx context.Context, quizID, userID string) error {
	return u.interactionRepo.AddParticipant(ctx, quizID, userID)
}

func (u *interactionUsecase) SubmitAnswer(
    ctx context.Context, quizID, userID, answer string, questionID string) (bool, int, int, error) {
	totalScore, _ := u.scoreRepo.GetTotalScore(ctx, quizID, userID)

	correct, score := u.interactionRepo.CheckAnswer(ctx, quizID, userID, answer, questionID)
	if correct {
		err := u.scoreRepo.UpdateScore(ctx, quizID, userID, score)
		u.scoreRepo.PublishScoreUpdate(ctx, quizID, userID, score)
		return true, score, totalScore, err
	}

	return false, 0, totalScore, nil
}
