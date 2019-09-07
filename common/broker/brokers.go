package broker

import (
	"context"
	"errors"
	"time"
)

type RegisterBroker interface{ RegisterBroker(ctx context.Context) }

var brokers map[string]*Broker

func InitBrokers(ctx context.Context) (close func()) {
	brokers[MemberRegistered] = NewStartedBroker(MemberRegistered, 0)

	return func() {
		for _, v := range brokers {
			v.Stop()
		}
	}
}

func GetBroker(ctx context.Context, name string) *Broker {
	broker, ok := brokers[name]
	if !ok {
		panic(errors.New("no this broker: " + name))
	}
	return broker
}

// 用户注册
const MemberRegistered = "member_registered"

type MemberRegisteredMsg struct {
	MemberId     string
	RegisterTime time.Time
}
