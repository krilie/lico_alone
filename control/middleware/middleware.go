package middleware

import (
	"github.com/krilie/lico_alone/common/comlog"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/userbase/auth"
	"github.com/krilie/lico_alone/module/userbase/user"
)

var apiUser user.User
var apiAuthUser auth.UserAuth
var log = comlog.NewLog(context.NewContext(), "alone.control.middleware")
