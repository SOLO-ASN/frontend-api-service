package main

import (
	"api-service/config"
	"api-service/internal/middleware/logger"
	"api-service/internal/model"
	"api-service/internal/server"
	"flag"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

const defaultHTTPAddr = "127.0.0.1:18080"

func main() {

	config.Init(getConfigPath())
	cfg := config.Get()

	// init logger
	_ = logger.Init(
		logger.WithLevel(cfg.Logger.Level),
		logger.WithFormat(cfg.Logger.Format),
		logger.WithSaveToFile(cfg.Logger.IsSave),
	)


	// init mysql
	model.InitMysql()

	host := defaultHTTPAddr
	if cfg := config.Get(); cfg.App.Host != "" {
		host = cfg.App.Host
	}

	server := server.NewHTTPServer(host,
		server.WithMode(cfg.App.Env))
	server.Start()

	var quit chan os.Signal
	quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP)
	<-quit

	server.Stop()
	os.Exit(0)
}

func getConfigPath() string {
	path := ""
	path = os.Getenv("CONFIG_PATH")
	if path != "" {
		return path
	}

	// read "-c" flag
	flag.StringVar(&path, "c", "", "config file path")
	flag.Parse()

	return path
}
