package database_sql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

/*
   功能说明: 数据库操作
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 16:03
*/

// 增删改
func ExecDB(sql string) (lastId int64, err error) {
	result, err := WDB.Exec(sql)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

// 列表查询
func getListDB(sql string, m interface{}, filed []string) (list []map[string]interface{}, err error) {
	fmt.Println(sql)
	list = make([]map[string]interface{}, 0)
	rows, err := RDB.Query(sql)
	if err != nil {
		return nil, err
	}
	l, err := fieldHandle(filed, m)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		newMap :=  make(map[string]interface{})
		err = rows.Scan(l...)
		if err != nil {
			return nil, err
		}
		b, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, &newMap); err != nil {
			return nil, err
		}
		list = append(list, newMap)
	}
	return
}

// 查询记录总数
func getTotalDB(sql string) (total int64) {
	_ = RDB.QueryRow(sql).Scan(&total)
	return
}

// 单条记录查询
func getByIdDB(m interface{}, sql string, filed []string) (err error) {
	l, err := fieldHandle(filed, m)
	if err != nil {
		return err
	}
	err = RDB.QueryRow(sql).Scan(l...)
	if err != nil {
		return err
	}
	return nil
}

// 查询字段处理
func fieldHandle(filed []string, m interface{}) (l []interface{},err error) {
	n := m
	t := reflect.TypeOf(n)
	v := reflect.ValueOf(n)
	l = make([]interface{},0)
	for _, item := range filed {
		for i := 0; i < t.Elem().NumField(); i++ {
			f := t.Elem().Field(i)
			if item == strings.Split(f.Tag.Get("json"), ",")[0] {
				l = append(l, v.Elem().Field(i).Addr().Interface())
				break
			}
		}
	}
	fmt.Println(l)
	return
}