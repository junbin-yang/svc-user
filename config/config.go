package config

import (
	"flag"
	"fmt"
	"github.com/junbin-yang/golib/logger"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"svc-user/constant"
)

var (
	VERSION    string = "undefined"
	BUILD_TIME string = "undefined"
	GO_VERSION string = "undefined"
)

type Config struct {
	Release bool
	Listen  string
	Logger  struct {
		Dir      string
		Level    uint
		Rotate   bool
		KeepDays int64
	}
	Mysql struct {
		Name string
		Host string
		Pass string
		Port string
		User string
	}
	Redis struct {
		Cluster bool
		Addr    []string
		Pass    string
	}
	Smtp struct {
		Host string
		Port string
		User string
		Pass string
	}
}

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, constant.AppName+", version: "+VERSION+" (built at "+BUILD_TIME+") "+GO_VERSION)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func Parse() *Config {
	conf := new(Config)
	data, e := ioutil.ReadFile("/etc/" + constant.AppName + ".yml")
	if e != nil {
		panic(e)
	}
	yaml.Unmarshal(data, &conf)

	if conf.Logger.Level <= (uint)(logger.InfoLevel) {
		conf.Release = true
	}

	(&logger.Options{AppName: constant.AppName, Path: conf.Logger.Dir, Level: (logrus.Level)(conf.Logger.Level), Rotate: conf.Logger.Rotate, KeepDays: conf.Logger.KeepDays, TakeStd: true}).New()
	go logger.Asyn()

	return conf
}
