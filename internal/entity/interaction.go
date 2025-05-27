package entity

type JoinRequest struct {
	QuizID string `json:"quiz_id"`
	UserID string `json:"user_id"`
}

type AnswerRequest struct {
	QuizID     string `json:"quiz_id"`
	UserID     string `json:"user_id"`
	Answer     string `json:"answer"`
	QuestionID string `json:"question_id"`
}

type ScoreUpdate struct {
	QuizID string `json:"quiz_id"`
	UserID string `json:"user_id"`
	Score  int    `json:"score"`
}
