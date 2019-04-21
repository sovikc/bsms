package server

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/sovikc/bsms/messaging"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	Messaging messaging.Service
	router    chi.Router
}

// New returns a new HTTP server.
func New(ms messaging.Service) *Server {
	s := &Server{
		Messaging: ms,
	}

	r := chi.NewRouter()
	r.Use(basicHeaders)
	r.Use(serveStatic)
	r.Use(middleware.Recoverer)

	r.Route("/messaging", func(r chi.Router) {
		msg := messagingHandler{s.Messaging}
		r.Mount("/v1", msg.router())
	})

	r.Get("/", index)

	s.router = r
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func basicHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

// Encode error for response
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}

func serveStatic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fPath := filepath.Clean(r.URL.Path)
		mainjsRequested, err := regexp.MatchString(`^/main.*.js$`, fPath)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		vendorjsRequested, err := regexp.MatchString(`^/vendor.*.js$`, fPath)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		if mainjsRequested || vendorjsRequested {
			fp := filepath.Join("static", fPath)
			http.ServeFile(w, r, fp)
			return
		}

		h.ServeHTTP(w, r)
	})
}
