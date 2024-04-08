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
	// get default mysql options
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
		logger.DefaultLogger().Error("Init mysql error: " + err.Error())
	}

	migrateErr := db.AutoMigrate(&User{})
	if migrateErr != nil {
		logger.DefaultLogger().Error("migrate mysql error: " + migrateErr.Error())
	}
}

func GetDb(init bool) *gorm.DB {
	if db == nil && init {
		InitMysql()
	}
	return db
}
