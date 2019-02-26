package models

import (
	"fmt"
)

type Tax struct {
	Id int `json:"id"`
	BusinessId int `json:"business_id"`
	TaxPeriod *string `json:"tax_period"`
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

func (t *Tax) UpdateTaxPayment () {

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