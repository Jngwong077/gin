package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Use(RequestInfos())
	//query解析
	e.GET("/query", Query)
	e.Run()
}

func Query(context *gin.Context) {
	context.JSON(404, map[string]interface{}{
		"code": 1,
		"msg":  context.FullPath(),
	})
}

//打印请求的中间件
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method

		fmt.Println("请求Path" + path)
		fmt.Println("请求method" + method)
		fmt.Println("请求前status", context.Writer.Status())

		//调用该路由前的其他其他中间件去执行
		//然后最后在执行该路由，可以做一些请求路由前的日志打印，参数验证
		context.Next()
		fmt.Println("请求后status", context.Writer.Status())

	}
}


