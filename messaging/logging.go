package messaging

import (
	"log"
	"time"

	"github.com/sovikc/bsms/sms"
)

type loggingService struct {
	next Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(s Service) Service {
	return &loggingService{s}
}

func (s *loggingService) SendSMS(message *sms.Message) string {
	defer func(begin time.Time) {
		log.Println("method", "SendSMS", "took", time.Since(begin))
	}(time.Now())
	return s.next.SendSMS(message)
}
