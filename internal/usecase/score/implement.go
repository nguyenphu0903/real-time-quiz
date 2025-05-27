package score

import (
	"context"
	"real-time-quiz/internal/infrastructure/repository"
)

type scoreUsecase struct {
	scoreRepo repository.ScoreRepository
}

func NewScoreUsecase(sr repository.ScoreRepository) ScoreUsecase {
	return &scoreUsecase{scoreRepo: sr}
}

func (u *scoreUsecase) GetLeaderboard(ctx context.Context, quizID string, topN int) ([]LeaderboardEntry, error) {
	users, scores, err := u.scoreRepo.GetLeaderboard(ctx, quizID, topN)
	if err != nil {
		return nil, err
	}
	entries := make([]LeaderboardEntry, 0, len(users))
	for i := range users {
		entries = append(entries, LeaderboardEntry{User: users[i], Score: scores[i], Rank: i + 1})
	}
	return entries, nil
}

func (u *scoreUsecase) GetTotalScore(ctx context.Context, quizID, userID string) (int, error) {
	return u.scoreRepo.GetTotalScore(ctx, quizID, userID)
}
