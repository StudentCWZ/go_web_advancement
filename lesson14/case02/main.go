package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var (
	db *sqlx.DB
)

type User struct {
	Id   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

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

// QueryByIDs in 查询
func QueryByIDs(ids []int) (users []User, err error) {
	// 动态填充 id
	query, args, err := sqlx.In("SELECT id, name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return nil, err
	}
	// sql.In 返回 `?` 查询语句，我们使用 Rebind() 重新绑定它
	query = db.Rebind(query)
	// 数据查询
	err = db.Select(&users, query, args...)
	return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT id, name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)",
		ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` 查询语句, 我们使用 Rebind() 重新绑定它
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}

func main() {
	// 连接数据库
	if err := InitMySQL(); err != nil {
		panic(err)
	}
	// 输出信息
	fmt.Println("Connect mysql database success!")
	ids := []int{1, 7, 5, 9}
	users, err := QueryByIDs(ids)
	if err != nil {
		fmt.Printf("Query data failed, err: %v\n", err)
		return
	}
	// 输出信息
	fmt.Println("Query data success!")
	fmt.Printf("users: %v\n", users)
	users, err = QueryAndOrderByIDs(ids)
	if err != nil {
		fmt.Printf("Query data failed, err: %v\n", err)
		return
	}
	// 输出信息
	fmt.Println("Query data success!")
	fmt.Printf("users: %v\n", users)
}
