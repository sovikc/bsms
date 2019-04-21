package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/sovikc/bsms/messaging"
)

type messagingHandler struct {
	msg messaging.Service
}

func (h *messagingHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/sms", func(r chi.Router) {
		r.Post("/", h.sendSMS)
	})

	return r
}

type Message struct {
	Message string
}

func (h *messagingHandler) sendSMS(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request struct {
		Phone    string
		Messages []Message
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("error reading sms post", err)
		encodeError(ctx, err, w)
		return
	}

	fmt.Println(request.Phone, request.Messages)

	var response = struct {
		Status string `json:"status"`
	}{
		Status: "Sent at " + time.Now().Format("Mon Jan 2 2006 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		encodeError(ctx, err, w)
		return
	}

}