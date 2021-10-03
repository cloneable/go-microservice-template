package service

import (
	"context"
	"fmt"
	"net/http"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"

	_ "net/http/pprof"
)

func Init(ctx context.Context, port int) (*zap.Logger, trace.TracerProvider, error) {
	syscall.Umask(0077)

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create zap logger: %w", err)
	}
	_, err = zap.RedirectStdLogAt(logger, zapcore.DebugLevel)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to redirect std/log output: %w", err)
	}

	grpclog.SetLoggerV2(newZapDepthLogger(logger))

	tp, err := newTracerProvider()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create tracer provider: %w", err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(metricsRegistry, promhttp.HandlerOpts{Registry: metricsRegistry}))

	go func() {
		logger.Sugar().Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	}()

	return logger, tp, nil
}
