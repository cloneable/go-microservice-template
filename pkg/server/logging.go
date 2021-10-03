package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc/grpclog"
)

type ZapDepthLogger struct {
	zapgrpc.Logger

	logger *zap.Logger
}

func NewZapDepthLogger(logger *zap.Logger) *ZapDepthLogger {
	return &ZapDepthLogger{
		Logger: *zapgrpc.NewLogger(logger),
		logger: logger,
	}
}

var _ grpclog.DepthLoggerV2 = (*ZapDepthLogger)(nil)

func (l *ZapDepthLogger) InfoDepth(depth int, args ...interface{}) {
	l.logger.WithOptions(zap.AddCallerSkip(depth + 2)).Sugar().Info(args...)
}

func (l *ZapDepthLogger) WarningDepth(depth int, args ...interface{}) {
	l.logger.WithOptions(zap.AddCallerSkip(depth + 2)).Sugar().Warn(args...)
}

func (l *ZapDepthLogger) ErrorDepth(depth int, args ...interface{}) {
	l.logger.WithOptions(zap.AddCallerSkip(depth + 2)).Sugar().Error(args...)
}

func (l *ZapDepthLogger) FatalDepth(depth int, args ...interface{}) {
	l.logger.WithOptions(zap.AddCallerSkip(depth + 2)).Sugar().Fatal(args...)
}
