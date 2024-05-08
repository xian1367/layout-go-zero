package main

import (
	"github.com/xian1367/layout-go-zero/app/user_http/route"
	"github.com/xian1367/layout-go-zero/config"
	"github.com/xian1367/layout-go-zero/pkg/console"
	"github.com/xian1367/layout-go-zero/pkg/http"
	"github.com/xian1367/layout-go-zero/pkg/orm"
)

func main() {
	config.Init("user_http")
	orm.Init()
	console.UI()

	http.RouterRegister = route.Routes{}
	http.Init()
}
