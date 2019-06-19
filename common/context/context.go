package context

import (
	"time"
)

// 运行上下文对象
// 可记录来自那个ip 用户的信息等
type context struct {
	TraceId       string    //微服务调用栈分析追踪
	StartTime     time.Time //开始调用时间
	LastTime      time.Time //调用结束时间
	ClientId      string    //client的id号
	ClientToken   string    //当前client的acctoken
	UserId        string
	UserToken     string //当前用户的acctoken jwt
	UserNickName  string
	UserLoginName string
}

func (c *context) GetTraceId() string {
	return c.TraceId
}

func (c *context) GetStartTime() time.Time {
	return c.StartTime
}

func (c *context) GetLastTime() time.Time {
	return c.LastTime
}

func (c *context) GetClientId() string {
	return c.ClientId
}

func (c *context) GetClientToken() string {
	return c.ClientToken
}

func (c *context) GetUserId() string {
	return c.UserId
}

func (c *context) GetUserToken() string {
	return c.UserToken
}

func (c *context) GetUserNickName() string {
	return c.UserNickName
}

func (c *context) GetUserLoginName() string {
	return c.UserLoginName
}

func (c *context) SetTraceId(staceId string) {
	c.TraceId = staceId
}

func (c *context) SetStartTime(start time.Time) {
	c.StartTime = start
}

func (c *context) SetLastTime(last time.Time) {
	c.LastTime = last
}

func (c *context) SetClientId(clientId string) {
	c.ClientId = clientId
}

func (c *context) SetClientToken(clientToken string) {
	c.ClientToken = clientToken
}

func (c *context) SetUserId(userId string) {
	c.UserId = userId
}

func (c *context) SetUserToken(userToken string) {
	c.UserToken = userToken
}

func (c *context) SetUserNickName(nickName string) {
	c.UserNickName = nickName
}

func (c *context) SetUserLoginName(loginName string) {
	c.UserLoginName = loginName
}
