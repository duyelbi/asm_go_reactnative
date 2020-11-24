package dal

import "crud-mysql/models"

//InsertQuestion to MathTestDB
func InsertQuestion(question models.Question) (int64, int64, error) {
	GetConnection()
	sqlQuery := "INSERT Questions SET question=?, answer=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(question.Question, question.Answer)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertId, err := res.LastInsertId()
	return rowsAffected, lastInsertId, err
}

//UpdateQuestion in MathTestDB
func UpdateQuestion(question models.Question) (int64, error) {
	GetConnection()
	sqlQuery := "UPDATE Questions SET question=?, answer=? WHERE question_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(question.Question, question.Answer, question.QuestionId)
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
func DeleteQuestion(questionId int64) (int64, error) {
	GetConnection()
	sqlQuery := "DELETE FROM Questions WHERE question_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(questionId)
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
func GetQuestion(questionId int64) (models.Question, error) {
	GetConnection()
	sqlQuery := "SELECT question_id, question, answer FROM Questions WHERE question_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var question models.Question
	if err != nil {
		return question, err
	}
	res, err := stmt.Query(questionId)
	defer CloseRows(res)
	if err != nil {
		return question, err
	}
	if res.Next() {
		res.Scan(&question.QuestionId, &question.Question, &question.Answer)
	}
	return question, err
}

func GetAllQuestion() ([]models.Question, error) {
	GetConnection()
	sqlQuery := "SELECT * FROM Questions"
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
		res.Scan(&question.QuestionId, &question.Question, &question.Answer)
		listQuestion = append(listQuestion, question)
	}

	return listQuestion, err
}
