/*
   @Author: StudentCWZ
   @Description:
   @File: mysql
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/16 14:50
*/

package mysql

import (
	"GoWeb/lesson25/web_app/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Db,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConnection)
	db.SetMaxIdleConns(cfg.MaxIdleConnection)
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
