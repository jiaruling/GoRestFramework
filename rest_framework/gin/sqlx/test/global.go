package main

import (
	"GoRestFramework/rest_framework/gin/sqlx"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:47
*/

func init() {
	database, err := sqlx.Open("mysql", "root:abc123456@tcp(127.0.0.1:3306)/imooc")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	//defer database.Close()  // 注意这行代码要写在上面err判断的下面
	grf.RDB = database
	grf.WDB = database
	grf.GlobalPageMax = 5
	grf.GlobalPageMin = 1
}


