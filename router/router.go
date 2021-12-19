package router

import (
	"chat/api"
	"chat/conf"
	"chat/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	conf.Init()
	r := gin.Default()
	r.Use(gin.Recovery(),gin.Logger())
	v1 := r.Group("/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200,"SUCCESS")
		})
		v1.POST("user/register", api.UserRegister)
		v1.GET("ws",service.WsHandler)
	}
	return r
}
