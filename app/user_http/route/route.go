package route

type Routes struct{}

// RegisterRoutes 注册 API 相关路由
func (r Routes) RegisterRoutes() {
	testRoutes()
	userRoutes()
}
