package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Options struct {
	GRPCPort         int
	RESTPort         int
	MonitoringPort   int
	RegisterServices ServiceRegistrationCallback
	GatewayServices  []GatewayRegistration
}

type ServiceRegistrationCallback func(s grpc.ServiceRegistrar)

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
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", opt.GRPCPort))
	if err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			serverMetrics.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(s.logger.Desugar()),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			serverMetrics.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(s.logger.Desugar()),
		)),
	)

	if opt.RegisterServices != nil {
		opt.RegisterServices(grpcServer)
	}
	serverMetrics.InitializeMetrics(grpcServer)

	monitoringServer := &http.Server{
		Handler: promhttp.HandlerFor(metricsRegistry, promhttp.HandlerOpts{Registry: metricsRegistry}),
		Addr:    fmt.Sprintf("0.0.0.0:%d", opt.MonitoringPort),
	}
	go func() {
		if err := monitoringServer.ListenAndServe(); err != nil {
			s.logger.Fatal("Unable to start a monitoring http server.")
		}
	}()

	go func() {
		s.logger.Fatal(grpcServer.Serve(grpcListener))
	}()

	conn, err := grpc.DialContext(
		ctx,
		grpcListener.Addr().String(),
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
	for _, regFunc := range opt.GatewayServices {
		if err := regFunc(ctx, gateway, conn); err != nil {
			return fmt.Errorf("failed to register service with gateway: %w", err)
		}
	}

	gatewayServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", opt.RESTPort),
		Handler: gateway,
	}

	s.logger.Info("Server running.")
	s.logger.Fatal(gatewayServer.ListenAndServe())

	return nil
}

func (s *Service) Logger() *zap.SugaredLogger {
	return s.logger
}
