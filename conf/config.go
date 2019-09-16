package conf

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type Config struct {
	JwtSecret 		string
}

var AppConfig = &Config{}

var cfg *ini.File

func SetUp() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		panic(err)
	}

	mapTo("app", AppConfig)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logrus.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
