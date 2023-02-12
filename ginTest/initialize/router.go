package initialize

import (
	"ginTest/middlewares"
	"ginTest/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 路由分组
	ApiGroup := Router.Group("/v1/")
	router.UserRouter(ApiGroup)
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	return Router
}
