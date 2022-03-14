package main

import (
	"database/sql/driver"
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

// Value 接口
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
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

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*User) (err error) {
	// 存放 (?, ?) 的 slice
	valueStrings := make([]string, 0, len(users))
	// 存放 values 的 slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历 users 准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自己拼接语句要执行的具体操作
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s", strings.Join(valueStrings, ","))
	_, err = db.Exec(stmt, valueArgs...)
	return
}

// BatchInsertUsersTwo 批量插入
func BatchInsertUsersTwo(users []interface{}) (err error) {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果 arg 实现了 driver.Valuer，sqlx.In 会通过调用 value 来展开它
	)
	fmt.Println(query) // 查看生成的 querystring
	fmt.Println(args)  // 查看生成的 args
	_, err = db.Exec(query, args...)
	return
}

// BatchInsertUsersThree 批量插入
// 注意：该功能需 1.3.1 版本以上，并且 1.3.1 版本目前还有点问题，sql 语句最后不能有空格和 ;。
func BatchInsertUsersThree(users []*User) (err error) {
	_, err = db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return
}

func main() {
	// 连接数据库
	if err := InitMySQL(); err != nil {
		panic(err)
	}
	fmt.Println("Connect mysql database success!")
	u1 := User{Name: "宋江", Age: 18}
	u2 := User{Name: "公孙胜", Age: 28}
	u3 := User{Name: "卢俊义", Age: 38}
	// 方法一
	users := []*User{&u1, &u2, &u3}
	err := BatchInsertUsers(users)
	if err != nil {
		fmt.Printf("BatchInsertUsers failed, err: %v\n", err)
		return
	}
	fmt.Println("Batch insert data success!")
	// 方法二
	users2 := []interface{}{u1, u2, u3}
	err = BatchInsertUsersTwo(users2)
	if err != nil {
		fmt.Printf("BatchInsertUsersTwo failed, err: %v\n", err)
		return
	}
	fmt.Println("Batch insert data success!")
	// 方法三
	users3 := []*User{&u1, &u2, &u3}
	err = BatchInsertUsersThree(users3)
	if err != nil {
		fmt.Printf("BatchInsertUsersThree failed, err: %v\n", err)
		return
	}
	fmt.Println("Batch insert data success!")
}
