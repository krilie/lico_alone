package middleware

import (
	"github.com/krilie/lico_alone/common/comlog"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/user/auth"
	"github.com/krilie/lico_alone/module/user/info"
)

var apiUser info.User
var apiAuthUser auth.UserAuth
var log = comlog.NewLog(context.NewContext(), "alone.control.middleware", "init")
