/*
Package echop Group.
Since methods and properties cannot be renamed, and to accommodate the habit of using Group methods to create new route groups, use nominal inheritance and override all methods, regardless of whether the method needs to add WrapHandlerFunc
*/
package echop

import (
	"github.com/labstack/echo/v4"
)

type (
	Group struct {
		G *echo.Group
	}
)

func (g *Group) Use(middleware ...echo.MiddlewareFunc) {
	g.G.Use(middleware...)
}

func (g *Group) CONNECT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.CONNECT(path, WrapHandlerFunc(h), m...)
}

func (g *Group) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.DELETE(path, WrapHandlerFunc(h), m...)
}

func (g *Group) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.GET(path, WrapHandlerFunc(h), m...)
}

func (g *Group) HEAD(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.HEAD(path, WrapHandlerFunc(h), m...)
}
func (g *Group) OPTIONS(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.OPTIONS(path, WrapHandlerFunc(h), m...)
}
func (g *Group) PATCH(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.PATCH(path, WrapHandlerFunc(h), m...)
}
func (g *Group) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.POST(path, WrapHandlerFunc(h), m...)
}
func (g *Group) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.PUT(path, WrapHandlerFunc(h), m...)
}
func (g *Group) TRACE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.TRACE(path, WrapHandlerFunc(h), m...)
}

func (g *Group) Any(path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	return g.G.Any(path, WrapHandlerFunc(handler), middleware...)
}

func (g *Group) Match(methods []string, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	return g.G.Match(methods, path, WrapHandlerFunc(handler), middleware...)
}
func (g *Group) Group(prefix string, m ...echo.MiddlewareFunc) *Group {
	return &Group{g.G.Group(prefix, m...)}
}

func (g *Group) File(path, file string) {
	g.G.File(path, file)
}

func (g *Group) RouteNotFound(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(RouteNotFound, path, h, m...)
}
func (g *Group) Add(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route {
	return g.G.Add(method, path, WrapHandlerFunc(handler), middleware...)
}
