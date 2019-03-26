package models

import (
	"database/sql"
	"fmt"
)

type Admin struct {
	Id string `json:"id"`
	Username  string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Token string `json:"token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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

func (a *Admin) GetAdminByToken () (err error) {
	var email, username string
	sqlQuery := `SELECT username, email FROM admins WHERE token=?`
	row := db.QueryRow(sqlQuery, a.Token)
	err = row.Scan(&username, &email)
	fmt.Println("username", a.Token)
	a.Username = username
	a.Email = email
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return
}

func (a Admin) GetAllAdminUsers () ([]interface{}, error) {
	sqlQuery := "SELECT * FROM admins"
	results, err := db.Query(sqlQuery)
	if err != nil {
		return []interface{}{}, err
	}
	var admins []interface{}
	for results.Next() {
		err = results.Scan(&a.Id, &a.Username, &a.Email, &a.Token, &a.Password, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return []interface{}{}, err
		}
		data := map[string]interface{}{
			"id": a.Id,
			"username": a.Username,
			"email": a.Email,
			"created_at": a.CreatedAt,
			"updated_at": a.UpdatedAt,
		}
		admins = append(admins, data)
	}
	return admins, err
}

