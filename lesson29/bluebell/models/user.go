/*
   @Author: StudentCWZ
   @Description:
   @File: user
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/18 10:41
*/

package models

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}
