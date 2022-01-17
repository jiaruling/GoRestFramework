package gorms

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strconv"
)

/*
   功能说明: 功能函数
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/15 18:04
*/

func inFields(field string, Fields []string) bool {
	return in(field, Fields)
}

func inExFields(field string, ExFields []string) bool {
	return in(field, ExFields)
}

// 判断元素是否在切片中
func in(f string, fl []string) bool {
	if len(fl) == 0 {
		return false
	}
	for _, item := range fl {
		if f == item {
			return true
		}
	}
	return false
}

// 分页器
func Paging(r *http.Request, PageMax, PageMin int64) (page, pageSize int) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	// 页码
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}

	// 每一页大小
	pageSize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		if PageMin > 0 {
			pageSize = int(PageMin)
		} else {
			pageSize = int(GlobalPageMin)
		}
	}
	if PageMax > 0 {
		if pageSize > int(PageMax) {
			pageSize = int(PageMax)
		}
	} else {
		if pageSize > int(GlobalPageMax) {
			pageSize = int(GlobalPageMax)
		}
	}
	if PageMin > 0 {
		if pageSize <= 0 {
			pageSize = int(PageMin)
		}
	} else {
		if pageSize <= 0 {
			pageSize = int(GlobalPageMin)
		}
	}
	return
}

// 验证器
func Validate(r *http.Request, m interface{}) (err error) {
	var (
		buf      = make([]byte, 4096)
		validate = validator.New()
	)
	n, err := r.Body.Read(buf)
	if err != nil && err != io.EOF {
		return
	}
	err = json.Unmarshal(buf[:n], m)
	if err != nil {
		return
	}
	if err = validate.Struct(m); err != nil {
		return
	}
	return nil
}
