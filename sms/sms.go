package sms

import "regexp"

// URLIndex has the start and end index of a URL
type URLIndex struct {
	Start int
	End   int
}

// MessageBody has the content and other details of the message
type MessageBody struct {
	HasURL     bool
	URLIndices []URLIndex
	Content    string
}

// Message has basic details required to send text message
type Message struct {
	Phone    string
	Messages []MessageBody
}

// Text represents an individual message
type Text struct {
	Phone string
	Body  MessageBody
}

// URLShortener shortens a long url still directing to the original page
type URLShortener interface {
	GetShortenedURL(longURL string) (string, error)
}

// New returns a new instance of Message
func New(phone string, messages []string) *Message {
	msgs := make([]MessageBody, 0)

	for _, v := range messages {
		urlIndices := getURLInfo(v)
		msgBody := MessageBody{
			URLIndices: urlIndices,
			Content:    v,
			HasURL:     len(urlIndices) > 0}
		msgs = append(msgs, msgBody)
	}

	return &Message{
		Phone:    phone,
		Messages: msgs}
}

func getURLInfo(message string) []URLIndex {
	re := regexp.MustCompile(`(https|http):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	indices := re.FindAllIndex([]byte(message), -1)
	urlIndices := make([]URLIndex, 0)

	for _, v := range indices {
		urlIndices = append(urlIndices, URLIndex{Start: v[0], End: v[1]})
	}

	return urlIndices
}
