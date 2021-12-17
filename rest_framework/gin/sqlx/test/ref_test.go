package test

import (
	"fmt"
	"reflect"
	"strings"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/17 14:44
*/


func main() {
	var s interface{}
	s = &Student{}
	//fmt.Println(strings.Split(s, ",")[0])
	t := reflect.TypeOf(s)
	for i := 0; i < t.Elem().NumField(); i++ {
		tag := t.Elem().Field(i).Tag.Get("binding")
		f := t.Elem().Field(i).Type.Name()
		fmt.Printf("%s -- %T\n", f, f)
		if tag != "" {
			fmt.Println(in("required", strings.Split(tag, ",")))
		}
	}
}

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
