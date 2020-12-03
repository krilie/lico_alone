package main

import (
	context2 "context"
	"fmt"
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	run_env "github.com/krilie/lico_alone/common/run_env"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module"
	service2 "github.com/krilie/lico_alone/module/service"
	"github.com/krilie/lico_alone/server"
	"github.com/krilie/lico_alone/server/http"
	_ "go.uber.org/automaxprocs"
	"os"
)

//go:generate swag init -g ./main.go
//go:generate go test -run Auto -v ./...

// @title lizo_alone
// @version 1.0.0
// @description  api docs for lizo_alone
// @license.name all right

// @host localhost:80
// @BasePath /

func main() {
	// cmd
	if CmdForInfo() {
		return
	}
	// dig config
	var container = appdig.
		NewAppDig().
		MustProvides(component.DigComponentProviderAll). // 基础组件
		MustProvides(module.DigProviderModuleAll).       // 功能模块组件
		MustProvides(service2.DigServiceProviderAll).    // 服务组件
		MustProvides(http.DigControllerProviderAll).     // http controller组件
		MustProvides(server.DigServerProviderAll)        // app组件
	// begin service
	container.MustInvoke(func(svc *server.Server, log *nlog.NLog) {

		ctxValues := context.NewAppCtxValues()
		ctxValues.Module = "main"
		ctxValues.Function = "main"
		ctxValues.UserId = "main"
		ctx := context.NewAppCtx(context2.Background(), ctxValues)

		svc.StartService(ctx)
		log.Get(ctx).Warning("service done")
	})
}

// 命令行处理 没用
func CmdForInfo() (hasCmd bool) { // 命令行 命令
	if len(os.Args) >= 2 {
		cmd := os.Args[1]
		switch cmd {
		case "version", "--version", "-version":
			fmt.Println(run_env.VERSION)
			return true
		case "git-commit", "--git-commit", "-git-commit":
			fmt.Println(run_env.GIT_COMMIT)
			return true
		case "go-version", "-go-version", "--go-version":
			fmt.Println(run_env.GO_VERSION)
			return true
		case "build-time", "-build-time", "--build-time":
			fmt.Println(run_env.BUILD_TIME)
			return true
		default:
			return true
		}
	} else {
		return false
	}
}
