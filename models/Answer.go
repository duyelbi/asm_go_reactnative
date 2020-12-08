package models

// Answer struct
type Answer struct {
	AnswerID int64 `json:"questionId"`
	Answer string `json:"answer"`
}