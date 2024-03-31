package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	formatJSON    = "json"
	formatConsole = "console"

	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
)

var defaultLogger *zap.Logger

func DefaultLogger() *zap.Logger {
	if defaultLogger == nil {
		logger := Init()
		defaultLogger = logger.zLog
	}
	return defaultLogger
}

type logger struct {
	format string
	level  string
	zLog   *zap.Logger
}

func Init(opts ...OptionFn) *logger {
	opt := defaultOption()
	opt.apply(opts...)

	logger := &logger{
		format: opt.encoding,
		level:  opt.level,
	}
	if opt.needToSave {
		logger.log2File(opt.fileOption)
	} else {
		logger.log2Console()
	}
	// set default logger
	defaultLogger = logger.zLog

	// return
	return logger
}

func (l *logger) log2File(fo *fileOption) {
	//l.zLog = zap.New()
}

func (l *logger) log2Console() {
	cfg := zap.Config{
		Level:         zap.NewAtomicLevelAt(getZapLevel(l.level)),
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}
	if l.format == formatConsole {
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zLog, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	l.zLog = zLog
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case levelDebug:
		return zapcore.DebugLevel
	case levelInfo:
		return zapcore.InfoLevel
	case levelWarn:
		return zapcore.WarnLevel
	case levelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
