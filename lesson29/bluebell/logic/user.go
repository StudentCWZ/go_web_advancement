/*
   @Author: StudentCWZ
   @Description:
   @File: user
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/17 16:57
*/

package logic

import (
	"GoWeb/lesson29/bluebell/dao/mysql"
	"GoWeb/lesson29/bluebell/models"
	"GoWeb/lesson29/bluebell/pkg/encrypt"
	"GoWeb/lesson29/bluebell/pkg/jwt"
	"GoWeb/lesson29/bluebell/pkg/snowflake"
	"fmt"
)

// SignUp 存放业务逻辑代码
func SignUp(p *models.ParamsSignUp) (err error) {
	// 1. 判断用户存不存在
	if err := mysql.CheckUserExit(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2. 生成 UID
	userID := snowflake.GenID()
	// 构造一个 User 实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	fmt.Printf("%#v\n", user)
	// 3. 密码加密
	user.Password = encrypt.EncryptPassword(user.Password)
	fmt.Printf("%#v\n", user)
	// 4. 保存进数据库
	err = mysql.InsertUser(user)
	// redis.xxx ...
	return
}

// Login 存放业务逻辑代码
func Login(p *models.ParamsLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到 user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", user)
	// 生成 JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
