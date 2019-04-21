package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sovikc/bsms/server"
)

const (
	defaultPort = "8000"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessToken := os.Getenv("BITLY_GENERIC_ACCESS_TOKEN")
	groupGUID := os.Getenv("BITLY_GROUP_GUID")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	log.Println("Loaded accessToken", accessToken)
	log.Println("Loaded groupGUID", groupGUID)
	log.Println("Loaded apiKey", apiKey)
	log.Println("Loaded apiSecret", apiSecret)

	var (
		httpAddr = ":" + defaultPort
	)

	srv := server.New()

	httpServer := &http.Server{Addr: httpAddr,
		Handler:      srv,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second}

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
}
