package echoserver

import (
	"context"

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

	if err := req.Validate(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "validation failed")
		return nil, err
	}

	span.AddEvent("echo called!")

	s.logger.Infof("Message received: %v", req.Msg)
	return &spb.EchoResponse{Msg: req.Msg}, nil
}
