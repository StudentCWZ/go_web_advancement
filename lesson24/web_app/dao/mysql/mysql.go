/*
   @Author: StudentCWZ
   @Description:
   @File: mysql
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/16 10:13
*/

package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql:host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.db"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("sql.max_open_connection"))
	db.SetMaxIdleConns(viper.GetInt("sql.max_idle_connection"))
	return
}

func Close() {
	err := db.Close()
	if err != nil {
		zap.L().Error("Closed mysql database failed", zap.Error(err))
		return
	}
	return
}
