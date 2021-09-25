package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/golang/glog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func Run(ctx context.Context, opt Options) error {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", opt.GRPCPort))
	if err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			serverMetrics.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			serverMetrics.StreamServerInterceptor(),
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
			glog.Fatal("Unable to start a monitoring http server.")
		}
	}()

	go func() {
		glog.Fatal(grpcServer.Serve(grpcListener))
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

	glog.Info("Server running.")
	glog.Fatal(gatewayServer.ListenAndServe())

	return nil
}
