package context_util

//运行上下文对象

type Context struct {
	StackId string //微服务调用栈分析追踪
	//开始时间和结束时间打在日志上这里不要加
	//StartTime time.Time //开始调用时间
	//LastTime time.Time //调用结束时间
	Auth *AuthInfo //一些认证信息，可以为nil
}

type AuthInfo struct {
	AppId    *string //认证的端的id 可以为nil
	UserId   *string //可以为nil 登录用户的id
	UserType *string //登录用户的类型
	NickName *string //登录用户的昵称
}

func (ctx *Context) GetAppIdOrEmpty() string {
	if ctx.Auth != nil && ctx.Auth.AppId != nil {
		return *ctx.Auth.AppId
	} else {
		return ""
	}
}

func (ctx *Context) GetUserIdOrEmpty() string {
	if ctx.Auth != nil && ctx.Auth.UserId != nil {
		return *ctx.Auth.UserId
	} else {
		return ""
	}
}

func (ctx *Context) GetUserIdOrDefault(def string) string {
	if ctx.Auth != nil && ctx.Auth.UserId != nil {
		return *ctx.Auth.UserId
	} else {
		return def
	}
}
