package context_enum

type ContextEnum string

func (c ContextEnum) Str() string {
	return string(c)
}

const (
	AppName    ContextEnum = "app_name"
	AppVersion ContextEnum = "app_version"
	AppHost    ContextEnum = "app_host"
	TraceId    ContextEnum = "trace_id"
	ClientId   ContextEnum = "client_id"
	UserId     ContextEnum = "user_id"
	Module     ContextEnum = "module"
	Function   ContextEnum = "function"
	Stack      ContextEnum = "stack"
	RemoteIp   ContextEnum = "remote_ip"
)

type ContextValues struct {
	AppName    string
	AppVersion string
	AppHost    string
	TraceId    string
	ClientId   string
	UserId     string
	Module     string
	Function   string
	Stack      string
	RemoteIp   string
}
