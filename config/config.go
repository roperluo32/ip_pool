package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// RedisConfig 定义配置文件解析后的结构
type RedisConfig struct {
	URL      string
	Password string
}

// XunDaiLiConfig 讯代理配置
type XunDaiLiConfig struct {
	OrderNo    string
	ReturnType string
	Count      string
	SpiderID   string
	Interval   int // 秒。多久去拉取一次
	Timeout    int // 秒。多久超时。Timeout最好小于Interval
}

// Config 配置文件
type Config struct {
	Port     int
	Redis    RedisConfig
	XunDaiLi XunDaiLiConfig
}

// C 全局配置
var C Config

// Init 初始化配置
func Init(name, path string) {
	viper.SetConfigName(name)   // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(path)   // 第一个搜索路径
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("read config file fail. err: %s", err))
	}
	viper.Unmarshal(&C) // 将配置信息绑定到结构体上
	fmt.Println(C)

	viper.WatchConfig()
}
