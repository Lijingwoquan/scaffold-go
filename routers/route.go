package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
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
	//静态文件
	r.Static("/img", "/app/statics/img")

	//下载功能
	r.GET("/download/:format/:filename", func(c *gin.Context) {
		format := c.Param("format")
		filename := c.Param("filename")
		filePath := filepath.Join("../scaffold-statics", format, filename)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")
		c.File(filePath)
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSONP(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}
