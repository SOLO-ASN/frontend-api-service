package model

import (
	"api-service/config"
	"api-service/internal/dbEntity/mysql"
	"api-service/internal/middleware/logger"
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

var (
	ErrGenerateMysql = errors.New("generate mysql error")
	ErrGetMysql      = errors.New("get mysql error")
)

func InitMysql() {
	opts := []mysql.OptionFn{
		mysql.WithMaxIdleConns(10),
		mysql.WithMaxOpenConns(100),
		mysql.WithMaxLifetime(time.Hour),
	}

	if config.Get().Mysql.EnableLog {
		opts = append(opts,
			mysql.WithLogEnable(),
			mysql.WithLogger(logger.DefaultLogger()),
		)
	}

	var err error
	db, err = mysql.Init(config.Get().Mysql.Dsn, opts...)
	if err != nil {
		panic("generate mysql error: " + err.Error())
	}
}

func GetDb() (*gorm.DB, error) {
	if db == nil {
		return nil, ErrGetMysql
	}
	return db, nil
}
