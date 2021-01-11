package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/broker/messages"
	"testing"
	"time"
)

func TestBlogArticleModule_HandleBrokerBlogArticleVisitedMessage(t *testing.T) {
	container.MustInvoke(func(svc *BlogArticleModule) {
		svc.broker.Send(context.EmptyAppCtx(), &messages.BlogArticleVisitedMessage{
			Ctx:         context.EmptyAppCtx(),
			VisitedTime: time.Now(),
			ArticleId:   "11",
			VisitorIp:   "22",
		})
		svc.broker.Close()
	})
}
