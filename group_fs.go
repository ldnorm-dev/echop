package echop

import (
	"github.com/labstack/echo/v4"
	"io/fs"
)

func (g *Group) Static(pathPrefix, fsRoot string) {
	g.G.Static(pathPrefix, fsRoot)
}

func (g *Group) StaticFS(pathPrefix string, filesystem fs.FS) {
	g.G.StaticFS(pathPrefix, filesystem)
}

func (g *Group) FileFS(path, file string, filesystem fs.FS, m ...echo.MiddlewareFunc) *echo.Route {
	return g.G.FileFS(path, file, filesystem, m...)
}
