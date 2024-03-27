package server

type IServer interface {
	Start() error
	Stop() error
	String() string
	IsRunning() bool
	GetPort() int
	GetAddress() string
}
