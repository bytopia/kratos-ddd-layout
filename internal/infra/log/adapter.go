package log

import (
	"github.com/bytopia/kratos-ddd-layout/pkg/logging"
	"github.com/go-kratos/kratos/v2/log"
)

type KratosLoggerAdapter struct {
	logger log.Logger
	helper *log.Helper
}

func (l *KratosLoggerAdapter) With(kv ...interface{}) logging.LoggerAdapter {
	logger := log.With(l.logger, kv)
	return NewLoggerAdapter(logger)
}

func (l *KratosLoggerAdapter) Debug(a ...interface{}) {
	l.helper.Debug(a)
}

func (l *KratosLoggerAdapter) Debugf(format string, a ...interface{}) {
	l.helper.Debugf(format, a)
}

func (l *KratosLoggerAdapter) Debugw(keyvals ...interface{}) {
	l.helper.Debugw(keyvals)
}

func (l *KratosLoggerAdapter) Info(a ...interface{}) {
	l.helper.Info(a)
}

func (l *KratosLoggerAdapter) Infof(format string, a ...interface{}) {
	l.helper.Infof(format, a)
}

func (l *KratosLoggerAdapter) Infow(keyvals ...interface{}) {
	l.helper.Infow(keyvals)
}

func (l *KratosLoggerAdapter) Warn(a ...interface{}) {
	l.helper.Warn(a)
}

func (l *KratosLoggerAdapter) Warnf(format string, a ...interface{}) {
	l.helper.Warnf(format, a)
}

func (l *KratosLoggerAdapter) Warnw(keyvals ...interface{}) {
	l.helper.Warnw(keyvals)
}

func (l *KratosLoggerAdapter) Error(a ...interface{}) {
	l.helper.Error(a)
}

func (l *KratosLoggerAdapter) Errorf(format string, a ...interface{}) {
	l.helper.Errorf(format, a)
}

func (l *KratosLoggerAdapter) Errorw(keyvals ...interface{}) {
	l.helper.Errorw(keyvals)
}

func (l *KratosLoggerAdapter) Fatal(a ...interface{}) {
	l.helper.Fatal(a)
}

func (l *KratosLoggerAdapter) Fatalf(format string, a ...interface{}) {
	l.helper.Fatalf(format, a)
}

func (l *KratosLoggerAdapter) Fatalw(keyvals ...interface{}) {
	l.helper.Fatalw(keyvals)
}

func NewLoggerAdapter(logger log.Logger) logging.LoggerAdapter {
	return &KratosLoggerAdapter{
		logger: logger,
		helper: log.NewHelper(logger),
	}
}
