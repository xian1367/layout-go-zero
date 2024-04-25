package http

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cast"
	"github.com/swaggo/http-swagger/v2"
	_ "github.com/xian1367/layout-go-zero/app/user_http/docs"
	"github.com/xian1367/layout-go-zero/config"
	"github.com/xian1367/layout-go-zero/pkg/http/validator"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

var Server *rest.Server

var Routers Router

type Router interface {
	RegisterRoutes()
}

func Init() {
	validator.Init()

	Server = rest.MustNewServer(
		config.Get().RestConf,
		rest.WithCors("*"),
	)
	defer Server.Stop()

	Routers.RegisterRoutes()
	RegisterSwagger()

	Server.Start()
}

func RegisterSwagger() {
	url := "http://" + config.Get().RestConf.Host + ":" + cast.ToString(config.Get().RestConf.Port) + "/swagger/doc.json"
	spew.Dump(url)
	Server.AddRoute(
		rest.Route{ // 添加路由
			Method: http.MethodGet,
			Path:   "/swagger/:file",
			Handler: httpSwagger.Handler(
				httpSwagger.URL(url), //The url pointing to API definition
			),
		},
	)
}
