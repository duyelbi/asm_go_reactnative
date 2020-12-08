package models

// QuestionAnswer struct
type QuestionAnswer struct {
	QuestionID int64 `json:"questionId"`
	Difficulty string `json:"difficulty"`
	Question string `json:"question"`
	CorrectAnswer string `json:"correctAnswer"`
	AnswerID int64 `json:"answerId"`
	Answer string `json:"answer"`
}