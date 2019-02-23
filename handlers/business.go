package handlers

import (
	"github.com/go-chi/chi"
	"net/http"
)

type businessHandler struct {
	handlerHelper
}

func businessRoute (r chi.Router) {
	b := businessHandler{}
	r.Post("/create", b.createBusiness)
}

func (b *businessHandler) createBusiness (w http.ResponseWriter, r *http.Request) {

}