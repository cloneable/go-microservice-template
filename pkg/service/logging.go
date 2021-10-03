package service

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc/grpclog"
)

type zapDepthLogger struct {
	zapgrpc.Logger

	logger *zap.Logger
}

func newZapDepthLogger(logger *zap.Logger) *zapDepthLogger {
	return &zapDepthLogger{
		Logger: *zapgrpc.NewLogger(logger),
		logger: logger,
	}
}

var _ grpclog.DepthLoggerV2 = (*zapDepthLogger)(nil)

func (l *zapDepthLogger) InfoDepth(depth int, args ...interface{}) {
	l.logger.WithOptions(zap.AddCallerSkip(depth + 2)).Sugar().Info(args...)
}

func (l *zapDepthLogger) WarningDepth(depth int, args ...interface{}) {
	l.logger.WithOptions(zap.AddCallerSkip(depth + 2)).Sugar().Warn(args...)
}

func (l *zapDepthLogger) ErrorDepth(depth int, args ...interface{}) {
	l.logger.WithOptions(zap.AddCallerSkip(depth + 2)).Sugar().Error(args...)
}

func (l *zapDepthLogger) FatalDepth(depth int, args ...interface{}) {
	l.logger.WithOptions(zap.AddCallerSkip(depth + 2)).Sugar().Fatal(args...)
}
