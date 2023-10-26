package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var envPtr = pflag.String("env", "dev", "Environment: dev or prod")

func NewConfig() *viper.Viper {
	v := viper.New()

	pflag.Parse()

	// 获取项目根路径
	dir, _ := os.Getwd()

	// 读取路径: 项目根路径/configs/
	v.AddConfigPath(dir)

	// 配置文件名称
	v.SetConfigName(fmt.Sprintf("config-%s", *envPtr))

	// 配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Panicf("Failed to read in config err: %s \n", err)
	}

	return v
}
