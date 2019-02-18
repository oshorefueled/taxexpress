package handlers

import (
	"github.com/go-chi/chi"
	"net/http"
)

type authHandler struct {

}

func authRoute (r chi.Router) {
	h := authHandler{}
	r.Get("/login", h.login)
}

func (h authHandler) login (w http.ResponseWriter, r *http.Request) {

}