package logsyshook

import "github.com/krilie/lico_alone/common/dig"

func init() {
	dig.Container.MustProvide(NewSyslogHook)
	dig.Container.MustProvide(NewElfLogHook)
}
