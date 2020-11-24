package models

type Answer struct {
	AnswerId int64 `json:"questionId"`
	Answer string `json:"answer"`
	QuestionId int64 `json:"questionId"`
}