package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	Common
	Http
	Cmd
	Cron
	Queue
}

type Common struct {
	redis.RedisConf
	Database
	Debug bool `json:",default=true"`
}

type Http struct {
	rest.RestConf
}

type Cmd struct {
	logx.LogConf
}

type Cron struct {
	logx.LogConf
}

type Queue struct {
	logx.LogConf
}

type Database struct {
	Connection string `json:",optional"`
	Mysql      struct {
		Host               string `json:",default=0.0.0.0"`
		Port               string `json:",default=3306"`
		Username           string `json:",optional"`
		Password           string `json:",optional"`
		DBName             string `json:",optional"`
		Charset            string `json:",optional"`
		MaxIdleConnections int    `json:",optional"`
		MaxOpenConnections int    `json:",optional"`
		MaxLifeSeconds     int    `json:",optional"`
	} `json:",optional"`
	Postgres struct {
		Host     string `json:",default=0.0.0.0"`
		Port     string `json:",default=6379"`
		Username string `json:",optional"`
		Password string `json:",optional"`
		DBName   string `json:",optional"`
	} `json:",optional"`
}
