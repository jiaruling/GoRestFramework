package gorms

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

/*
   功能说明: 错误处理
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:56
*/

func Handler200(writer http.ResponseWriter, data interface{}) {
	handler(writer, http.StatusOK, "success", data)
	return
}

func Handler201(writer http.ResponseWriter, data interface{}) {
	handler(writer, http.StatusCreated, "success", data)
	return
}

func Handler204(writer http.ResponseWriter) {
	handler(writer, http.StatusNoContent, "success", nil)
	return
}

func Handler400(writer http.ResponseWriter, msg string, data interface{}) {
	handler(writer, http.StatusBadRequest, msg, data)
	return
}

func Handler500(writer http.ResponseWriter, msg string) {
	handler(writer, http.StatusInternalServerError, msg, nil)
	return
}

func handler(writer http.ResponseWriter, code int, msg string, data interface{}) {
	d := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	b, err := json.Marshal(d)
	if err != nil {
		return
	}
	writer.WriteHeader(code)
	writer.Write(b)
	return
}

func handlerValidate(w http.ResponseWriter, err error) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		var resultData = make(map[string]interface{})
		for _, err := range err.(validator.ValidationErrors) {
			resultData[strings.ToLower(err.StructField())] = err.ActualTag()
		}
		Handler400(w, err.Error(), resultData)
		return
	} else {
		Handler400(w, err.Error(), nil)
		return
	}
}