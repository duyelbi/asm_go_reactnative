package models

type Question struct {
	QuestionId int64 `json:"questionId"`
	Question string `json:"question"`
	Answer string `json:"answer"`
}