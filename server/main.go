package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc"

	"github.com/mp-hl-2021/grpc-example/api"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterEchoServer(grpcServer, &EchoServer{})

	if err := grpcServer.Serve(l); err != nil {
		panic(err)
	}
}

type EchoServer struct {
	api.UnimplementedEchoServer
}

func (e *EchoServer) Do(_ context.Context, r *api.EchoRequest) (*api.EchoReply, error) {
	reply := strings.ToUpper(r.Line)
	reply = fmt.Sprintf("((( %s ))) %d", reply, r.Num)
	return &api.EchoReply{EchoLine: reply}, nil
}
