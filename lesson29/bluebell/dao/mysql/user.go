/*
   @Author: StudentCWZ
   @Description:
   @File: user
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/17 17:00
*/

package mysql

import (
	"GoWeb/lesson29/bluebell/models"
	"errors"
)

// CheckUserExit 把每一步数据库操作封装成函数
// 待 logic 层根据业务需求调用
// CheckUserExit 检查指定用户名的用户是否存在
func CheckUserExit(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 执行 SQL 语句入库
	sqlStr := "insert into user(user_id, username, password) values(?, ?, ?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}
