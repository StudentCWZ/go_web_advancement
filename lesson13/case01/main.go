package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

// QueryRow 查询单条数据示例
func QueryRow() (err error) {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err = db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("Query data failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d name: %s age: %d\n", u.Id, u.Name, u.Age)
	return
}

// QueryMultiRow 查询多行数据示例
func QueryMultiRow() (err error) {
	sqlStr := "select id, name, age from user where id > ?"
	var users []User
	err = db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("Query data failed, err: %v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", users)
	return
}

// InsertRow 插入数据
func InsertRow() (err error) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("Insert data failed, err: %v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的 id
	if err != nil {
		fmt.Printf("Get last insertID failed, err: %v\n", err)
		return
	}
	fmt.Printf("Insert data success, the id is %d.\n", theID)
	return
}

// UpdateRow 更新数据
func UpdateRow() (err error) {
	sqlStr := "Update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("Update data failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("Get ret.RowsAffected() failed, err: %v\n", err)
		return
	}
	fmt.Printf("Update data success, affected row: %d\n", n)
	return
}

// DeleteRow 删除数据
func DeleteRow() (err error) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响行数
	if err != nil {
		fmt.Printf("Get ret.RowsAffected() failed, err: %v\n", err)
		return
	}
	fmt.Printf("Delete data success, affected row: %d\n", n)
	return
}

// InsertUser NamedExec 使用
func InsertUser() (err error) {
	_, err = db.NamedExec(`INSERT INTO user (name, age) values (:name, :age)`,
		map[string]interface{}{
			"name": "七米",
			"age":  17,
		})
	if err != nil {
		fmt.Printf("Insert data failed, err: %v\n", err)
		return
	}
	return
}

// QueryUser NamedQuery 使用
func QueryUser() (err error) {
	// 使用 map 做命名查询
	rows, err := db.NamedQuery(`select * from user where name=:name`, map[string]interface{}{"name": "q1mi"})
	if err != nil {
		fmt.Printf("db.NamedQuery() failed, err: %v\n", err)
		return
	}
	defer func(rows *sqlx.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var u User
		err = rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			continue
		}
		fmt.Printf("user: %#v\n", u)
	}
	return
}

// Transaction 事务
func Transaction() (err error) {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		fmt.Printf("Begin transaction failed, err: %v\n", err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("Transaction rollback ...")
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
			fmt.Println("Transaction commit ...")
		}
	}()
	sqlStr1 := "Update user set age=20 where id=?"
	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return
	}
	if n != 1 {
		return errors.New("execute sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where id=?"
	rs, err = tx.Exec(sqlStr2, 3)
	if err != nil {
		return
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return
	}
	if n != 1 {
		return errors.New("execute sqlStr2 failed")
	}
	return
}

func main() {
	// 连接数据库
	if err := InitMySQL(); err != nil {
		panic(err)
	}
	fmt.Println("Connect mysql database success!")
	// 查询数据
	if err := QueryRow(); err != nil {
		return
	}
	// 批量查询数据
	if err := QueryMultiRow(); err != nil {
		return
	}
	// 插入数据
	//if err := InsertRow(); err != nil {
	//	return
	//}
	// 更新数据
	//if err := UpdateRow(); err != nil {
	//	return
	//}
	// 删除数据
	//if err := DeleteRow(); err != nil {
	//	return
	//}
	// NamedExec 使用
	//if err := InsertUser(); err != nil {
	//	return
	//}
	// NamedQuery 使用
	if err := QueryUser(); err != nil {
		return
	}
	// 事务
	if err := Transaction(); err != nil {
		return
	}
}
