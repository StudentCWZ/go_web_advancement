/*
   @Author : cuiweizhi
*/

package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ = zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))
	} else {
		sugarLogger.Info("Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		err = resp.Body.Close()
		if err != nil {
			return
		}
	}
}

func main() {
	InitLogger()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Printf("write log failed, err: %v\n", err)
		}
	}(logger)
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("https://www.baidu.com")
}
