package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/models"
	"net/http"
	"strconv"
)

type taxHandler struct {
	handlerHelper
}

func taxRoute (r chi.Router) {
	t := taxHandler{}
	r.Post("/revenue", t.updateTaxRevenue)
	r.Post("/payment", t.updateTaxPayment)
	r.Get("/paid", t.getAllPaidTaxes)
	r.Get("/record", t.getBusinessTaxRecord)
	r.Get("/unpaid", t.getAllUnPaidTaxes)
	r.Get("/date", t.getData)
}

func (h *taxHandler) getBusinessTaxRecord (w http.ResponseWriter, r *http.Request) {
	businessId, _ := strconv.Atoi(r.URL.Query().Get("business_id"))
	t := models.Tax{
		BusinessId:businessId,
	}
	records, err := t.GetBusinessTaxRecord()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": records,
		})
	}
}

func (h *taxHandler) getAllPaidTaxes (w http.ResponseWriter, r *http.Request) {
	var tax models.Tax
	_ = json.NewDecoder(r.Body).Decode(&tax)
	paidTaxes, err := tax.GetAllPaidTaxes()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": paidTaxes,
		})
	}
}

func (h *taxHandler) getAllUnPaidTaxes (w http.ResponseWriter, r *http.Request) {
	var tax models.Tax
	_ = json.NewDecoder(r.Body).Decode(&tax)
	paidTaxes, err := tax.GetAllUnPaidTaxes()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": paidTaxes,
		})
	}
}

func (h *taxHandler) updateTaxPayment (w http.ResponseWriter, r *http.Request) {
	var tax models.Tax
	_ = json.NewDecoder(r.Body).Decode(&tax)
	err := tax.UpdateTaxPayment()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": map[string]interface{}{
				"message": "tax payment updated successfully",
				"revenue":tax.TaxPaid,
				"tax_period": tax.TaxPeriod,
			},
		})
	}
}


func (h *taxHandler) updateTaxRevenue (w http.ResponseWriter, r *http.Request) {
	var tax models.Tax
	_ = json.NewDecoder(r.Body).Decode(&tax)
	err := tax.SaveBusinessRevenue()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": map[string]interface{}{
				"message": "tax revenue updated successfully",
				"revenue":tax.Revenue,
				"tax_period": tax.TaxPeriod,
			},
		})
	}
}

func (h *taxHandler) getData (w http.ResponseWriter, r *http.Request) {
	businessId, _ := strconv.Atoi(r.URL.Query().Get("business_id"))
	taxPeriod := string(r.URL.Query().Get("tax_period"))
	t := models.Tax{
		BusinessId:businessId,
		TaxPeriod: taxPeriod,
	}
	err := t.GetTaxByDateAndId()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": t,
		})
	}
}
