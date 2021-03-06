package service

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

var (
	metricsRegistry = prometheus.NewRegistry()
	serverMetrics   = grpc_prometheus.NewServerMetrics()
	clientMetrics   = grpc_prometheus.NewClientMetrics()
)

func init() {
	metricsRegistry.MustRegister(collectors.NewGoCollector())
	metricsRegistry.MustRegister(collectors.NewBuildInfoCollector())
	metricsRegistry.MustRegister(serverMetrics)
	metricsRegistry.MustRegister(clientMetrics)
}
