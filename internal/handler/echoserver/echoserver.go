package echoserver

import (
	"context"
	"fmt"

	spb "github.com/cloneable/go-microservice-template/api/proto/server"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type EchoServer struct {
	logger *zap.SugaredLogger
	tracer trace.Tracer
}

func New(logger *zap.SugaredLogger, tracer trace.Tracer) *EchoServer {
	return &EchoServer{
		logger: logger,
		tracer: tracer,
	}
}

var _ spb.EchoServiceServer = (*EchoServer)(nil)

func (s *EchoServer) Echo(ctx context.Context, req *spb.EchoRequest) (*spb.EchoResponse, error) {
	var span trace.Span
	ctx, span = s.tracer.Start(ctx, "echo request")
	defer span.End()

	if req.Msg == "error" {
		err := fmt.Errorf("echo error")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, "request failed")
		return nil, err
	}

	span.AddEvent("echo called!")

	s.logger.Infof("Message received: %v", req.Msg)
	return &spb.EchoResponse{Msg: req.Msg}, nil
}
