package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

var db *sql.DB

type User struct {
	Id   int
	Age  int
	Name string
}

// InitMySQl 连接数据库
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
	return
}

// Transaction 事务示例
func Transaction() {
	tx, err := db.Begin() //开启事务
	if err != nil {
		if tx != nil {
			err := tx.Rollback()
			if err != nil {
				return
			} // 回滚
		}
		fmt.Printf("Begin transaction failed, err: %v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
		fmt.Printf("Execute sql1 failed, err: %v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
		fmt.Printf("Execute ret1.RowsAffected() failed, err: %v\n", err)
		fmt.Printf("affRow1: %d\n", affRow1)
		return
	}
	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		err := tx.Rollback() // 回滚
		if err != nil {
			return
		}
		fmt.Printf("Execute sql2 failed, err: %v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return
		}
		fmt.Printf("Execute ret1.RowsAffected() failed, err: %v\n", err)
		fmt.Printf("affRow2: %d\n", affRow2)
		return
	}
	fmt.Printf("affRow1: %d affRow2: %d\n", affRow1, affRow2)
	// 当 affRow1 == 1 && affRow2 ==1
	if affRow1 == 1 && affRow2 == 1 {
		err = tx.Commit() // 提交事务
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			} // 回滚
			fmt.Printf("Commit transaction failed, err: %v\n", err)
			return
		}
		fmt.Println("Execute transaction success!")
	} else {
		err := tx.Rollback() // 回滚
		if err != nil {
			return
		}
		fmt.Println("Transaction rollback ...")
	}
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
	// 事务示例
	Transaction()
}
