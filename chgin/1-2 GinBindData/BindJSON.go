package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	e:=gin.Default()

	e.POST("/stu", func(context *gin.Context) {
		var stu Stu

		if err:=context.BindJSON(&stu);err!=nil{
			log.Fatal(err.Error())
			return
		}
		fmt.Println(stu.Name + "学号" + stu.Number)
		context.Writer.Write([]byte("姓名"+stu.Name))

	})

	e.Run()
}

type Stu struct {
	Name string `form:"name"`
	Number string `form:"number"`
}