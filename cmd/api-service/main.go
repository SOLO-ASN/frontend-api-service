package main

import (
	"api-service/config"
	"api-service/internal/server"
	"flag"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

const defaultHTTPAddr = "127.0.0.1:18080"

func main() {

	config.Init(getConfigPath())

	host := defaultHTTPAddr
	if cfg := config.Get(); cfg.App.Host != "" {
		host = cfg.App.Host
	}

	server := server.NewHTTPServer(host,
		server.WithMode(gin.DebugMode))
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
