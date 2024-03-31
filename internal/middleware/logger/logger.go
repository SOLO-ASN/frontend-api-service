package logger

import (
	"github.com/natefinch/lumberjack"
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
	//new encoder config
	cfg := getZapEncoderConfig(formatJSON)
	// new encoder
	var encoder zapcore.Encoder
	if l.format == formatJSON {
		encoder = zapcore.NewJSONEncoder(cfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(cfg)
	}

	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename: fo.filename,
		Compress: fo.isCompression,
	})

	core := zapcore.NewCore(encoder, writeSyncer, getZapLevel(l.level))

	l.zLog = zap.New(core, zap.AddCaller())
}

func (l *logger) log2Console() {
	cfg := zap.Config{
		Level:         zap.NewAtomicLevelAt(getZapLevel(l.level)),
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}
	cfg.EncoderConfig = getZapEncoderConfig(l.format)
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

func getZapEncoderConfig(format string) zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder
	if format == formatConsole {
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	return cfg
}
