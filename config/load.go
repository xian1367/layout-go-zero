package config

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"path/filepath"
	"strings"
)

var config = new(Config)

func Init(name string) {
	var configFile = flag.String("c", "./config/"+name+".yaml", "the config file")
	flag.Parse()

	conf.MustLoad(filepath.Dir(*configFile)+"/common.yaml", &config.Common)

	if strings.Contains(name, "http") {
		conf.MustLoad(*configFile, &config.Http)
	} else if name == "cmd" {
		conf.MustLoad(*configFile, &config.Cmd)
	} else if name == "cron" {
		conf.MustLoad(*configFile, &config.Cron)
	} else if name == "queue" {
		conf.MustLoad(*configFile, &config.Queue)
	}
}

func Get() Config {
	return *config
}
