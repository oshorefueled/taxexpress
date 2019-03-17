package models

import "fmt"

type Message struct {
	Id int `json:"id"`
	Message  string `json:"message"`
	Type  string `json:"type"`
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

func (m Message) GetAllMessages () ([]Message, error) {
	sqlQuery := "SELECT * FROM messages"
	results, err := db.Query(sqlQuery)
	if err != nil {
		return []Message{}, err
	}
	var messages []Message
	for results.Next() {
		err = results.Scan(&m.Id, &m.Message, &m.Type, &m.CreatedAt, &m.UpdateAt)
		if err != nil {
			return []Message{}, err
		}
		messages = append(messages, m)
	}
	return messages, err
}