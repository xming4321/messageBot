package conf

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"runtime"
)

type TgBot struct {
	Token   string `yaml:"token"`
	Debug   bool   `yaml:"debug"`
	Timeout int    `yaml:"timeout"`
}

type TConf struct {
	Port     string `yaml:"port"`
	LogLevel string `yaml:"logLevel"`
	Mysql    struct {
		Addr         string `yaml:"addr"`
		User         string `yaml:"user"`
		PassWord     string `yaml:"password"`
		DataBase     string `yaml:"database"`
		MaxIdleConns int    `yaml:"maxidleconns"`
		MaxOpenConns int    `yaml:"maxopenconns"`
	}
	TgBot TgBot `yaml:"tgBot"`
}

var Conf TConf

func init() {
	yamlFile, err := ioutil.ReadFile(getCurrentPath() + "/config.yaml")
	if err != nil {
		logrus.Panicf("yamlfile get error: %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		logrus.Panicf("yaml unmarshal error: %v", err)
	}
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}
