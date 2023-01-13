package config

import (
	"github.com/spf13/viper"
	"log"
)

var GlobalConfig Config

func init() {
	config := viper.New()
	config.AddConfigPath("./")     //设置读取的文件路径
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yaml")   //设置文件的类型
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	err := config.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Println(err.Error())
	}
}

const (
	PrefixModel     = 1
	AutoAnswerModel = 2
)

type Config struct {
	GptToken     string `mapstructure:"gpt_token"`
	WorkMode     int8   `mapstructure:"work_mode"` // 1-前缀问答模式 2-任务模式
	StartKeyWord string `mapstructure:"start_key_word"`
	EndKeyWord   string `mapstructure:"end_key_word"`
	PrefixWord   string `mapstructure:"prefix_word"`
}
