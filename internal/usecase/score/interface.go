package score

import (
	"context"
)

type LeaderboardEntry struct {
	User  string `json:"user"`
	Score int    `json:"score"`
	Rank  int    `json:"rank"`
}

type ScoreUsecase interface {
	GetLeaderboard(ctx context.Context, quizID string, topN int) ([]LeaderboardEntry, error)
	GetTotalScore(ctx context.Context, quizID, userID string) (int, error)
}
