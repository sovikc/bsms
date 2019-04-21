package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sovikc/bsms/bitly"
	"github.com/sovikc/bsms/messaging"
	"github.com/sovikc/bsms/server"
	"github.com/sovikc/bsms/sms"
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

	var (
		httpAddr = ":" + defaultPort
		ms       messaging.Service
		us       sms.URLShortener
	)

	us = bitly.NewURLShortener(accessToken, groupGUID)
	ms = messaging.NewService(apiKey, apiSecret, us)
	srv := server.New(ms)

	httpServer := &http.Server{Addr: httpAddr,
		Handler:      srv,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second}

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
}
