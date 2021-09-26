package echoserver

import (
	"context"

	spb "github.com/cloneable/go-microservice-template/api/proto/server"
	"go.uber.org/zap"
)

type EchoServer struct {
	Logger *zap.SugaredLogger
}

var _ spb.EchoServiceServer = (*EchoServer)(nil)

func (s *EchoServer) Echo(ctx context.Context, req *spb.EchoRequest) (*spb.EchoResponse, error) {
	s.Logger.Infof("Message received: %v", req.Msg)
	return &spb.EchoResponse{Msg: req.Msg}, nil
}
