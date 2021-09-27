package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	healthz_proto "github.com/cloneable/go-microservice-template/api/proto/healthz"
	server_proto "github.com/cloneable/go-microservice-template/api/proto/server"
	"github.com/cloneable/go-microservice-template/pkg/handler/echoserver"
	"github.com/cloneable/go-microservice-template/pkg/handler/healthz"
	"github.com/cloneable/go-microservice-template/pkg/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "net/http/pprof"
)

var (
	restPort       = flag.Int("rest_port", 8080, "port of the REST server")
	grpcPort       = flag.Int("grpc_port", 0, "port of the gRPC server")
	monitoringPort = flag.Int("monitoring_port", 9090, "port of the monitoring metrics HTTP server")
	pprofPort      = flag.Int("pprof_port", 6060, "port of the pprof handler")
)

func main() {
	ctx := context.Background()
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to create zap logger: %v", err)
	}
	undoRedirect := zap.RedirectStdLog(logger)
	defer undoRedirect()

	go func() {
		// pprof endpoint
		logger.Sugar().Info(http.ListenAndServe(fmt.Sprintf(":%d", *pprofPort), nil))
	}()

	srv, err := server.New(logger)
	if err != nil {
		logger.Sugar().Fatalf("Failed to create server: %v", err)
	}

	err = srv.Run(ctx, server.Options{
		GRPCPort:       *grpcPort,
		RESTPort:       *restPort,
		MonitoringPort: *monitoringPort,
		RegisterServices: func(s grpc.ServiceRegistrar) {
			healthz_proto.RegisterHealthzServer(s, &healthz.HealthzServer{})
			server_proto.RegisterEchoServiceServer(s, &echoserver.EchoServer{Logger: logger.Sugar()})
		},
		GatewayServices: []server.GatewayRegistration{
			healthz_proto.RegisterHealthzHandler,
			server_proto.RegisterEchoServiceHandler,
		},
	})

	if err != nil {
		logger.Sugar().Fatalf("Failed to start server: %v", err)
	}
}
