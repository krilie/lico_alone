package service_user

import "context"

type InitAppData struct {
	Version   string `json:"version"`
	GitSha    string `json:"git_sha"`
	BuildTime string `json:"build_time"`
}

func (a *UserService) InitAppData(ctx context.Context) (data *InitAppData) {
	return &InitAppData{
		Version:   a.NCfg.RunEnv.Version,
		GitSha:    a.NCfg.RunEnv.GitCommit,
		BuildTime: a.NCfg.RunEnv.BuildTime,
	}
}
