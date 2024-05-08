package route

import (
	"github.com/xian1367/layout-go-zero/app/user_http/controller"
	server "github.com/xian1367/layout-go-zero/pkg/http"
	"github.com/xian1367/layout-go-zero/pkg/http/middleware"
)

func userRoutes() {
	ctrl := new(controller.UserController)
	r := server.Router.Group("/user").Use(middleware.Jwt)
	{
		r.GET("", server.ControllerHandler(ctrl.Index))
		r.GET("/:id", server.ControllerHandler(ctrl.Show))
		r.POST("/:id", server.ControllerHandler(ctrl.Store))
		r.PUT("/:id", server.ControllerHandler(ctrl.Update))
		r.DELETE("/:id", server.ControllerHandler(ctrl.Destroy))
	}
}
