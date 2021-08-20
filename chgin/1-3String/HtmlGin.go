package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	e := gin.Default()
	//设置html路径
	e.LoadHTMLGlob("../html/*")

	//设置金静态资源
	e.Static("/img","../img")

	e.GET("htmlgin", func(context *gin.Context) {

		context.HTML(http.StatusOK,"home.html",map[string]interface{}{
			"title":"htmlgin",
			"h1":"dabiaoti",
		})


	})


	e.Run()
}
