package handlers

import (
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/models"
	"net/http"
)

type adminHandler struct {
	handlerHelper
}

func adminRoute (r chi.Router) {
	m := messageHandler{}
	r.Get("/all", m.getAllAdmins)
}

func (h *messageHandler) getAllAdmins (w http.ResponseWriter, r *http.Request) {
	a := models.Admin{}
	admins, err := a.GetAllAdminUsers()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": admins,
		})
	}
}