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

// PrepareQuery 预处理查询示例
func PrepareQuery() (err error) {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare failed, err: %v\n", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			return
		}
	}(stmt)
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("Query data failed, err: %v\n", err)
		return
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
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
		fmt.Printf("id: %d name: %s age: %d\n", u.Id, u.Name, u.Age)
	}
	return
}

// PrepareInsert 预处理插入示例
func PrepareInsert() (err error) {
	sqlStr := "Insert into user(name, age) values (?, ?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare failed, err: %v\n", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			return
		}
	}(stmt)
	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("Insert data failed, err: %v\n", err)
		return
	}
	_, err = stmt.Exec("七米", 29)
	if err != nil {
		fmt.Printf("Insert data failed, err: %v\n", err)
		return
	}
	fmt.Println("Insert data success!")
	return
}

// SqlInject sql 注入示例
func SqlInject(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name = '%s'", name)
	fmt.Printf("SQL: %s\n", sqlStr)
	var u User
	err := db.QueryRow(sqlStr).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("Exec failed, err: %v\n", err)
		return
	}
	fmt.Printf("user: %#v\n", u)
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
	// 查询数据
	if err := PrepareQuery(); err != nil {
		return
	}
	// 批量插入数据
	//if err := PrepareInsert(); err != nil {
	//	return
	//}
	// sql 注入示例
	SqlInject("七米")            // 正常
	SqlInject("xxx ' or 1=1#") // 	引发 sql 注入问题
}
