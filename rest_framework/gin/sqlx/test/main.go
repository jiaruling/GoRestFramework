package main

import (
	"GoRestFramework/rest_framework/gin/sqlx"
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:08
*/

type name struct {
	N string
	A interface{}
}

var n = &name{N: "张三", A: Student{}}

func main() {
	r := gin.Default()
	r.GET("/ping/*id", func(c *gin.Context) {
		id := c.Param("id")[1:]
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    id,
		})
	})
	r.Any("/stu/*id", func(c *gin.Context) {
		//stu.Lock.Lock()
		//defer stu.Lock.Unlock()
		//stu.M = new(Student)
		fmt.Printf("%p\n", &stu)
		//stu.M = new(Student)
		s := stu
		//ss := Student{Id: 1}
		//s.M = ss
		fmt.Printf("stu: %p -- s:%p \n", &stu, &s)
		fmt.Printf("stu.M: %p -- s.M:%p \n", &stu.M, &s.M)
		fmt.Printf("%v -- %v \n", s.M, stu.M)
		fmt.Printf("---------------------------------------------\n")
		m := n
		fmt.Printf("n: %v, m: %v\n", n, m)
		fmt.Printf("n: %p, m: %p; n.N: %p, m.N: %p; n.A: %p, m.A: %p\n", &n, &m, &n.N, &m.N, &n.A, &m.A)
		n.N = "李四"
		n.A = Student{Id: 2}
		fmt.Printf("n: %v, m: %v\n", n, m)
		fmt.Printf("n: %p, m: %p; n.N: %p, m.N: %p; n.A: %p, m.A: %p\n", &n, &m, &n.N, &m.N, &n.A, &m.A)
		fmt.Printf("---------------------------------------------\n")
		sqlx.Dispatcher(stu, c)
		//c.JSON(http.StatusOK, gin.H{})
		return
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
