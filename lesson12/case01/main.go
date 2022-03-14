package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// InitMySQL 连接数据库
func InitMySQL() (err error) {
	dsn := "root:gree123@tcp(127.0.0.1:13306)/sql_demo?charset=utf8mb4&parseTime=True"
	// 也可以使用 MustConnect 连接不成功就 panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Connect mysql database failed, err: %v\n", err)
		return
	}
	return
}

func main() {
	// 连接数据库
	if err := InitMySQL(); err != nil {
		panic(err)
	}
	fmt.Println("Connect mysql database success!")
}
