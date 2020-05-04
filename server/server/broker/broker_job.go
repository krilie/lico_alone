package broker

import (
	"context"
	"github.com/krilie/lico_alone/service"
)

// 消息处理工作
func RegisterHandler(ctx context.Context, app *service.App) {
	app.All.UserService.RegisterBroker(ctx)
	app.All.FileService.RegisterBroker(ctx)
	app.All.AccountService.RegisterBroker(ctx)
}
