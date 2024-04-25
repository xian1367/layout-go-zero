package route

import (
	"github.com/xian1367/layout-go-zero/app/user_http/controller"
	server "github.com/xian1367/layout-go-zero/pkg/http"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func userRoutes() {
	ctrl := new(controller.UserController)
	server.Server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: server.ControllerHandler(ctrl.Index),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: server.ControllerHandler(ctrl.Show),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: server.ControllerHandler(ctrl.Update),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: server.ControllerHandler(ctrl.Destroy),
			},
		},
		rest.WithPrefix("/user"),
	)
}
