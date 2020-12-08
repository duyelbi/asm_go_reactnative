package dal

import "crud-mysql/models"

//InsertQuestionAnswer to MathTestDB
func InsertQuestionAnswer(questionAnswer models.QuestionAnswer) (int64, int64, error) {
	GetConnection()
	sqlQuery := "INSERT Question_Answer SET question_id=?, answer_id=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(questionAnswer.QuestionID, questionAnswer.AnswerID)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertID, err := res.LastInsertId()
	return rowsAffected, lastInsertID, err
}

//UpdateQuestionAnswer in MathTestDB
func UpdateQuestionAnswer(questionAnswer models.QuestionAnswer) (int64, error) {
	GetConnection()
	sqlQuery := "UPDATE Question_Answer SET question_id=?, answer_id=? WHERE answer_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(questionAnswer.QuestionID, questionAnswer.AnswerID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

//DeleteQuestionAnswer in QuestionAnswer
func DeleteQuestionAnswer(answerID int64) (int64, error) {
	GetConnection()
	sqlQuery := "DELETE FROM Question_Answer WHERE answer_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(answerID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

//GetQuestionAnswer from answerId
func GetQuestionAnswer(answerID int64) (models.QuestionAnswer, error) {
	GetConnection()
	sqlQuery := "SELECT question_id,  answer_id FROM Question_Answer WHERE answer_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var questionAnswer models.QuestionAnswer
	if err != nil {
		return questionAnswer, err
	}
	res, err := stmt.Query(answerID)
	defer CloseRows(res)
	if err != nil {
		return questionAnswer, err
	}
	if res.Next() {
		res.Scan(&questionAnswer.QuestionID,  &questionAnswer.AnswerID)
	}
	return questionAnswer, err
}

// GetAllQuestionAnswer select all dada from QuestionAnswer table
func GetAllQuestionAnswer() ([]models.QuestionAnswer, error) {
	GetConnection()
	sqlQuery := "SELECT q.*, a.* FROM Questions q INNER JOIN Question_Answer qa ON qa.question_id = q.question_id INNER JOIN Answers a ON a.answer_id = qa.answer_id"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var listQuestionAnswer = make([]models.QuestionAnswer, 0)
	// var question models.QuestionAnswer
	if err != nil {
		return listQuestionAnswer, err
	}
	res, err := stmt.Query()
	defer CloseRows(res)
	if err != nil {
		return listQuestionAnswer, err
	}
	for res.Next() {
		var questionAnswer models.QuestionAnswer
		res.Scan(&questionAnswer.QuestionID, &questionAnswer.Difficulty, &questionAnswer.Question, &questionAnswer.CorrectAnswer,  &questionAnswer.AnswerID, &questionAnswer.Answer)
		listQuestionAnswer = append(listQuestionAnswer, questionAnswer)
	}

	return listQuestionAnswer, err
}


// GetQuestionAnswerID select all dada from QuestionAnswer table
func GetQuestionAnswerID(questionID int64) ([]models.QuestionAnswer, error) {
	GetConnection()
	sqlQuery := "SELECT q.*, a.* FROM Questions q INNER JOIN Question_Answer qa ON qa.question_id = q.question_id INNER JOIN Answers a ON a.answer_id = qa.answer_id WHERE qa.question_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var listQuestionAnswer = make([]models.QuestionAnswer, 0)
	// var question models.QuestionAnswer
	if err != nil {
		return listQuestionAnswer, err
	}
	res, err := stmt.Query(questionID)
	defer CloseRows(res)
	if err != nil {
		return listQuestionAnswer, err
	}
	for res.Next() {
		var questionAnswer models.QuestionAnswer
		res.Scan(&questionAnswer.QuestionID, &questionAnswer.Difficulty, &questionAnswer.Question, &questionAnswer.CorrectAnswer,  &questionAnswer.AnswerID, &questionAnswer.Answer)
		listQuestionAnswer = append(listQuestionAnswer, questionAnswer)
	}

	return listQuestionAnswer, err
}