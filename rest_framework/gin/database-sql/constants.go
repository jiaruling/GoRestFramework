package grf

import (
	"database/sql"
)

/*
   功能说明: 包内常量
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/16 16:20
*/

var (
	RDB           *sql.DB
	WDB           *sql.DB
	GlobalPageMax int64
	GlobalPageMin int64
	LogPath       = "./database-sql.log"
)
