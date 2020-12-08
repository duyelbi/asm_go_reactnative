package dal

import "crud-mysql/models"

//InsertAnswer to DB
func InsertAnswer(answer models.Answer) (int64, int64, error) {
	GetConnection()
	sqlQuery := "INSERT Answers SET answer_id=?, answer=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(answer.AnswerID, answer.Answer)
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

//UpdateAnswer in MathTestDB
func UpdateAnswer(answer models.Answer) (int64, error) {
	GetConnection()
	sqlQuery := "UPDATE Answers SET answer_id=?, answer=? WHERE answer_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(answer.AnswerID, answer.Answer)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

//DeleteAnswer in OrderDB
func DeleteAnswer(answerID int64) (int64, error) {
	GetConnection()
	sqlQuery := "DELETE FROM Answers WHERE answer_id = ?"
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

//GetAnswer from answerId
func GetAnswer(answerID int64) (models.Answer, error) {
	GetConnection()
	sqlQuery := "SELECT answer_id, answer FROM Answers WHERE answer_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var answer models.Answer
	if err != nil {
		return answer, err
	}
	res, err := stmt.Query(answerID)
	defer CloseRows(res)
	if err != nil {
		return answer, err
	}
	if res.Next() {
		res.Scan(&answer.AnswerID, &answer.Answer)
	}
	return answer, err
}

// GetAllAnswer select all dada from Answers table
func GetAllAnswer() ([]models.Answer, error) {
	GetConnection()
	sqlQuery := "SELECT answer_id, answer FROM Answers"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var listAnswer = make([]models.Answer, 0)
	// var answer models.Answer
	if err != nil {
		return listAnswer, err
	}
	res, err := stmt.Query()
	defer CloseRows(res)
	if err != nil {
		return listAnswer, err
	}
	for res.Next() {
		var answer models.Answer
		res.Scan(&answer.AnswerID, &answer.Answer)
		listAnswer = append(listAnswer, answer)
	}

	return listAnswer, err
}