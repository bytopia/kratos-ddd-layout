package stdlog

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"os"
)

func NewLogger(serviceId, serviceName, serviceVersion string) log.Logger {
	l := log.NewStdLogger(os.Stdout)
	return log.With(
		l,
		"service.id", serviceId,
		"service.name", serviceName,
		"service.version", serviceVersion,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}
