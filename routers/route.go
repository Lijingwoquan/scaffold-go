package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var r *gin.Engine

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()
	r := gin.Default()
	//r.Use(cors.Default()) --> 这里没有Authorization！！！妈的被坑惨了
	// 创建新的CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	r.Static("/img", "/app/statics/img")

	r.NoRoute(func(c *gin.Context) {
		c.JSONP(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}
