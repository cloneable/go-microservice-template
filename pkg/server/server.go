package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	healthzpb "github.com/cloneable/go-microservice-template/api/proto/healthz"
	serverpb "github.com/cloneable/go-microservice-template/api/proto/server"
	"github.com/cloneable/go-microservice-template/pkg/echoserver"
	"github.com/golang/glog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Options struct {
	GrpcPort       int
	RestPort       int
	MonitoringPort int
}

func Run(ctx context.Context, opt Options) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opt.GrpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			serverMetrics.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			serverMetrics.StreamServerInterceptor(),
		)),
	)
	healthzpb.RegisterHealthzServer(s, &HealthzServer{})
	serverpb.RegisterEchoServiceServer(s, &echoserver.EchoServer{})

	grpc_prometheus.Register(s)
	monitoringServer := &http.Server{Handler: promhttp.HandlerFor(metricsRegistry, promhttp.HandlerOpts{Registry: metricsRegistry}), Addr: fmt.Sprintf("0.0.0.0:%d", opt.MonitoringPort)}
	go func() {
		if err := monitoringServer.ListenAndServe(); err != nil {
			glog.Fatal("Unable to start a monitoring http server.")
		}
	}()

	go func() {
		glog.Fatal(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		ctx,
		lis.Addr().String(),
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			clientMetrics.UnaryClientInterceptor(),
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			clientMetrics.StreamClientInterceptor(),
		)),
	)
	if err != nil {
		return fmt.Errorf("failed to dial grpc server: %w", err)
	}

	gateway := runtime.NewServeMux()
	err = healthzpb.RegisterHealthzHandler(ctx, gateway, conn)
	if err != nil {
		return fmt.Errorf("failed to register healthz handler with gateway: %w", err)
	}
	err = serverpb.RegisterEchoServiceHandler(ctx, gateway, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	gatewayServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", opt.RestPort),
		Handler: gateway,
	}

	glog.Info("Server running.")
	glog.Fatal(gatewayServer.ListenAndServe())

	return nil
}

var OK_BODY = []byte("ok\r\n")

type HealthzServer struct{}

func (s *HealthzServer) Check(ctx context.Context, _ *emptypb.Empty) (*httpbodypb.HttpBody, error) {
	return &httpbodypb.HttpBody{ContentType: "text/plain;charset=utf-8", Data: OK_BODY}, nil
}
