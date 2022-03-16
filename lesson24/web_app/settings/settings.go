/*
   @Author: StudentCWZ
   @Description:
   @File: settings
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/16 09:13
*/

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")  // 指定配置文件名称（不需要带后缀）
	viper.SetConfigType("yaml")    // 指定配置文件类型
	viper.AddConfigPath("./conf/") // 指定查找配置文件的路径（这里使用相对路径）
	err = viper.ReadInConfig()     // 读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("configure file changed ...")
	})
	return
}
