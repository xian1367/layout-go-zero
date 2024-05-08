package http

import (
	"github.com/zeromicro/go-zero/rest"
	"math"
	"net/http"
	"path"
	"regexp"
)

type IRouter interface {
	IRoutes
	Group(string, ...rest.Middleware) *RouterGroup
}

type IRoutes interface {
	Use(...rest.Middleware) IRoutes

	Handle(string, string, HandlerFunc) IRoutes
	Any(string, HandlerFunc) IRoutes
	GET(string, HandlerFunc) IRoutes
	POST(string, HandlerFunc) IRoutes
	DELETE(string, HandlerFunc) IRoutes
	PATCH(string, HandlerFunc) IRoutes
	PUT(string, HandlerFunc) IRoutes
	OPTIONS(string, HandlerFunc) IRoutes
	HEAD(string, HandlerFunc) IRoutes
}

type RouterGroup struct {
	Middlewares Middlewares
	Handler     HandlerFunc
	BasePath    string
}

var _ IRouter = &RouterGroup{}

type HandlerFunc = http.HandlerFunc

type Middlewares = []rest.Middleware

const abortIndex int8 = math.MaxInt8 / 2

func (group *RouterGroup) Use(middleware ...rest.Middleware) IRoutes {
	group.Middlewares = append(group.Middlewares, middleware...)
	return group
}

func (group *RouterGroup) Group(relativePath string, middleware ...rest.Middleware) *RouterGroup {
	return &RouterGroup{
		Middlewares: group.combineHandlers(middleware),
		BasePath:    group.calculateAbsolutePath(relativePath),
	}
}

func (group *RouterGroup) Handle(httpMethod, relativePath string, handler HandlerFunc) IRoutes {
	if matches, err := regexp.MatchString("^[A-Z]+$", httpMethod); !matches || err != nil {
		panic("http method " + httpMethod + " is not valid")
	}
	return group.handle(httpMethod, relativePath, handler)
}

// POST is a shortcut for router.Handle("POST", path, handle).
func (group *RouterGroup) POST(relativePath string, handler HandlerFunc) IRoutes {
	return group.handle(http.MethodPost, relativePath, handler)
}

// GET is a shortcut for router.Handle("GET", path, handle).
func (group *RouterGroup) GET(relativePath string, handler HandlerFunc) IRoutes {
	return group.handle(http.MethodGet, relativePath, handler)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle).
func (group *RouterGroup) DELETE(relativePath string, handler HandlerFunc) IRoutes {
	return group.handle(http.MethodDelete, relativePath, handler)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle).
func (group *RouterGroup) PATCH(relativePath string, handler HandlerFunc) IRoutes {
	return group.handle(http.MethodPatch, relativePath, handler)
}

// PUT is a shortcut for router.Handle("PUT", path, handle).
func (group *RouterGroup) PUT(relativePath string, handler HandlerFunc) IRoutes {
	return group.handle(http.MethodPut, relativePath, handler)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle).
func (group *RouterGroup) OPTIONS(relativePath string, handler HandlerFunc) IRoutes {
	return group.handle(http.MethodOptions, relativePath, handler)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle).
func (group *RouterGroup) HEAD(relativePath string, handler HandlerFunc) IRoutes {
	return group.handle(http.MethodHead, relativePath, handler)
}

// Any registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (group *RouterGroup) Any(relativePath string, handler HandlerFunc) IRoutes {
	group.handle(http.MethodGet, relativePath, handler)
	group.handle(http.MethodPost, relativePath, handler)
	group.handle(http.MethodPut, relativePath, handler)
	group.handle(http.MethodPatch, relativePath, handler)
	group.handle(http.MethodHead, relativePath, handler)
	group.handle(http.MethodOptions, relativePath, handler)
	group.handle(http.MethodDelete, relativePath, handler)
	group.handle(http.MethodConnect, relativePath, handler)
	group.handle(http.MethodTrace, relativePath, handler)
	return group
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handler HandlerFunc) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	Server.AddRoutes(
		rest.WithMiddlewares(
			group.Middlewares,
			[]rest.Route{
				{
					Method:  httpMethod,
					Path:    absolutePath,
					Handler: handler,
				},
			}...,
		),
	)
	return group
}

func (group *RouterGroup) combineHandlers(handlers Middlewares) Middlewares {
	finalSize := len(group.Middlewares) + len(handlers)
	if finalSize >= int(abortIndex) {
		panic("too many handlers")
	}
	mergedHandlers := make(Middlewares, finalSize)
	copy(mergedHandlers, group.Middlewares)
	copy(mergedHandlers[len(group.Middlewares):], handlers)
	return mergedHandlers
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(group.BasePath, relativePath)
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
