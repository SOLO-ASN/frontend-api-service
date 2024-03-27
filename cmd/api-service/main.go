package main

import (
	"api-service/internal/server"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

const defaultHTTPAddr = "0.0.0.0:18080"

func main() {
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
