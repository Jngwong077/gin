package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	e := gin.Default()

	e.POST("/tea", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var tea Tea
		if err := context.ShouldBind(&tea); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(tea.Name + "电话" + tea.Tel)

		context.Writer.Write([]byte("姓名" + tea.Name))

	})

	e.Run()
}

type Tea struct {
	Name string `form:"name"`
	Tel  string `form:"tel"`
}
