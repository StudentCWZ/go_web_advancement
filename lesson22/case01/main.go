/*
   @Author : cuiweizhi
*/

package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MysqlConfig `mapstructure:"mysql"`
}

type MysqlConfig struct {
	Host string `mapstructure:"host"` // 一定记住要写 tag
	Port int    `mapstructure:"port"`
	Db   string `mapstructure:"db"`
}

func readConfig() (err error) {
	// 设置默认值
	viper.SetDefault("fileDir", "./")
	// 读取配置文件
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.SetConfigFile("config.yaml") 等价以上两步
	viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在路径
	viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")              // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()            // 查找并读取配置文件
	return err
}

func main() {
	// 读取配置文件
	if err := readConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			fmt.Println("config file not found ...")
			panic(err)
		} else {
			// 配置文件找到，但产生了另外的错误
			panic(err)
		}
	}
	// 实时监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("Unmarshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("c: %#v\n", c)
}
