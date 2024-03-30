package logger

import (
	"context"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

var (
	formatJSON    = "json"
	formatConsole = "console"

	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
)

type logger struct {
}

func (l logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	//TODO implement me
	panic("implement me")
}

func (l logger) Info(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l logger) Warn(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l logger) Error(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//TODO implement me
	panic("implement me")
}

func Init() gormlogger.Interface {
	return &logger{}
}
