package sms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	messages := []string{
		"This is a test message with a url http://www.abcd.com and some more text",
		"Another message with 3 urls https://xyz.co.uk/page1.html some more text, http://www.llvm.com/dfghj-yuiyuiyui, and http://www.tyu.co",
		"Third message with a url http://ppp.com",
	}
	message := New("61444444444", messages)
	for _, msg := range message.Messages {
		assert.Equal(t, true, msg.HasURL, "they should be equal")
	}
}
