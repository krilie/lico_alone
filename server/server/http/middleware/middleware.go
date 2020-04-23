package middleware

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/nlog"
)

var log = log.NewLog(context.NewContext(), "alone.control.middleware", "init")
