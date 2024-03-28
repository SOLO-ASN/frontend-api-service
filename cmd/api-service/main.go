package main

import (
	"os"
	"os/signal"
	"syscall"

	"api-service/config"
	"api-service/internal/server"
	"github.com/gin-gonic/gin"
)

const defaultHTTPAddr = "0.0.0.0:18080"

func main() {

	path := os.Getenv("CONFIG_PATH")
	config.Init(path)

	server := server.NewHTTPServer(defaultHTTPAddr,
		server.WithMode(gin.DebugMode))
	server.Start()

	var quit chan os.Signal
	quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP)
	<-quit

	server.Stop()
	os.Exit(0)
}
