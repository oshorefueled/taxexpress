package handlers

import (
	"github.com/go-chi/chi"
	"net/http"
)

func InitHandlers (route *chi.Mux) {
	route.Get("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to TaxExpress"))
	})

	route.Get("/auth", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the auth page"))
	})
}
