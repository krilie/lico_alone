// +build !auto_test

package grpc

import (
	"context"
	"github.com/krilie/lico_alone/server/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

const (
	port = ":50051"
)

func TestGRpcService_Search(t *testing.T) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterSearchServiceServer(s, &GRpcService{})
	println("服务端开始")
	s.Serve(lis)
	println("服务端退出")
}

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func TestGRpcService_Search2(t *testing.T) {
	// 准备连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewSearchServiceClient(conn)
	// 请求接口
	name := defaultName
	r, err := c.Search(context.Background(), &proto.Test{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Name)
}
