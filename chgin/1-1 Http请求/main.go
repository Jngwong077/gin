package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//r.Handle("GET", "/hello", func(context *gin.Context) {
	//
	//	path := context.FullPath() //指上述的relativePath
	//	fmt.Println(path)
	//
	//	name := context.DefaultQuery("name", "") //传进来的参数
	//	fmt.Println(name)
	//	context.Writer.Write([]byte("name = " + name))
	//})
	//
	//r.Handle("POST", "/login", func(context *gin.Context) {
	//	username := context.PostForm("username")
	//	password := context.PostForm("password")
	//
	//	fmt.Println(username + password)
	//	context.Writer.Write([]byte(username + "登录"))
	//})

	r.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.Query("name")
		fmt.Println(name)

		context.Writer.Write([]byte("hello" + name))

	})

	r.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		//可判断是否存在表中
		username, exit := context.GetPostForm("username")
		if exit {
			fmt.Println(username)
		}
		password, exit := context.GetPostForm("password")
		if exit {
			fmt.Println(password)
		}

		context.Writer.Write([]byte("hello" + username + password))
	})

	r.DELETE("/user/:id", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		userid := context.Param("id")  //获取参数,路由变量
		fmt.Println(userid)
		context.Writer.Write([]byte("delete id :" + userid))

	})



	r.Run()
}
