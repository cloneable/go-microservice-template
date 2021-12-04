package healthz

import (
	"context"

	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/types/known/emptypb"
)

var okBody = []byte("ok\r\n")

type HealthzServer struct{}

func (s *HealthzServer) Alive(ctx context.Context, _ *emptypb.Empty) (*httpbodypb.HttpBody, error) {
	return &httpbodypb.HttpBody{ContentType: "text/plain; charset=UTF-8", Data: okBody}, nil
}

func (s *HealthzServer) Ready(ctx context.Context, _ *emptypb.Empty) (*httpbodypb.HttpBody, error) {
	return &httpbodypb.HttpBody{ContentType: "text/plain; charset=UTF-8", Data: okBody}, nil
}
