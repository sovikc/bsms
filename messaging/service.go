package messaging

// Service is the interface that provides sms service
type Service interface {}

type service struct {}

// NewService creates a sms service
func NewService() Service {
	return &service{}
}