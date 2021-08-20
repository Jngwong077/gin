package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	e := gin.Default()

	e.GET("/user", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var user User

		err := context.ShouldBindQuery(&user)


		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(user.Username)
		context.Writer.Write([]byte(user.Username +"你好"))
	})

	e.Run()
}

type User struct {
	Username string `form:"name"`
}
