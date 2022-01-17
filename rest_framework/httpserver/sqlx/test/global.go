package main

import (
	sqlx2 "GoRestFramework/rest_framework/httpserver/sqlx"
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
	database, err := sqlx.Open("mysql", "root:abc123456@tcp(192.168.18.100:3306)/imooc")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	//defer database.Close()  // 注意这行代码要写在上面err判断的下面
	sqlx2.RDB = database
	sqlx2.WDB = database
	sqlx2.GlobalPageMax = 5
	sqlx2.GlobalPageMin = 1
}


