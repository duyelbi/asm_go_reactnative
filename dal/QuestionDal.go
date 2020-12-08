package dal

import "crud-mysql/models"

//InsertQuestion to MathTestDB
func InsertQuestion(question models.Question) (int64, int64, error) {
	GetConnection()
	sqlQuery := "INSERT Questions SET difficulty=?, question=?, correct_answer=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(question.Difficulty, question.Question, question.CorrectAnswer)
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

//UpdateQuestion in MathTestDB
func UpdateQuestion(question models.Question) (int64, error) {
	GetConnection()
	sqlQuery := "UPDATE Questions SET difficulty=?, question=?, correct_answer=? WHERE question_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(question.Difficulty, question.Question, question.CorrectAnswer, question.QuestionID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

//DeleteQuestion in OrderDB
func DeleteQuestion(questionID int64) (int64, error) {
	GetConnection()
	sqlQuery := "DELETE FROM Questions WHERE question_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(questionID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

//GetQuestion from questionId
func GetQuestion(questionID int64) (models.Question, error) {
	GetConnection()
	sqlQuery := "SELECT question_id, difficulty, question, correct_answer FROM Questions WHERE question_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var question models.Question
	if err != nil {
		return question, err
	}
	res, err := stmt.Query(questionID)
	defer CloseRows(res)
	if err != nil {
		return question, err
	}
	if res.Next() {
		res.Scan(&question.QuestionID,  &question.Difficulty,  &question.Question, &question.CorrectAnswer)
	}
	return question, err
}

// GetAllQuestion select all dada from Questions table
func GetAllQuestion() ([]models.Question, error) {
	GetConnection()
	sqlQuery := "SELECT question_id, difficulty, question, correct_answer FROM Questions"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var listQuestion = make([]models.Question, 0)
	// var question models.Question
	if err != nil {
		return listQuestion, err
	}
	res, err := stmt.Query()
	defer CloseRows(res)
	if err != nil {
		return listQuestion, err
	}
	for res.Next() {
		var question models.Question
		res.Scan(&question.QuestionID, &question.Difficulty, &question.Question, &question.CorrectAnswer)
		listQuestion = append(listQuestion, question)
	}

	return listQuestion, err
}
