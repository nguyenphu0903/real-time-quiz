package repository

import (
	"context"
	"fmt"
	"real-time-quiz/internal/common/constant"
	"real-time-quiz/pkg"

	"github.com/go-redis/redis/v8"
)

type ScoreRepository interface {
	UpdateScore(ctx context.Context, quizID, userID string, score int) error
	PublishScoreUpdate(ctx context.Context, quizID, userID string, score int) error
	GetLeaderboard(ctx context.Context, quizID string, topN int) ([]string, []int, error)
	GetTotalScore(ctx context.Context, quizID, userID string) (int, error)
}

type scoreRepo struct{}

func NewScoreRepository() ScoreRepository {
	return &scoreRepo{}
}

func (r *scoreRepo) UpdateScore(ctx context.Context, quizID, userID string, score int) error {
	key := constant.LeaderboardKeyPrefix + quizID
    return pkg.Rdb.ZIncrBy(ctx, key, float64(score), userID).Err()
}

func (r *scoreRepo) PublishScoreUpdate(ctx context.Context, quizID, userID string, score int) error {
	channel := constant.ScoreChannelPrefix + quizID
	msg := fmt.Sprintf("%s:%d", userID, score)
	return pkg.Rdb.Publish(ctx, channel, msg).Err()
}

func (r *scoreRepo) GetLeaderboard(ctx context.Context, quizID string, topN int) ([]string, []int, error) {
	key := constant.LeaderboardKeyPrefix + quizID
	res, err := pkg.Rdb.ZRevRangeWithScores(ctx, key, 0, int64(topN-1)).Result()
	if err != nil {
		return nil, nil, err
	}
	users := []string{}
	scores := []int{}
	for _, z := range res {
		users = append(users, z.Member.(string))
		scores = append(scores, int(z.Score))
	}
	return users, scores, nil
}

func (r *scoreRepo) GetTotalScore(ctx context.Context, quizID, userID string) (int, error) {
    key := constant.LeaderboardKeyPrefix + quizID
    score, err := pkg.Rdb.ZScore(ctx, key, userID).Result()
    if err != nil {
        if err == redis.Nil {
            return 0, nil
        }
        return 0, err
    }
    return int(score), nil
}