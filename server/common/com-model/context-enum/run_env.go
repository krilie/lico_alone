package context_enum

// RunEnv is the app's run env
type RunEnv struct {
	AppName    string
	AppVersion string
	AppHost    string
	ClientId   string
	UserId     string
	Version    string
	BuildTime  string
	GoVersion  string
	GitCommit  string
}
