package echoserver

import (
	"context"

	spb "github.com/cloneable/go-microservice-template/api/proto/server"
	"github.com/golang/glog"
)

type EchoServer struct {
	spb.UnimplementedEchoServiceServer
}

var _ spb.EchoServiceServer = (*EchoServer)(nil)

func (s *EchoServer) Echo(ctx context.Context, req *spb.EchoRequest) (*spb.EchoResponse, error) {
	glog.Infof("Message received: %v", req.Msg)
	return &spb.EchoResponse{Msg: req.Msg}, nil
}
