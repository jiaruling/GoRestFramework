package test

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:46
*/

type Student struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty" db:"name" binding:""`
	Age         int64  `json:"age,omitempty" db:"age" binding:"required"`
	Class       string `json:"class,omitempty" db:"class" binding:"required"`
	CreatedAt   int64  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   int64  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAtAt int64  `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Teacher struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Age         int64  `json:"age" db:"age" binding:"required"`
	Subject     string `json:"subject" db:"subject" binding:"required"`
	CreatedAt   int64  `json:"created_at" db:"created_at"`
	UpdatedAt   int64  `json:"updated_at" db:"updated_at"`
	DeletedAtAt int64  `json:"deleted_at" db:"deleted_at"`
}