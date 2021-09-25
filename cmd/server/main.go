package main

import (
	"context"
	"flag"

	"github.com/cloneable/go-microservice-template/pkg/server"
	"github.com/golang/glog"
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

	if err := server.Run(ctx, server.Options{
		GrpcPort:       *grpcPort,
		RestPort:       *restPort,
		MonitoringPort: *monitoringPort,
	}); err != nil {
		glog.Fatal(err)
	}
}
