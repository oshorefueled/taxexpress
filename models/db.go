package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

const (
	databaseUser string = "root"
	databaseName string = "taxexpress"
	password string = "bubble28"
)

func InitializeDB () {
	fmt.Println("Connecting to database")
	dbConnection, err := sql.Open("mysql", databaseUser+":" +password+"@(localhost:3306)/"+databaseName)
	if err != nil {
		fmt.Printf("Connection to %s database failed \n", databaseName)
	}
	err = dbConnection.Ping()
	if err == nil {
		fmt.Printf("Connected to %s successfully \n", databaseName)
	} else {
		fmt.Println("Failed to ping database")
	}
	db = dbConnection
}

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}