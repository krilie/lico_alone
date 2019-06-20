package context

import (
	"context"
	"time"
)

// 运行上下文对象
// 可记录来自那个ip 用户的信息等
type Context struct {
	context.Context
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

func (c *Context) GetTraceId() string {
	return c.TraceId
}

func (c *Context) GetStartTime() time.Time {
	return c.StartTime
}

func (c *Context) GetLastTime() time.Time {
	return c.LastTime
}

func (c *Context) GetClientId() string {
	return c.ClientId
}

func (c *Context) GetClientToken() string {
	return c.ClientToken
}

func (c *Context) GetUserId() string {
	return c.UserId
}

func (c *Context) GetUserToken() string {
	return c.UserToken
}

func (c *Context) GetUserNickName() string {
	return c.UserNickName
}

func (c *Context) GetUserLoginName() string {
	return c.UserLoginName
}

func (c *Context) SetTraceId(staceId string) {
	c.TraceId = staceId
}

func (c *Context) SetStartTime(start time.Time) {
	c.StartTime = start
}

func (c *Context) SetLastTime(last time.Time) {
	c.LastTime = last
}

func (c *Context) SetClientId(clientId string) {
	c.ClientId = clientId
}

func (c *Context) SetClientToken(clientToken string) {
	c.ClientToken = clientToken
}

func (c *Context) SetUserId(userId string) {
	c.UserId = userId
}

func (c *Context) SetUserToken(userToken string) {
	c.UserToken = userToken
}

func (c *Context) SetUserNickName(nickName string) {
	c.UserNickName = nickName
}

func (c *Context) SetUserLoginName(loginName string) {
	c.UserLoginName = loginName
}
