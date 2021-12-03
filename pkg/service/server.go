package service

import (
	"context"
	"fmt"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Options struct {
	GRPCPort         int
	RESTPort         int
	RegisterServices ServiceRegistrationCallback
	GatewayServices  []GatewayRegistration
}

type ServiceRegistrationCallback func(s *grpc.Server)

type GatewayRegistration func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

type Service struct {
	logger *zap.SugaredLogger
}

func New(logger *zap.Logger) (*Service, error) {
	return &Service{
		logger: logger.Sugar(),
	}, nil
}

func (s *Service) Run(ctx context.Context, opt Options) error {
	s.logger.Info("Server starting.")
	// grpcListener, err := net.Listen("unix", "/tmp/gateway.sock")
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", opt.GRPCPort))
	if err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(),
			serverMetrics.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(s.logger.Desugar()),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			otelgrpc.StreamServerInterceptor(),
			serverMetrics.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(s.logger.Desugar()),
		)),
	)

	if opt.RegisterServices != nil {
		opt.RegisterServices(grpcServer)
	}

	serverMetrics.InitializeMetrics(grpcServer)

	go func() {
		s.logger.Fatal(grpcServer.Serve(grpcListener))
	}()

	conn, err := grpc.DialContext(
		ctx,
		"<not used>",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			var dialer net.Dialer
			return dialer.DialContext(ctx, grpcListener.Addr().Network(), grpcListener.Addr().String())
		}),
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			otelgrpc.UnaryClientInterceptor(),
			clientMetrics.UnaryClientInterceptor(),
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			otelgrpc.StreamClientInterceptor(),
			clientMetrics.StreamClientInterceptor(),
		)),
	)
	if err != nil {
		return fmt.Errorf("failed to dial grpc server: %w", err)
	}

	gateway := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		return metadata.Pairs("correlation-id", "moo")
	}))
	for _, regFunc := range opt.GatewayServices {
		if err := regFunc(ctx, gateway, conn); err != nil {
			return fmt.Errorf("failed to register service with gateway: %w", err)
		}
	}

	gatewayServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", opt.RESTPort),
		Handler: otelhttp.NewHandler(gateway, "gateway"),
	}

	s.logger.Info("Server running.")
	s.logger.Fatal(gatewayServer.ListenAndServe())

	return nil
}

func (s *Service) Logger() *zap.SugaredLogger {
	return s.logger
}
