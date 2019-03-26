package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/models"
	"net/http"
	"strconv"
)

type businessHandler struct {
	handlerHelper
}

func businessRoute (r chi.Router) {
	r.Use(needTokenMiddleware)
	b := businessHandler{}
	r.Post("/create", b.createBusiness)
	r.Get("/all", b.getAllBusinesses)
	r.Get("/", b.getBusiness)
}

func (h *businessHandler) createBusiness (w http.ResponseWriter, r *http.Request) {
	var business models.Business
	_ = json.NewDecoder(r.Body).Decode(&business)
	_, _, err := business.StoreBusiness()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		_ = business.GetBusinessByEmail()
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": map[string]interface{}{
				"message": "business created successfully",
				"business": business,
			},
		})
	}
}

func (h *businessHandler) getAllBusinesses (w http.ResponseWriter, r *http.Request) {
	b := models.Business{}
	businesses, err := b.GetAllBusinesses()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": businesses,
		})
	}
}

func (h *businessHandler) getBusiness (w http.ResponseWriter, r *http.Request) {
	b := models.Business{}
	businessId, _ := strconv.Atoi(r.URL.Query().Get("id"))
	b.Id = businessId
	err := b.GetBusinessById()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": b,
		})
	}
}