package route

import (
	"github.com/xian1367/layout-go-zero/app/user_http/controller"
	server "github.com/xian1367/layout-go-zero/pkg/http"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func testRoutes() {
	ctrl := new(controller.TestController)
	server.Server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: server.ControllerFunc(ctrl.Test),
		},
		rest.WithPrefix("/test"),
	)
}
