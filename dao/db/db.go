package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"sz_resume_202005/utils/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//Init 初始化数据库
func Init() {
	err := innoDB(setting.MysqlAddr)
	if err != nil {
		log.Printf("connect database failed,err:%v\n", err)
	}

}
func innoDB(dsn string) (err error) {
	db, err = sqlx.Connect("mysql", dsn)
	return

}

//ToNullString validates a sql.NullInt64
func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

//ToNullInt64 validates a sql.NullInt64 if incoming string evaluates to an integer, invalidates if it does not
func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
}

func batchStringParam(n int) (str string) {
	if n >= 1 {
		str = "(?)"
		for i := 1; i < n; i++ {
			str += ",(?)"
		}
	}
	fmt.Println(str)
	return str
}
