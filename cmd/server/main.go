package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"syscall"

	healthz_proto "github.com/cloneable/go-microservice-template/api/proto/healthz"
	server_proto "github.com/cloneable/go-microservice-template/api/proto/server"
	"github.com/cloneable/go-microservice-template/internal/handler/echoserver"
	"github.com/cloneable/go-microservice-template/internal/handler/healthz"
	"github.com/cloneable/go-microservice-template/pkg/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	_ "net/http/pprof"
)

var (
	restPort       = flag.Int("rest_port", 8080, "port of the REST server")
	grpcPort       = flag.Int("grpc_port", 0, "port of the gRPC server")
	monitoringPort = flag.Int("monitoring_port", 9090, "port of the monitoring metrics HTTP server")
	pprofPort      = flag.Int("pprof_port", 6060, "port of the pprof handler")
)

func main() {
	syscall.Umask(0077)

	ctx := context.Background()
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to create zap logger: %v", err)
	}
	undoRedirect, err := zap.RedirectStdLogAt(logger, zapcore.DebugLevel)
	if err != nil {
		log.Fatalf("Failed to redirect std/log output: %v", err)
	}
	defer undoRedirect()

	grpclog.SetLoggerV2(service.NewZapDepthLogger(logger))

	go func() {
		// pprof endpoint
		logger.Sugar().Info(http.ListenAndServe(fmt.Sprintf(":%d", *pprofPort), nil))
	}()

	tp, err := service.NewTracerProvider()
	if err != nil {
		logger.Sugar().Fatalf("Failed to create tracer provider: %v", err)
	}

	srv, err := service.New(logger)
	if err != nil {
		logger.Sugar().Fatalf("Failed to create server: %v", err)
	}

	err = srv.Run(ctx, service.Options{
		GRPCPort:       *grpcPort,
		RESTPort:       *restPort,
		MonitoringPort: *monitoringPort,
		RegisterServices: func(s grpc.ServiceRegistrar) {
			healthz_proto.RegisterHealthzServer(s, &healthz.HealthzServer{})
			server_proto.RegisterEchoServiceServer(s, echoserver.New(logger.Sugar(), tp.Tracer("echoserver")))
		},
		GatewayServices: []service.GatewayRegistration{
			healthz_proto.RegisterHealthzHandler,
			server_proto.RegisterEchoServiceHandler,
		},
	})

	if err != nil {
		logger.Sugar().Fatalf("Failed to start server: %v", err)
	}
}
