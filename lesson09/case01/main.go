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

// QueryRow 查询单条数据示例
func QueryRow() (err error) {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	// 非常重要：确保 QueryRow 之后调用 Scan 方法，否则持有的数据库连接不会被释放
	row := db.QueryRow(sqlStr, 1)
	err = row.Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("Scan data failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d name: %s age: %d \n", u.Id, u.Name, u.Age)
	return
}

// QueryMultiRow 查询多条数据示例
func QueryMultiRow() (err error) {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("Query failed, err: %v\n", err)
		return
	}
	// 非常重要： 关闭 rows 释放持有的数据库连接
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Printf("Close rows failed, err: %v\n", err)
			return
		}
	}(rows)
	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("Scan data failed, err: %v\n", err)
			continue
		}
		fmt.Printf("id: %d name: %s age: %d \n", u.Id, u.Name, u.Age)
	}
	return
}

// InsertRow 插入单条数据
func InsertRow() (err error) {
	sqlStr := "insert into user(name, age) values (?, ?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("Insert data failed, err: %v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的 id
	if err != nil {
		fmt.Printf("Get last insert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("Insert data success, the id is %d.\n", theID)
	return
}

// UpdateRow 更新单条数据
func UpdateRow() (err error) {
	sqlStr := "Update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("Update data failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("Update data success, affected rows: %d\n", n)
	return
}

// DeleteRow 删除单条数据
func DeleteRow() (err error) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("Delete data failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("Get rowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("Delete data success, affected rows: %d\n", n)
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
	// 查询单条数据
	if err := QueryRow(); err != nil {
		return
	}
	// 查询多条数据
	if err := QueryMultiRow(); err != nil {
		return
	}
	// 插入单条数据
	//if err := InsertRow(); err != nil {
	//	panic(err)
	//}
	// 更新单条数据
	//if err := UpdateRow(); err != nil {
	//	return
	//}
	// 删除单条数据
	//if err := DeleteRow(); err != nil {
	//	return
	//}
}
