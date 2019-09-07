package broker

import (
	"context"
	"fmt"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/sirupsen/logrus"
	"sync"
)

// broker 生产者 消费者
// 简单模拟消息队列

var log = clog.NewLog(context.Background(), "lico_alone.common.broker", "brokers")

type Broker struct {
	event     chan interface{}
	funcs     []func(interface{})
	muFuncs   sync.RWMutex
	Name      string
	wait      *sync.WaitGroup
	onceStart sync.Once
	onceStop  sync.Once
}

// 创建broker,并开始
func NewStartedBroker(name string, chanBuf int) *Broker {
	b := &Broker{
		event: make(chan interface{}, chanBuf),
		funcs: make([]func(interface{}), 0),
		Name:  name,
		wait:  &sync.WaitGroup{},
	}
	b.Start()
	log.Infof("broker %v created and started", b.Name)
	return b
}

// 注册事件
func (b *Broker) Send(o interface{}) (err error) {
	defer func() {
		if errs := recover(); errs != nil {
			err = fmt.Errorf("%v", errs)
		}
	}()
	b.event <- o
	return nil
}

// 注册事件
func (b *Broker) Register(f func(interface{})) {
	b.muFuncs.Lock()
	defer b.muFuncs.Unlock()
	b.funcs = append(b.funcs, func(o interface{}) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Error(err)
			}
		}()
		f(o)
	})
}

func (b *Broker) Clear(f func(interface{})) {
	b.muFuncs.Lock()
	defer b.muFuncs.Unlock()
	b.funcs = b.funcs[0:0]
}

// 调用stop之前确保写入方都已经退出了,不然要panic
func (b *Broker) Stop() {
	b.onceStop.Do(func() {
		close(b.event)
		b.wait.Wait()
		log.Infof("broker %v has stop success", b.Name)
	})
}

func (b *Broker) Start() {
	b.onceStart.Do(func() {
		b.wait.Add(1)
		go func() {
			for {
				event, ok := <-b.event
				if ok {
					// 事件分发
					b.muFuncs.RLock()
					for _, v := range b.funcs {
						v(event) // 有recover
					}
					if len(b.funcs) == 0 {
						log.Infof("broker %v no events %v", b.Name, str_util.ToJson(event))
					}
					b.muFuncs.RUnlock()
				} else {
					// 通道已经关闭
					b.wait.Done()
					return
				}
			}
		}()
		log.Infof("broker %v init success", b.Name)
	})
}
