package zaplog

import (
	confpb "github.com/bytopia/kratos-ddd-layout/internal/pkg/proto/conf"
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/url"
	"strings"
)

const (
	prefixLumberjack  = "lumberjack"
	outputFile        = "file"
	outputConsole     = "console"
	defaultMaxSizeMB  = 100
	defaultMaxAgeDays = 24
	defaultMaxBackups = 50
	defaultCompress   = true
)

func NewLogger(c *confpb.Logging, devMode bool, flagLevel string) log.Logger {
	output := outputConsole
	if c.Output == outputFile {
		output = outputFile
	}
	if output == "file" && (strings.HasPrefix(c.DefaultFile, prefixLumberjack) || strings.HasPrefix(c.ErrorFile, prefixLumberjack)) {
		initLumberjack(c.Lumberjack)
	}

	config := zap.NewProductionConfig()
	if devMode {
		config = zap.NewDevelopmentConfig()
	}

	options := []zap.Option{
		zap.AddStacktrace(zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	}
	// 开启开发模式
	if devMode {
		options = append(options, zap.Development())
	}

	// 处理 level
	level := zapcore.DebugLevel
	_level := c.Level
	if flagLevel != "" {
		_level = flagLevel
	}

	switch strings.ToLower(_level) {
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	}
	config.Level = zap.NewAtomicLevelAt(level)

	// 处理文件
	if output == outputFile && c.DefaultFile != "" {
		config.OutputPaths = append(config.OutputPaths, c.DefaultFile)
	}

	if output == outputFile && c.ErrorFile != "" {
		config.ErrorOutputPaths = append(config.ErrorOutputPaths, c.ErrorFile)
	}

	zlogger, err := config.Build(options...)
	if err != nil {
		panic(err)
	}
	logger := kzap.NewLogger(zlogger)
	// 设置默认 logger
	log.SetLogger(logger)
	return logger
}

func newLumberjackLogger(c *confpb.Logging_Lumberjack, path string) *lumberjack.Logger {
	maxSize := defaultMaxSizeMB
	if c != nil && c.MaxSizeMb != 0 {
		maxSize = int(c.MaxSizeMb)
	}

	maxAge := defaultMaxAgeDays
	if c != nil && c.MaxAgeDays != 0 {
		maxAge = int(c.MaxAgeDays)
	}

	maxBackups := defaultMaxBackups
	if c != nil && c.MaxBackups != 0 {
		maxBackups = int(c.MaxBackups)
	}
	compress := defaultCompress
	if c != nil {
		compress = c.Compress
	}
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  true,
		Compress:   compress,
	}

}

func initLumberjack(c *confpb.Logging_Lumberjack) {
	zap.RegisterSink(prefixLumberjack, func(url *url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: newLumberjackLogger(c, url.Opaque),
		}, nil
	})
}
