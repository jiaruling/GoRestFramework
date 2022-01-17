package main

import (
	"GoRestFramework/rest_framework/httpserver/sqlx"
	"fmt"
	"net/http"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:08
*/

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		sqlx.Handler200(w, "pong")
		return
	})
	http.HandleFunc("/stu", func(w http.ResponseWriter, r *http.Request) {
		// 通过全局变量赋值给局部变量，可以实现并发
		s := stu
		s.M = new(Student)
		sqlx.Dispatcher(s, w, r)
		fmt.Printf("stu: %p -- s:%p \n", &stu, &s)
		fmt.Printf("stu.M: %p -- s.M:%p \n", &stu.M, &s.M)
		fmt.Printf("stu.T: %p -- s.T:%p \n", &stu.Table, &s.Table)
		fmt.Printf("stu.M: %v -- s.M:%v \n", stu.M, s.M)
		fmt.Printf("-----------------------------------------------------\n")
		return
	})
	http.ListenAndServe(":8001", nil)
}


















