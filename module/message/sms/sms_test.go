package sms

import (
	"context"
	"testing"
)

func TestSendSms(t *testing.T) {
	err := SendSms(context.Background(), "test")
	t.Log(err)
}
