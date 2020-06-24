package messages

import (
	"context"
	"reflect"
	"time"
)

// 文章被访问的消息
type BlogArticleVisitedMessage struct {
	Ctx         context.Context
	VisitedTime time.Time
	ArticleId   string
	VisitorIp   string
}

func (b *BlogArticleVisitedMessage) GetName() string {
	return "BlogArticleVisitedMessage"
}

func (b *BlogArticleVisitedMessage) GetCtx() context.Context {
	return b.Ctx
}

func (b *BlogArticleVisitedMessage) GetType() reflect.Type {
	return reflect.TypeOf(b)
}
