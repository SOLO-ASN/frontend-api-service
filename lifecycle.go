package explorer_api_service

type LifeCycle interface {
	Start() error
	Stop() error
}
