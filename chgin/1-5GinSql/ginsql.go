package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//连接数据库
	connstr := "root:123456@tcp(127.0.0.1:3306)/ginsql"


	//打开数据库
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建数据库表
	_, err = db.Exec("create table person(" +
		"id int auto_increment primary key," +
		"name varchar(12) not null" +
		");")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
