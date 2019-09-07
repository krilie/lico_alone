package broker

import (
	"context"
	"fmt"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"sync"
)

type RegisterBroker interface{ RegisterBroker(ctx context.Context) }

var brokers map[string]*Broker
var mutex sync.Mutex

func Close() {
	for _, v := range brokers {
		v.Stop()
		log.Infof("brokers is closed:%v", v.Name)
	}
	log.Infoln("brokers is closed.")
}

func SendMessage(ctx context.Context, name string, msg interface{}) {
	broker, ok := brokers[name]
	if !ok {
		panic(fmt.Errorf("err get broker:%v when send msg:%v", name, str_util.ToJson(msg)))
	}
	broker.Send(msg)
	log.Debugf("broker:%v send msg:%v", broker.Name, str_util.ToJson(msg))
}

func RegisterHandler(ctx context.Context, name string, f func(p interface{})) {
	broker, ok := brokers[name]
	if !ok {
		mutex.Lock()
		defer mutex.Unlock()
		broker = NewStartedBroker(name, 1)
		brokers[name] = broker
		log.Infof("brokers created:%v", broker.Name)
	}
	broker.Register(f)
}
