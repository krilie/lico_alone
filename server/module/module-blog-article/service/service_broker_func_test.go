package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/broker/messages"
	"testing"
	"time"
)

func TestBlogArticleModule_HandleBrokerBlogArticleVisitedMessage(t *testing.T) {
	dig.Container.MustInvoke(func(svc *BlogArticleModule) {
		svc.broker.Send(context.NewContext(), &messages.BlogArticleVisitedMessage{
			Ctx:         context.NewContext(),
			VisitedTime: time.Now(),
			ArticleId:   "11",
			VisitorIp:   "22",
		})
		svc.broker.Close()
	})
}
