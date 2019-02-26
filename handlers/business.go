package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/models"
	"net/http"
	"strconv"
)

type businessHandler struct {
	handlerHelper
}

func businessRoute (r chi.Router) {
	b := businessHandler{}
	r.Post("/create", b.createBusiness)
	r.Get("/all", b.getBusinesses)
	r.Get("/", b.getBusiness)
}

func (h *businessHandler) createBusiness (w http.ResponseWriter, r *http.Request) {
	var business models.Business
	_ = json.NewDecoder(r.Body).Decode(&business)
	fmt.Println(business)
	_, _, err := business.StoreBusiness()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": map[string]string{
				"message": "business created successfully",
			},
		})
	}
}

func (h *businessHandler) getBusinesses (w http.ResponseWriter, r *http.Request) {
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