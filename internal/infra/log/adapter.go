package log

import (
	"github.com/bytopia/kratos-ddd-template/pkg/logging"
	"github.com/go-kratos/kratos/v2/log"
)

func NewLoggerAdapter(logger log.Logger) logging.LoggerAdapter {
	return log.NewHelper(logger)
}
