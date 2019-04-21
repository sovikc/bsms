package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sovikc/bsms/sms"

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

// Message is a value object for decoding
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

	message := sms.New(request.Phone, transform(request.Messages))
	status := h.msg.SendSMS(message)

	var response = struct {
		Status string `json:"status"`
	}{
		Status: status + " at " + time.Now().Format("Mon Jan 2 2006 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		encodeError(ctx, err, w)
		return
	}

}

func transform(messages []Message) []string {
	msgs := make([]string, 0)
	for _, v := range messages {
		msgs = append(msgs, v.Message)
	}
	return msgs
}
