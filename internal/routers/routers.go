package routers

import (
	"api-service/config"

	"api-service/internal/handler"
	"github.com/gin-gonic/gin"
)

var (
	base_path = "/api"

	group_UserId      []func(rg *gin.RouterGroup)
	group_UserId_Path = "/user"

	group_Space      []func(rg *gin.RouterGroup)
	group_Space_Path = "/space"

	group_Explore      []func(rg *gin.RouterGroup)
	group_Explore_Path = "/explore"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET(base_path+"/health", handler.CheckHealth)

	// get config
	cfg := config.Get()
	if cfg.MiddleWare.EnableTrace {
		// todo implement me
		panic("implement me")
	}

	if cfg.MiddleWare.EnableCircuitBreaker {
		// todo implement me
		panic("implement me")
	}

	if cfg.MiddleWare.EnableRateLimit {
		// todo implement me
		panic("implement me")
	}

	if cfg.MiddleWare.EnableMetrics {
		// todo implement me
		panic("implement me")
	}

	regRouters(r, base_path+group_UserId_Path, group_UserId)
	regRouters(r, base_path+group_Space_Path, group_Space)
	regRouters(r, base_path+group_Explore_Path, group_Explore)

	return r
}

func regRouters(r *gin.Engine, groupPath string, routerFns []func(rg *gin.RouterGroup), handlers ...gin.HandlerFunc) {
	rg := r.Group(groupPath, handlers...)
	for _, fn := range routerFns {
		fn(rg)
	}
}
