package email

import (
	"context"
	"testing"
)

func TestSendEmail(t *testing.T) {
	err := SendServiceUpEmail(context.Background(), "测试消息")
	t.Log(err)
}
