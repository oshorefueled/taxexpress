package handlers

import (
	"encoding/json"
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

func (h mailHandler) sendMail (w http.ResponseWriter, r *http.Request) {
	var mail = h.MailHelper
	_ = json.NewDecoder(r.Body).Decode(&mail)
	err := mail.SendMail()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"body": mail.Body,
			"to": mail.ToIDs,
			"subject": mail.Subject,
		})
	}
}