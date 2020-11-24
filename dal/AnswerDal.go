package dal

import "crud-mysql/models"

//InsertAnswer to MathTestDB
func InsertAnswer(answer models.Answer) (int64, int64, error) {
	GetConnection()
	sqlQuery := "INSERT Answers SET answer=?, question_id=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(answer.Answer, answer.QuestionId)
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

//UpdateAnswer in MathTestDB
func UpdateAnswer(answer models.Answer) (int64, error) {
	GetConnection()
	sqlQuery := "UPDATE Answers SET answer=?, question_id=? WHERE answer_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(answer.Answer, answer.QuestionId, answer.AnswerId)
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
func DeleteAnswer(answerId int64) (int64, error) {
	GetConnection()
	sqlQuery := "DELETE FROM Answers WHERE answer_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(answerId)
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
func GetAnswer(answerId int64) (models.Answer, error) {
	GetConnection()
	sqlQuery := "SELECT answer_id, answer, question_id FROM Answers WHERE answer_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var answer models.Answer
	if err != nil {
		return answer, err
	}
	res, err := stmt.Query(answerId)
	defer CloseRows(res)
	if err != nil {
		return answer, err
	}
	if res.Next() {
		res.Scan(&answer.AnswerId, &answer.Answer, &answer.QuestionId)
	}
	return answer, err
}
