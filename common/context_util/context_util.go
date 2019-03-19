package context_util

import (
	"github.com/lico603/lico-my-site-user/common/jwt"
	"time"
)

// 运行上下文对象
// 可记录来自那个ip 用户的信息等
type Context struct {
	TraceId string //微服务调用栈分析追踪
	//开始时间和结束时间打在日志上这里不要加
	StartTime      time.Time //开始调用时间
	LastTime       time.Time //调用结束时间
	ClientId       *string   //client的id号
	ClientAccToken *string   //当前client的acctoken
	NowUserToken   *string   //当前用户的acctoken

	UserClaims *jwt.UserClaims //一些认证信息，可以为nil
}

func (ctx *Context) GetAppIdOrEmpty() string {
	if ctx.UserClaims != nil {
		return ctx.UserClaims.AppId
	} else {
		return ""
	}
}

func (ctx *Context) GetUserIdOrEmpty() string {
	if ctx.UserClaims != nil {
		return ctx.UserClaims.UserId
	} else {
		return ""
	}
}

func (ctx *Context) GetUserIdOrDefault(def string) string {
	if ctx.UserClaims != nil {
		return ctx.UserClaims.UserId
	} else {
		return def
	}
}

func (ctx *Context) GetClientIdOrEmpty() string {
	if ctx.ClientId != nil {
		return *ctx.ClientId
	} else {
		return ""
	}
}

func (ctx *Context) GetNowUserTokenOrEmpty() string {
	if ctx.NowUserToken != nil {
		return *ctx.NowUserToken
	} else {
		return ""
	}
}

func MustGetContext(ctx interface{}) *Context {
	if c, ok := ctx.(*Context); !ok {
		panic("err convert from interface{} to *Context")
	} else {
		return c
	}
}

func GetContextOrNil(ctx interface{}) *Context {
	if c, ok := ctx.(*Context); !ok {
		return nil
	} else {
		return c
	}
}
