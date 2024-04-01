package mysql

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type OptionFn func(*option)

type option struct {
	enableLog     bool
	enableTrace   bool
	slowThreshold time.Duration

	maxIdleConns   int
	maxOpenConns   int
	conMaxLifetime time.Duration

	logger        *zap.Logger
	logLevel      zapcore.Level
	enableLogFile bool
	logFilePath   string

	masterDsn []string
	slaveDsn  []string

	plugins []gorm.Plugin
}

func defaultOption() *option {
	return &option{
		enableLog:     false,            // whether to enable gorm log, default off
		enableTrace:   false,            //
		slowThreshold: time.Duration(0), // when greater than 0, pnly print logs longer than slowThreshold

		maxIdleConns:   3,                // set the maximum number of connections in the idle connection pool
		maxOpenConns:   50,               // set the maximum number of open database connections
		conMaxLifetime: 30 * time.Minute, // set the maximum time a connection can be reused

		logger:        nil,
		logLevel:      zap.InfoLevel,
		enableLogFile: false,
		logFilePath:   "logs.log",
	}
}

func (o *option) apply(opts ...OptionFn) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithLogEnable() OptionFn {
	return func(o *option) { o.enableLog = true }
}

func WithLogTrace() OptionFn {
	return func(o *option) { o.enableTrace = true }
}

func WithSlowThreshold(threshold time.Duration) OptionFn {
	return func(o *option) { o.slowThreshold = threshold }
}

func WithMaxIdleConns(conns int) OptionFn {
	return func(o *option) { o.maxIdleConns = conns }
}

func WithMaxOpenConns(conns int) OptionFn {
	return func(o *option) { o.maxOpenConns = conns }
}

func WithMaxLifetime(time time.Duration) OptionFn {
	return func(o *option) { o.conMaxLifetime = time }
}

func WithLogLevel(level zapcore.Level) OptionFn {
	return func(o *option) { o.logLevel = level }
}

func WithLogPath(path string) OptionFn {
	return func(o *option) { o.logFilePath = path }
}

func WithRWCluster(slaves []string, masters []string) OptionFn {
	return func(o *option) { o.slaveDsn = slaves; o.masterDsn = masters }
}

func WithPlugins(plugins ...gorm.Plugin) OptionFn {
	return func(o *option) { o.plugins = plugins }
}
