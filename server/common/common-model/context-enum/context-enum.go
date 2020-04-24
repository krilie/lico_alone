package context_enum

type ContextEnum string

const (
	AppName    = "app_name"
	AppVersion = "app_version"
	AppHost    = "app_host"
	TraceId    = "trace_id"
	ClientId   = "client_id"
	UserId     = "user_id"
	Module     = "module"
	Function   = "function"
	Stack      = "stack"
)
