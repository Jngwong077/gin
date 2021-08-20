package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	e := gin.Default()

	e.GET("/write", func(context *gin.Context) {
		context.Writer.WriteString("Lujing"+context.FullPath())
	})



	e.Run()
}
