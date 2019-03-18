package handlers

import (
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/helpers"
	"net/http"
)

type mailHandler struct {
	handlerHelper
	helpers.MailHelper
}

func mailRoute (r chi.Router) {
	m := mailHandler{}
	r.Post("/send", m.sendMail)
}

func (m mailHandler) sendMail (w http.ResponseWriter, r *http.Request) {

}