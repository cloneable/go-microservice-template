package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	serverpb "github.com/cloneable/go-microservice-template/api/proto/server"
	"github.com/cloneable/go-microservice-template/pkg/echoserver"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	restPort = flag.Int("rest_port", 9090, "port of the REST server")
	grpcPort = flag.Int("grpc_port", 8080, "port of the gRPC server")
)

func main() {
	ctx := context.Background()

	flag.Parse()
	defer glog.Flush()

	if err := run(ctx); err != nil {
		glog.Fatal(err)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}
	s := grpc.NewServer()
	serverpb.RegisterEchoServiceServer(s, &echoserver.EchoServer{})
	go func() {
		glog.Fatal(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("127.0.0.1:%d", *grpcPort),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial grpc server: %w", err)
	}

	gateway := runtime.NewServeMux()
	err = serverpb.RegisterEchoServiceHandler(ctx, gateway, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	gatewayServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *restPort),
		Handler: gateway,
	}

	glog.Fatal(gatewayServer.ListenAndServe())

	return nil
}
