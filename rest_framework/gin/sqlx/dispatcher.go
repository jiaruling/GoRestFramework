package sqlx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
   功能说明: 分发器
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:30
*/

func Dispatcher(m ViewAPI, c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		if len(c.Param("id")) == 1 {
			m.ListViewAPI(c)
		} else {
			m.RetrieveViewAPI(c)
		}
	case "POST":
		m.CreateViewAPI(c)
	case "PUT":
		m.UpdateViewAPI(c)
	case "DELETE":
		m.DeleteViewAPI(c)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"code":http.StatusBadRequest,"msg":"请求方式不被允许","data":""})
	}
	return
}