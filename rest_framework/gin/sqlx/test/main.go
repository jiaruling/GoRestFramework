package test

import (
	"GoRestFramework/rest_framework/gin/sqlx"
	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:08
*/

func main() {
	r := gin.Default()
	r.GET("/ping/*id", func(c *gin.Context) {
		id := c.Param("id")[1:]
		c.JSON(200, gin.H{
			"message": "pong",
			"data": id,
		})
	})
	r.Any("/stu/*id", func(c *gin.Context) {
		sqlx.Dispatcher(stu, c)
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

