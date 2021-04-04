package router

import (
	"github.com/anthonyzero/go-quick-api/controller"
	"github.com/anthonyzero/go-quick-api/docs"
	"github.com/anthonyzero/go-quick-api/router/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func LoadRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	docs.SwaggerInfo.Title = "go-quick-api"
	docs.SwaggerInfo.Description = "基于GO实现的快速开发接口框架，封装了常用的功能，使用简单，致力于进行快速的业务研发"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.NoCache) // 强制浏览器不使用缓存
	router.Use(middleware.Options) // 浏览器跨域 OPTIONS 请求设置
	router.Use(middleware.Secure)  // 一些安全设置
	router.Use(middlewares...)
	// 设置请求ID生成全局中间件
	router.Use(middleware.RequestId())

	router.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
		c.String(200, "%s", "success")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	indexRouter := router.Group("/index")
	indexRouter.Use() //...
	{
		controller.IndexRegister(indexRouter)
	}
	return router
}
