package proxy

import (
	"github.com/babaev/logrus_proxy_hook"

	"github.com/sirupsen/logrus"
	"github.com/sniperkit/logrus_mate"
)

type Config struct {
	AppName          string
	Address          string
	Protocol         string
	Host             string // 0.0.0.0
	Port             int    // 5044
	LJV              int
	Keepalive        int
	Timeout          int
	AlwaysSentFields logrus.Fields
}

func init() {
	logrus_mate.RegisterHook("Config", NewProxyHook)
}

func NewProxyHook(config logrus_mate.Configuration) (hook logrus.Hook, err error) {
	conf := Config{}

	if config != nil {

		conf.AppName = config.GetString("app-name")
		conf.Address = config.GetString("address")
		conf.Host = config.GetString("host")
		conf.Port = config.GetString("port")
		conf.Protocol = config.GetString("protocol")
		conf.LJV = config.GetInt32("ljv")
		conf.Keepalive = config.GetInt32("keepalive")
		conf.Timeout = config.GetInt32("timeout")

		alwaysSentFieldsConf := config.GetConfig("always-sent-fields")
		keys := alwaysSentFieldsConf.Keys()
		fields := make(logrus.Fields, len(keys))

		for i := 0; i < len(keys); i++ {
			fields[keys[i]] = alwaysSentFieldsConf.GetString(keys[i])
		}

		conf.AlwaysSentFields = fields
	}
	hook, err = logrus_proxy_hook.NewHook()

	return
}
