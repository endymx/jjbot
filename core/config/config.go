package config

import (
	"gopkg.in/yaml.v3"
	"jjbot/core/logger"
	"jjbot/core/v12"
	"os"
)

var (
	C = Config{}
)

type Config struct {
	User      string    `yaml:"user"`    //使用此机器人的用户
	Mongodb   []MongoDB `yaml:"mongodb"` //使用的mongodb，留空会爆炸
	Redis     Redis     `yaml:"redis"`
	QQ        string    `yaml:"qq"` //QQ机器人的ws地址，留空为不使用QQ机器人
	BwhVeid   int       `yaml:"bwh_veid"`
	BwhApiKey string    `yaml:"bwh_api_key"`
	OpenaiKey string    `yaml:"openai_key"`
}

type MongoDB struct {
	Name string
	Addr string
}

type Redis struct {
	Addr     string
	Password string
}

func PullConf() {
	err := yaml.Unmarshal(Parse("conf/config.yaml"), &C)
	if err != nil {
		logger.SugarLogger.Fatalf("转换Config错误：%v", err)
	}
}

func Parse(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		v12.SendPrivateMsg(345793738, "QQ机器人获取"+path+"文件失败", false)
		logger.SugarLogger.Fatalf("无法找到%s文件", path)
	}
	return data
}
