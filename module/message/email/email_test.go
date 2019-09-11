package email

import (
	"context"
	"testing"
)

func TestSendEmail(t *testing.T) {
	err := SendEmail(context.Background(), "aa", "测试消息")
	t.Log(err)
}
