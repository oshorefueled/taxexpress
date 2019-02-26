package models

import (
	"fmt"
)

type Business struct {
	Id int `json:"id"`
	Name  string `json:"name"`
	RCNumber *int `json:"rc_number"`
	BusinessDescription string `json:"business_desc"`
	Email string `json:"email"`
	TotalRevenue float64 `json:"total_revenue"`
	TaxStatus *int `json:"tax_status"`
	CreatedAt *string `json:"created_at"`
	UpdateAt *string  `json:"updated_at"`
}

const businessTableName = "businesses"

func (b *Business) StoreBusiness () (int64, int64, error) {
	sqlQuery := fmt.Sprintf("INSERT %s SET name = ?, rc_number = ?, business_desc = ?, " +
		"email = ?, total_revenue = ?, tax_status = ?, created_at = NOW(), updated_at = NOW()", businessTableName)
	stmt, err := db.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	fmt.Println("email ", b.Email)
	res, err := stmt.Exec(b.Name, b.RCNumber, b.BusinessDescription, b.Email, b.TotalRevenue, b.TaxStatus)
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

func (b Business) GetAllBusinesses () ([]Business, error) {
	sqlQuery := "SELECT * FROM businesses"
	results, err := db.Query(sqlQuery)
	if err != nil {
		return []Business{}, err
	}
	var businesses []Business
	for results.Next() {
		err = results.Scan(&b.Id, &b.Name, &b.RCNumber, &b.BusinessDescription, &b.Email,
			&b.TotalRevenue, &b.TaxStatus, &b.CreatedAt, &b.UpdateAt)
		if err != nil {
			return []Business{}, err
		}
		fmt.Println(b.Name)
		businesses = append(businesses, b)
	}
	return businesses, err
}

func (b *Business) GetBusinessById () (err error) {
	sqlQuery := "SELECT * FROM businesses WHERE id=?"
	result := db.QueryRow(sqlQuery, b.Id)
	err = result.Scan(&b.Id, &b.Name, &b.RCNumber, &b.BusinessDescription, &b.Email,
		&b.TotalRevenue, &b.TaxStatus, &b.CreatedAt, &b.UpdateAt)
	return
}

func (b *Business) UpdateBusinessRevenue (revenue float64) (err error) {
	err = b.GetBusinessById()
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("Update businesses set total_revenue=?, updated_at=NOW() where id=?")
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	updatedRevenue := revenue + b.TotalRevenue
	_, err = stmt.Exec(updatedRevenue, b.Id)
	return
}



