package entity

type Question struct {
	ID      string   `json:"id"`
	QuizID  string   `json:"quiz_id"`
	Content string   `json:"content"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}

type QuestionBank struct {
	Questions []Question `json:"questions"`
}
