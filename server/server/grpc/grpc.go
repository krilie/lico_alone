package grpc

import (
	"context"
	"github.com/krilie/lico_alone/server/grpc/proto"
)

//go:generate protoc --go_opt=paths=source_relative --proto_path=./ --go_out=plugins=grpc:./ ./proto/test.proto

type GRpcService struct{}

func (G *GRpcService) Search(ctx context.Context, test *proto.Test) (*proto.Test, error) {
	return test, nil
}
