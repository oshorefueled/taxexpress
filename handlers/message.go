package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/models"
	"net/http"
)

type messageHandler struct {
	handlerHelper
}

func messageRoute (r chi.Router) {
	m := messageHandler{}
	r.Use(needTokenMiddleware)
	r.Post("/", m.createMessage)
	r.Get("/", m.getAllMessages)
}

func (h *messageHandler) createMessage (w http.ResponseWriter, r *http.Request) {
	message := models.Message{}
	_ = json.NewDecoder(r.Body).Decode(&message)
	_, _, err := message.CreateMessage()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": map[string]interface{}{
				"message": "tax message created successfully",
				"data": map[string]interface{}{
					"message": message.Message,
					"type": message.Type,
				},
			},
		})
	}
}

func (h *messageHandler) getAllMessages (w http.ResponseWriter, r *http.Request) {
	m := models.Message{}
	messages, err := m.GetAllMessages()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": messages,
		})
	}
}