package messaging

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/sovikc/bsms/sms"
)

const (
	smsURL = "https://api.transmitsms.com/send-sms.json"
)

// Service is the interface that provides sms service
type Service interface {
	SendSMS(*sms.Message) string
}

type service struct {
	apiKey       string
	apiSecret    string
	urlShortener sms.URLShortener
}

// NewService creates a sms service with necessary dependencies
func NewService(apiKey, apiSecret string, urlShortener sms.URLShortener) Service {
	return &service{
		apiKey:       apiKey,
		apiSecret:    apiSecret,
		urlShortener: urlShortener,
	}
}

type serviceResponse struct {
	SentAt string
	Error  error
}

func (s *service) hasAPIKey() bool {
	return len(s.apiKey) > 0
}

func (s *service) hasAPISecret() bool {
	return len(s.apiSecret) > 0
}

// SendSMS is used to send sms
func (s *service) SendSMS(message *sms.Message) string {
	msgCount := len(message.Messages)
	texts := make(chan sms.Text, msgCount)
	responses := make(chan serviceResponse, msgCount)

	if !s.hasAPIKey() || !s.hasAPISecret() {
		return fmt.Sprintf("Sent SMS with %d successes and %d failures", 0, msgCount)
	}

	for i := 0; i < msgCount; i++ {
		go s.sender(texts, responses)
	}

	for j := 0; j < msgCount; j++ {
		text := sms.Text{
			Phone: message.Phone,
			Body:  message.Messages[j]}
		texts <- text
	}
	close(texts)

	var success, fail int
	for r := 0; r < msgCount; r++ {
		response := <-responses
		if response.Error != nil {
			fail++
			log.Println("SMS send error", response.Error)
			continue
		}
		success++
	}

	response := fmt.Sprintf("Sent SMS with %d successes and %d failures", success, fail)
	return response

}

func (s *service) shortenURLs(messageBody sms.MessageBody) (string, error) {
	var sb strings.Builder
	var index int
	for _, urlIndex := range messageBody.URLIndices {
		sb.WriteString(messageBody.Content[index:urlIndex.Start])

		url := messageBody.Content[urlIndex.Start:urlIndex.End]
		shortURL, err := s.urlShortener.GetShortenedURL(url)
		if err != nil {
			return messageBody.Content, err
		}

		index = urlIndex.End
		sb.WriteString(shortURL)
	}
	sb.WriteString(messageBody.Content[index:len(messageBody.Content)])
	return sb.String(), nil

}

func (s *service) sender(texts <-chan sms.Text, responses chan<- serviceResponse) {
	for text := range texts {

		var smsText = text.Body.Content
		sr := serviceResponse{}
		var err error

		if text.Body.HasURL {
			smsText, err = s.shortenURLs(text.Body)
			if err != nil {
				sr.Error = err
				responses <- sr
				return
			}
		}

		var sb strings.Builder
		sb.WriteString("message=")
		sb.WriteString(smsText)
		sb.WriteString("&to=")
		sb.WriteString(text.Phone)
		body := strings.NewReader(sb.String())

		req, err := http.NewRequest("POST", smsURL, body)
		if err != nil {
			sr.Error = err
			responses <- sr
			return
		}

		req.SetBasicAuth(s.apiKey, s.apiSecret)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			sr.Error = err
			responses <- sr
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			sr.Error = errors.New("SMS service returned an error")
			responses <- sr
			return
		}

		var response struct {
			SentAt string `json:"send_at"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			sr.Error = err
			responses <- sr
			return
		}

		if len(strings.Trim(response.SentAt, " ")) == 0 {
			sr.Error = errors.New("SMS service returned an error")
			responses <- sr
			return
		}

		sr.SentAt = response.SentAt
		responses <- sr
		return

	}

}
