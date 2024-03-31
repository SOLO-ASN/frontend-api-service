package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap/zapcore"
	"log"
	"os"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	dmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var (
	ErrCreateMysql        = errors.New("create mysql error")
	ErrCreateGormInstance = errors.New("create gorm instance error")
	ErrUsingPlugin        = errors.New("error using otelgorm plugin")
)

type Mysql struct {
}

func initMysql(dns string, opts ...OptionFn) (*gorm.DB, error) {
	// get default mysql options
	opt := defaultOption()
	opt.apply(opts...)

	//
	sql, err := sql.Open("mysql-driver", dns)
	if err != nil {
		return nil, ErrCreateMysql
	}
	sql.SetMaxIdleConns(opt.maxIdleConns)
	sql.SetMaxOpenConns(opt.maxOpenConns)
	sql.SetConnMaxLifetime(opt.conMaxLifetime)

	// gorm
	db, err := gorm.Open(dmysql.New(dmysql.Config{Conn: sql}), gormConfig(opt))
	if err != nil {
		return nil, ErrCreateGormInstance
	}
	// automatic appending of table suffixes when creating tables
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	//
	if opt.enableTrace {
		err = db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("solo-mission-db")))
		if err != nil {
			return nil, ErrUsingPlugin
		}
	}

	// todo add read-write separation
	if len(opt.slaveDsn) > 0 {

	}

	// reg user's plugins
	for _, p := range opt.plugins {
		db.Use(p)
	}
	return db, nil

}

func gormConfig(o *option) *gorm.Config {
	cfg := &gorm.Config{}
	if o.enableLog {
		if o.slowThreshold > 0 {
			cfg.Logger = gormlogger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				gormlogger.Config{
					SlowThreshold: o.slowThreshold,
					Colorful:      true,
					LogLevel:      gormlogger.Info,
				})
		} else if o.logger == nil {
			cfg.Logger = gormlogger.Default.LogMode(zapLogLevel2GormLogLevel(o.logLevel))
		} else {
			cfg.Logger = InitGormLogger(o)
		}
	}
	return cfg
}

// zapLogLevel2GormLogLevel convert zapcore.Level to gormlogger.LogLevel
// this sucks
func zapLogLevel2GormLogLevel(level zapcore.Level) gormlogger.LogLevel {
	switch level {
	case zapcore.InfoLevel:
		return gormlogger.Info
	case zapcore.WarnLevel:
		return gormlogger.Warn
	case zapcore.ErrorLevel:
		return gormlogger.Error
	default:
		return gormlogger.Info
	}
}
