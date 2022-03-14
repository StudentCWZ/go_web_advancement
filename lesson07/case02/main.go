package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
	"time"
)

var db *sql.DB

func InitMySQl() (err error) {
	// DSN: Data Source Name
	dsn := "root:gree123@tcp(127.0.0.1:13306)/sql_demo"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// 尝试与数据库建立连接(校验 dsn 是否正确)
	err = db.Ping()
	if err != nil {
		fmt.Printf("Connect mysql databases failed, err: %v\n", err)
		return
	}
	// 数值根据业务具体情况来确定(常用连接配置)
	db.SetConnMaxLifetime(time.Second * 10) // 连接存活的最长时间
	db.SetMaxOpenConns(200)                 // 最大连接数
	db.SetConnMaxIdleTime(50)               // 最大空闲连接数
	return
}

func main() {
	// 连接数据库
	if err := InitMySQl(); err != nil {
		panic(err)
	}
	// close() 用来释放掉数据库连接相关的资源
	defer func(db *sql.DB) { // 注意：defer 语句要在上面的 err 判断下面
		err := db.Close()
		if err != nil {
			fmt.Printf("Close mysql database failed, err: %v\n", err)
		}
	}(db)
	fmt.Println("Connect mysql database success!")
}
