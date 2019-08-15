package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type server struct {
	Port  int64
	Token string
}

// Server 服务基础配置
var Server server

func initServer(v *viper.Viper) {
	Server.Port = v.GetInt64("server.port")
	Server.Token = v.GetString("server.token")
}

// AliSMS 阿里短信相关
type aliSMS struct {
	RegionID     string
	AccessKeyID  string
	AccessSecret string
	SignName     string
	TemplateCode string
}

// AliSMS 阿里短信
var AliSMS aliSMS

func initAliSMS(v *viper.Viper) {
	AliSMS.RegionID = v.GetString("sms.regionId")
	AliSMS.AccessKeyID = v.GetString("sms.accessKeyId")
	AliSMS.AccessSecret = v.GetString("sms.accessSecret")
	AliSMS.SignName = v.GetString("sms.signName")
	AliSMS.TemplateCode = v.GetString("sms.templateCode")
}

type dingTalk struct {
	Robot map[string]string
}

// DingTalk 钉钉机器人
var DingTalk dingTalk

func initDingTalk(v *viper.Viper) {
	DingTalk.Robot = v.GetStringMapString("dingtalk.robot")
}

func init() {
	v := viper.New()
	v.SetConfigName("conf")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf(time.Now().Format("2006-01-02 15:04:05"), "Fatal error config file: %s \n", err))
	}
	initServer(v)
	initAliSMS(v)
	initDingTalk(v)
}
