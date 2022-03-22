/*
   @Author: StudentCWZ
   @Description:
   @File: jwt
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/21 09:18
*/

package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("迷途小书童")

// MyClaims 自定义声明结构体并内嵌 jwt.StandardClaims
// jwt 包自带的 jwt.StandardClaims 只包含了官方字段
// 我们这里需要额外记录一个 UserID 字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成 JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		UserID:   userID,
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的 secret 签名并获得完整的编码后的字符串 token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验 token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
