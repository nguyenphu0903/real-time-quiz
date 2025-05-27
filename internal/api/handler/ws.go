package handler

import (
	"context"
	"net/http"
	"real-time-quiz/internal/common/constant"
	"real-time-quiz/internal/usecase/score"
	"real-time-quiz/pkg"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler gin.HandlerFunc

func NewWSHandler(scoreUC score.ScoreUsecase) gin.HandlerFunc {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	return func(c *gin.Context) {
		quizID := c.Query("quiz_id")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		ctx := context.Background()
		pubsub := pkg.Rdb.Subscribe(ctx, constant.ScoreChannelPrefix+quizID)
		ch := pubsub.Channel()
		for range ch {
			entries, _ := scoreUC.GetLeaderboard(ctx, quizID, 10)
			conn.WriteJSON(gin.H{"leaderboard": entries})
		}
	}
}
