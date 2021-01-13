package router

import (
	"vss/app"
	"vss/app/api/finance"
	"vss/app/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	engine := app.GetEngine()
	//注册中间件
	engine.Use(gin.Logger())         //日志
	engine.Use(middleware.Recover()) //异常捕获
	//api相关路由
	apiRouter := engine.Group("/api")
	{
		financeRouter := apiRouter.Group("/finance")
		{
			financeRouter.POST("/getMergeList", finance.GetMergeList)
		}
	}

}
