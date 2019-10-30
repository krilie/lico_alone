package ccontext

import (
	"time"
)

// 运行上下文对象
// 可记录来自那个ip 用户的信息等
type Context struct {
	TraceId   string    //微服务调用栈分析追踪
	StartTime time.Time //开始调用时间
	LastTime  time.Time //调用结束时间
	ClientId  string    //client的id号
	UserId    string
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Context) Done() <-chan struct{} {
	return nil
}

func (c *Context) Err() error {
	return nil
}

func (c *Context) Value(key interface{}) interface{} {
	return nil
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
func (c *Context) GetUserId() string {
	return c.UserId
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
func (c *Context) SetUserId(userId string) {
	c.UserId = userId
}
