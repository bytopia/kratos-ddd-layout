package logging

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Adapter log.Logger

type Logger struct {
	*log.Helper
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	return &Logger{
		Helper: l.Helper.WithContext(ctx),
	}
}

type Option log.Option

func NewLogger(adapter Adapter, opts ...Option) *Logger {
	var ops []log.Option
	for _, o := range opts {
		ops = append(ops, log.Option(o))
	}

	return &Logger{
		Helper: log.NewHelper(adapter, ops...),
	}
}
