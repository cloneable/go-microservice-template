package main

import (
	"context"
	"flag"

	healthz_proto "github.com/cloneable/go-microservice-template/api/proto/healthz"
	server_proto "github.com/cloneable/go-microservice-template/api/proto/server"
	"github.com/cloneable/go-microservice-template/pkg/handler/echoserver"
	"github.com/cloneable/go-microservice-template/pkg/handler/healthz"
	"github.com/cloneable/go-microservice-template/pkg/server"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	restPort       = flag.Int("rest_port", 8080, "port of the REST server")
	grpcPort       = flag.Int("grpc_port", 0, "port of the gRPC server")
	monitoringPort = flag.Int("monitoring_port", 9090, "port of the monitoring metrics HTTP server")
)

func main() {
	ctx := context.Background()

	flag.Parse()
	defer glog.Flush()
	glog.Info("Server starting up...")

	err := server.Run(ctx, server.Options{
		GRPCPort:       *grpcPort,
		RestPort:       *restPort,
		MonitoringPort: *monitoringPort,
	}, func(s grpc.ServiceRegistrar) {
		healthz_proto.RegisterHealthzServer(s, &healthz.HealthzServer{})
		server_proto.RegisterEchoServiceServer(s, &echoserver.EchoServer{})
	}, []server.GatewayRegistration{
		healthz_proto.RegisterHealthzHandler, server_proto.RegisterEchoServiceHandler,
	})

	if err != nil {
		glog.Fatal(err)
	}
}
