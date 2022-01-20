package grf

import (
	"net/http"
)

/*
   功能说明: 分发器
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:30
*/

func Dispatcher(m ViewAPI, w http.ResponseWriter, r *http.Request) {
	if !m.GetModelIsInit() {
		Handler500(w, "模型没有初始化, 不能进行操作")
		return
	}
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		if id == "" {
			m.ListViewAPI(w, r)
		} else {
			m.RetrieveViewAPI(w, r)
		}
	case "POST":
		m.CreateViewAPI(w, r)
	case "PUT":
		m.UpdateViewAPI(w, r)
	case "DELETE":
		m.DeleteViewAPI(w, r)
	default:
		Handler400(w, "请求方式不被允许", nil)
	}
	return
}
