package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	e := gin.Default()

	e.GET("/mapjson", func(context *gin.Context) {
		fullpath := context.FullPath()
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"name": fullpath,
		})

	})

	e.GET("structjson", func(context *gin.Context) {
		res := Res{1, "hello", "wow"}
		context.JSON(http.StatusAccepted, &res)

	})

	e.Run()
}

type Res struct {
	Code    int
	Name    string
	Message string
}
