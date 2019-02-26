package models

import (
	"fmt"
)

type Tax struct {
	Id int `json:"id"`
	BusinessId int `json:"business_id"`
	TaxPeriod string `json:"tax_period"`
	Revenue float64 `json:"revenue"`
	TaxPaid *float32 `json:"tax_paid"`
	DatePaid *string `json:"date_paid"`
	CreatedAt *string `json:"created_at"`
	UpdateAt *string  `json:"updated_at"`
}

const taxTableName  = "tax"

func (t *Tax) SaveBusinessRevenue () error {
	sqlQuery := fmt.Sprintf("INSERT %s SET business_id = ?, tax_period = ?, revenue = ?, " +
		"created_at = NOW(), updated_at = NOW()", taxTableName)
	stmt, err := db.Prepare(sqlQuery)
	fmt.Println(err)
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(t.BusinessId, t.TaxPeriod, t.Revenue)
	if err != nil {
		return err
	}
	b := Business{}
	b.Id = t.BusinessId
	err = b.UpdateBusinessRevenue(t.Revenue)
	return err
}

func (t Tax) GetAllPaidTaxes () ([]map[string]interface{}, error) {
	var b Business
	var businessId int
	var revenue float64
	var datePaid string
	sqlQuery := "SELECT business_id, revenue, date_paid FROM tax WHERE tax_paid>0"
	results, err := db.Query(sqlQuery)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	var taxArray []map[string]interface{}
	for results.Next() {
		err = results.Scan(&businessId, &revenue, &datePaid)
		b.Id = businessId
		err := b.GetBusinessById()
		if err != nil {
			return taxArray, err
		}
		taxArray = append(taxArray, map[string]interface{}{
			"business_id": businessId,
			"revenue": revenue,
			"date_paid": datePaid,
			"business": b,
		})
	}
	return taxArray, err
}

func (t *Tax) GetTaxByDateAndId () (err error) {
	fmt.Println("tax period", t.TaxPeriod)
	sqlQuery := "SELECT * FROM tax WHERE business_id=? AND tax_period=?"
	result := db.QueryRow(sqlQuery, t.BusinessId, t.TaxPeriod)
	err = result.Scan(&t.Id, &t.BusinessId, &t.TaxPeriod, &t.Revenue,
		&t.TaxPaid, &t.DatePaid, &t.CreatedAt, &t.UpdateAt)
	return
}

func (t *Tax) UpdateTaxPayment () (err error) {
	stmt, err := db.Prepare("Update tax set date_paid=?, tax_paid=?, updated_at=NOW() where business_id=? AND tax_period=?")
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(t.DatePaid, t.TaxPaid, t.BusinessId, t.TaxPeriod)
	return err
}

func (t *Tax) IsTaxCompliant () {
	// todo
}

func (t *Tax) TaxCompliantList () {
	// todo
}

func (t *Tax) ListNonCompliant () {
	// todo
}