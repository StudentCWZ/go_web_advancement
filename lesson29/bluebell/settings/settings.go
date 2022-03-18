/*
   @Author: StudentCWZ
   @Description:
   @File: settings
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/17 16:07
*/

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(Config)

type Config struct {
	*AppConfig       `mapstructure:"app"`
	*LogConfig       `mapstructure:"log"`
	*MySQLConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*SnowFlakeConfig `mapstructure:"snowflake"`
	*EncryptConfig   `mapstructure:"encrypt"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host              string `mapstructure:"host"`
	User              string `mapstructure:"user"`
	Password          string `mapstructure:"password"`
	Db                string `mapstructure:"db"`
	Port              int    `mapstructure:"port"`
	MaxOpenConnection int    `mapstructure:"max_open_connection"`
	MaxIdleConnection int    `mapstructure:"max_idle_connection"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type SnowFlakeConfig struct {
	StartTime string `mapstructure:"start_time"`
	MachineId int64  `mapstructure:"machine_id"`
}

type EncryptConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件名称（不需要带后缀）
	viper.SetConfigType("yaml")   // 指定配置文件类型
	//viper.AddConfigPath(".")      // 指定查找配置文件的路径（这里使用相对路径）
	viper.AddConfigPath("./conf/")
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.AddConfigPath() failed, err: %v\n", err)
		return
	}
	// 把读取到的信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal() failed, err: %v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Configure file changed ...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal() failed, err: %v\n", err)
		}
	})
	return
}
