package repository

import (
	"context"
	"real-time-quiz/internal/common"
	"real-time-quiz/internal/common/constant"
	"real-time-quiz/pkg"
)

type InteractionRepository interface {
	AddParticipant(ctx context.Context, quizID, userID string) error
	CheckAnswer(ctx context.Context, quizID, userID, answer string, questionID string) (bool, int)
}

type interactionRepo struct{}

func NewInteractionRepository() InteractionRepository {
	return &interactionRepo{}
}

func (r *interactionRepo) AddParticipant(ctx context.Context, quizID, userID string) error {
	key := constant.ParticipantsKeyPrefix + quizID
	return pkg.Rdb.SAdd(ctx, key, userID).Err()
}

func (r *interactionRepo) CheckAnswer(
	ctx context.Context, quizID, userID, answer string, questionID string,
) (bool, int) {
	if userID == "" || quizID == "" {
		return false, 0
	}

	key := constant.ParticipantsKeyPrefix + quizID
	isMember, err := pkg.Rdb.SIsMember(ctx, key, userID).Result()
	if err != nil || !isMember {
		return false, 0
	}

	for _, q := range common.DemoQuestions {
		if q.QuizID == quizID && q.ID == questionID {
			if answer == q.Answer {
				return true, 10
			}
			return false, 0
		}
	}
	return false, 0
}
