package mysql

import (
	"context"
	"go.uber.org/zap"
	logger "gorm.io/gorm/logger"
	"time"
)

var (
	requestIDKey = "request_id"
)

type gormLogger struct {
	logger       *zap.Logger
	level        logger.LogLevel
	requestIdKey string
}

func (l gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func (l gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Info {
		return
	}
	l.logger.Info(msg, zap.Any("data", data), l.requestIdFromContext(ctx))
}

func (l gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Warn {
		return
	}
	l.logger.Warn(msg, zap.Any("data", data), l.requestIdFromContext(ctx))
}

func (l gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Error {
		return
	}
	l.logger.Error(msg, zap.Any("data", data), l.requestIdFromContext(ctx))
}

func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//TODO implement me
	panic("implement me")
}

func InitGormLogger(o *option) logger.Interface {
	return &gormLogger{
		logger:       o.logger,
		level:        zapLogLevel2GormLogLevel(o.logLevel),
		requestIdKey: requestIDKey,
	}
}

func (l *gormLogger) requestIdFromContext(ctx context.Context) zap.Field {
	if l.requestIdKey == "" {
		return zap.Skip()
	}
	field := zap.Skip()
	if v := ctx.Value(l.requestIdKey); v != nil {
		field = zap.String(l.requestIdKey, v.(string))
	}
	return field
}
