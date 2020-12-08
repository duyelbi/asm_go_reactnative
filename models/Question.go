package models

// Question struct
type Question struct {
	QuestionID int64 `json:"questionId"`
	Difficulty string `json:"difficulty"`
	Question string `json:"question"`
	CorrectAnswer string `json:"correctAnswer"`
}