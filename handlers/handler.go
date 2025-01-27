package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func InitHandlers (route *chi.Mux) {
	route.Get("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to TaxExpress"))
	})
	route.Route("/auth", authRoute)
	route.Route("/business", businessRoute)
	route.Route("/tax", taxRoute)
	route.Route("/message", messageRoute)
	route.Route("/admin", adminRoute)
	route.Route("/mail", mailRoute)
}

type handlerHelper struct {

}

func (h handlerHelper) catch(err error) {
	if err != nil {
		panic(err)
	}
}

func (h handlerHelper) RespondWithError(w http.ResponseWriter, code int, msg string) {
	h.RespondWithJSON(w, code, map[string]string{"status": "error", "message": msg})
}

func (h handlerHelper) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
