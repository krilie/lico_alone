package service

import (
	"github.com/krilie/lico_alone/component/broker/messages"
	"time"
)

// HandleBrokerBlogArticleVisitedMessage 文章被访问的消息
func (s *BlogArticleModule) HandleBrokerBlogArticleVisitedMessage(msg *messages.BlogArticleVisitedMessage) {
	log := s.log.Get(msg.GetCtx()).WithField("param_msg", msg)
	log.Info("handle a broker of blog article visited message")
	err := s.Dao.GetDb(msg.GetCtx()).
		Exec("update tb_article_master set updated_at=?,pv=pv+1 where deleted_at is null and id=?", time.Now(), msg.ArticleId).Error
	if err != nil {
		log.WithField("err", err).Error("update pv err")
	}
	log.Info("update pv success")
}
