package grf

import (
	"encoding/json"
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
func getListDB(sql string, m interface{}) (list []map[string]interface{}, err error) {
	list = make([]map[string]interface{}, 0)
	rows, err := RDB.Queryx(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(m)
		if err != nil {
			return nil, err
		}
		b, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		newMap :=  make(map[string]interface{})
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
func getByIdDB(m interface{}, sql string) (err error) {
	err = RDB.Get(m, sql)
	if err != nil {
		return err
	}
	return nil
}
