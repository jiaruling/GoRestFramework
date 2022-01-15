package main

import (
	database_sql "GoRestFramework/rest_framework/gin/database-sql"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:46
*/

type Student struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty" db:"name" binding:""`
	Age       int64  `json:"age,omitempty" db:"age" binding:"required"`
	Class     string `json:"class,omitempty" db:"class" binding:"required"`
	CreatedAt int64  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt int64  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt int64  `json:"deleted_at,omitempty" db:"deleted_at"`
}

var stu = database_sql.Model{
	M:     nil, // M: new(Student) 传入模型的结构体指针
	Table: "student",
	CreateField: database_sql.CreateField{
		CreatedFields:        nil,
		CreatedIgnoreFields:  []string{"deleted_at"},
		CreatedSetTimeFields: []string{"created_at", "updated_at"},
	},
	SoftDeleteField: database_sql.SoftDeleteField{
		DeletedFields: "deleted_at",
	},
	UpdateField: database_sql.UpdateField{
		UpdateFields:        nil,
		UpdateIgnoreFields:  []string{"created_at", "deleted_at"},
		UpdateSetTimeFields: []string{"updated_at"},
	},
	SelectField: database_sql.SelectField{
		SelectFields:       nil,
		SelectIgnoreFields: []string{"deleted_at"},
	},
	SelectFieldList: database_sql.SelectFieldList{
		Search:  []string{"name", "age"},
		Filter:  nil,
		Sort:    []string{"id"},
		PageMax: 100,
		PageMin: 17,
	},
}

type Teacher struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	Age       int64  `json:"age" db:"age" binding:"required"`
	Subject   string `json:"subject" db:"subject" binding:"required"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
	DeletedAt int64  `json:"deleted_at" db:"deleted_at"`
}
