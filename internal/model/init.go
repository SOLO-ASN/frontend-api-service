package model

import (
	"api-service/config"
	"api-service/internal/dbEntity/cache"
	"api-service/internal/dbEntity/mysql"
	"api-service/internal/middleware/logger"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	cacheDb *cache.Client
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

func InitRedis() {
	cacheOpts := []cache.OptionFn{
		cache.WithDialTimeout(30),
		cache.WithReadTimeout(10),
		cache.WithWriteTimeout(10),
	}

	cacheDb = cache.NewRedisClientWithPassword(
		config.Get().Redis.AddressList[0],
		config.Get().Redis.Password,
		config.Get().Redis.DB,
		cacheOpts...,
	)
}

func GetCacheDb() *cache.Client {
	if cacheDb == nil {
		InitRedis()
	}
	return cacheDb
}
