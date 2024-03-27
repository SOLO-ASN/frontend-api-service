package server

import (
	"api-service/internal/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ IServer = (*httpServer)(nil)

type httpServer struct {
	addr   string
	server *http.Server
}

func NewHTTPServer(addr string, opts ...HTTPOptionFunc) IServer {
	// generate default option, then apply options
	opt := defaultHTTPOptions()
	opt.applyHTTPOptions(opts)

	// create http server
	gin.SetMode(opt.mode)

	handler := routers.NewRouter()

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  opt.readTimeout,
		WriteTimeout: opt.writeTimeout,
		Handler:      handler,
	}

	return &httpServer{
		addr:   addr,
		server: server,
	}
}

func (h *httpServer) Start() error {
	//TODO implement me
	if err := h.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}
	return nil
}

func (h *httpServer) Stop() error {
	//TODO implement me
	fmt.Println("stop http server")
	return nil
}

func (h *httpServer) String() string {
	//TODO implement me
	panic("implement me")
}

func (h *httpServer) IsRunning() bool {
	//TODO implement me
	panic("implement me")
}

func (h *httpServer) GetPort() int {
	//TODO implement me
	panic("implement me")
}

func (h *httpServer) GetAddress() string {
	//TODO implement me
	panic("implement me")
}
