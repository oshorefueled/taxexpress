package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/models"
	"net/http"
)

type taxHandler struct {
	handlerHelper
}

func taxRoute (r chi.Router) {
	t := taxHandler{}
	r.Post("/revenue", t.updateTaxRevenue)
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
