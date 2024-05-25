package initialize

import (
	"fmt"
	"github.com/ohwin/core/config"
	"github.com/spf13/viper"
)

// Viper 解析配置
func Viper() {
	viperConfig := viper.New()
	// 设置配置文件名，没有后缀
	viperConfig.SetConfigName("config")
	// 设置读取文件格式为: yaml
	viperConfig.SetConfigType("yaml")
	// 设置配置文件目录(可以设置多个,优先级根据添加顺序来)
	viperConfig.AddConfigPath(".")
	viperConfig.AddConfigPath("./config")
	// 读取解析
	if err := viperConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("配置文件未找到！%v\n", err)
			return
		} else {
			fmt.Printf("找到配置文件,但是解析错误,%v\n", err)
			return
		}
	}
	// 映射到结构体
	var serverConfig config.ServerConfig
	if err := viperConfig.Unmarshal(&serverConfig); err != nil {
		fmt.Printf("配置映射错误,%v\n", err)
	}
	fmt.Printf("config: %+v\n", serverConfig)

}
