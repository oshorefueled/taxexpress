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
	var taxPaid float64
	var taxPeriod string
	sqlQuery := "SELECT business_id, revenue, date_paid, tax_paid, tax_period FROM tax WHERE tax_paid>0"
	results, err := db.Query(sqlQuery)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	var taxArray []map[string]interface{}
	for results.Next() {
		err = results.Scan(&businessId, &revenue, &datePaid, &taxPaid, &taxPeriod)
		b.Id = businessId
		err := b.GetBusinessById()
		if err != nil {
			return taxArray, err
		}
		taxArray = append(taxArray, map[string]interface{}{
			"business_id": businessId,
			"revenue": revenue,
			"date_paid": datePaid,
			"tax_paid": taxPaid,
			"tax_period": taxPeriod,
			"business": b,
		})
	}
	return taxArray, err
}

func (t Tax) GetAllUnPaidTaxes () ([]map[string]interface{}, error) {
	var b Business
	var businessId int
	var revenue float64
	var taxPeriod string
	sqlQuery := "SELECT business_id, revenue, tax_period FROM tax WHERE tax_paid=0"
	results, err := db.Query(sqlQuery)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	var taxArray []map[string]interface{}
	for results.Next() {
		err = results.Scan(&businessId, &revenue, &taxPeriod)
		b.Id = businessId
		err := b.GetBusinessById()
		if err != nil {
			return taxArray, err
		}
		taxArray = append(taxArray, map[string]interface{}{
			"business_id": businessId,
			"revenue": revenue,
			"tax_period": taxPeriod,
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

func (t Tax) GetTaxRecordById () (interface{}, error) {
	var b Business
	sqlQuery := "SELECT * FROM tax WHERE business_id=?"
	result := db.QueryRow(sqlQuery, t.Id)
	err := result.Scan(&t.Id, &t.BusinessId, &t.TaxPeriod, &t.Revenue,
		&t.TaxPaid, &t.DatePaid, &t.CreatedAt, &t.UpdateAt)
	if err != nil {
		return map[string]interface{}{}, err
	}
	b.Id = t.BusinessId
	err = b.GetBusinessById()
	if err != nil {
		return map[string]interface{}{}, err
	}
	return map[string]interface{}{
		"id": t.Id,
		"business_id": t.BusinessId,
		"tax_period": t.TaxPeriod,
		"revenue": t.Revenue,
		"tax_paid": t.TaxPaid,
		"date_paid": t.DatePaid,
		"created_at": t.CreatedAt,
		"updated_at": t.UpdateAt,
		"business": b,
	}, err
}


func (t Tax) GetBusinessTaxRecord () ([]Tax, error) {
	fmt.Println("tax period", t.TaxPeriod)
	sqlQuery := "SELECT * FROM tax WHERE business_id=?"
	results, err := db.Query(sqlQuery, t.BusinessId)
	if err != nil {
		return []Tax{}, err
	}
	var records []Tax
	for results.Next() {
		err = results.Scan(&t.Id, &t.BusinessId, &t.TaxPeriod, &t.Revenue,
			&t.TaxPaid, &t.DatePaid, &t.CreatedAt, &t.UpdateAt)
		records = append(records, t)
	}
	return records, err
}

func (t *Tax) UpdateTaxPayment () (err error) {
	stmt, err := db.Prepare("Update tax set date_paid=?, tax_paid=?, " +
		"updated_at=NOW() where business_id=? AND tax_period=?")
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
