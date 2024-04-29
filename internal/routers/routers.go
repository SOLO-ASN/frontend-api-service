package routers

import (
	"api-service/config"
	"api-service/internal/handler"
	"api-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

var (
	base_path = "/api"

	group_UserId      []func(rg *gin.RouterGroup)
	group_UserId_Path = "/user"

	group_Space      []func(rg *gin.RouterGroup)
	group_Space_Path = "/space"

	group_Spaces      []func(rg *gin.RouterGroup)
	group_Spaces_Path = "/spaces"

	group_Explore      []func(rg *gin.RouterGroup)
	group_Explore_Path = "/explore"

	group_Campaign      []func(rg *gin.RouterGroup)
	group_Campaign_Path = "/campaign"

	group_Campaigns      []func(rg *gin.RouterGroup)
	group_Campaigns_Path = "/campaigns"

	group_Images      []func(rg *gin.RouterGroup)
	group_Images_Path = "/images"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// use cors middleware
	r.Use(middleware.Cors())

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
	regRouters(r, base_path+group_Spaces_Path, group_Spaces)
	regRouters(r, base_path+group_Explore_Path, group_Explore)
	regRouters(r, base_path+group_Campaign_Path, group_Campaign)
	regRouters(r, base_path+group_Campaigns_Path, group_Campaigns)

	return r
}

func regRouters(r *gin.Engine, groupPath string, routerFns []func(rg *gin.RouterGroup), handlers ...gin.HandlerFunc) {
	rg := r.Group(groupPath, handlers...)
	for _, fn := range routerFns {
		fn(rg)
	}
}
