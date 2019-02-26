package models

import (
	"database/sql"
	"fmt"
)

type Admin struct {
	Username  string `json: "username"`
	Password string `json: "password"`
	Email string `json: "email"`
	Token string `json: "token"`
}

const adminTableName  = "admins"

func (a *Admin) CreateAdminUser () (int64, int64, error) {
	sqlQuery := fmt.Sprintf("INSERT %s SET username = ?, password = ?, email = ?, " +
		"token = ?, created_at = NOW(), updated_at = NOW()", adminTableName)
	stmt, err := db.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(a.Username, a.Password, a.Email, a.Token)
	if err != nil {
		return 0, 0, err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return affectedRows, lastInsertedId, err
}

func (a *Admin) GetAdminUser () (err error) {
	var password, email, token string
	sqlQuery := `SELECT password, email, token FROM admins WHERE username=?`
	row := db.QueryRow(sqlQuery, a.Username)
	err = row.Scan(&password, &email, &token)
	a.Password = password
	a.Email = email
	a.Token = token
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}
	return
}

