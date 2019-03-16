package models

import "fmt"

type Message struct {
	Id int `json:"id"`
	Type  string `json:"type"`
	Message  string `json:"message"`
	CreatedAt *string `json:"created_at"`
	UpdateAt *string  `json:"updated_at"`
}

const messageTableName = "messages"

func (m *Message) CreateMessage () (int64, int64, error) {
	sqlQuery := fmt.Sprintf("INSERT %s SET type = ?, message = ?, created_at = NOW(), " +
		"updated_at = NOW()", messageTableName)
	stmt, err := db.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(m.Type, m.Message)
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