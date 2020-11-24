package dal

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitializeMySQL to MathTestDb
func InitializeMySQL() {
	dBConnection, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/MathTestDB")
	if err != nil {
		fmt.Println("Connetion Failed!!")
	}
	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	}

	db = dBConnection
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
}

//GetConnection is get MySQL Connection
func GetConnection() *sql.DB {
	if db == nil {
		InitializeMySQL()
	}
	return db
}

//CloseStmt after run stmt
func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}

//CloseRows when select
func CloseRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}