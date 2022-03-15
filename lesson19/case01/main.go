/*
   @Author : cuiweizhi
*/

package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	// 方式一
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // ISO 时间格式
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//return zapcore.NewJSONEncoder(encoderConfig)
	// 方式二：自定义结构体
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./test.log")	// 每次创建
	//file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744) // 追加
	//return zapcore.AddSync(file)
	lumberJackLogger := &lumberjack.Logger{ // 日志切割和归档
		Filename:   "./test.log",
		MaxSize:    10,    // m
		MaxBackups: 5,     // 备份
		MaxAge:     30,    // 最大备份天数
		Compress:   false, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
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
	defer func(sugarLogger *zap.SugaredLogger) {
		err := sugarLogger.Sync()
		if err != nil {
			fmt.Printf("write log failed, err: %v\n", err)
		}
	}(sugarLogger)
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("https://www.baidu.com")
}
