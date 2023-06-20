/*
Package echop is a wrapper of echo framework, which provides some useful features.
*/
package echop

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type (
	Echop struct {
		*echo.Echo
	}
	HandlerFunc func(c Context) error
)

var (
	AppName = "echop"
	// RequestIDConfig middleware.RequestIDWithConfig
	RequestIDConfig     = middleware.DefaultRequestIDConfig
	RequestLoggerConfig = middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogRemoteIP: true,
		LogHost:     true,
		LogLatency:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var logFunc func(c echo.Context, msg string, fields ...zap.Field)
			if v.Status >= 500 {
				logFunc = LogErrorWithContext
			} else if v.Status >= 400 {
				logFunc = LogWarnWithContext
			} else {
				logFunc = LogInfoWithContext
			}
			logFunc(
				c, "request",
				zap.String("request_id", GetRequestId(c)),
				zap.String("method", v.URI),
				zap.String("method", v.Method),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("host", v.Host),
				zap.String("latency", v.Latency.String()),
			)
			return nil
		},
	}
	DefaultResponseSuccessCode = 0
	DefaultResponseFailCode    = 1
	DefaultResponseSuccessMsg  = "ok"
	DefaultResponseFailMsg     = "fail"
	RouteNotFound              = "echo_route_not_found"
	Logger                     *zap.Logger
)

func OverrideContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context{Context: c}
			return next(cc)
		}
	}
}

func GetRequestId(c echo.Context) (id string) {
	id = c.Request().Header.Get(RequestIDConfig.TargetHeader)
	if id == "" {
		id = c.Response().Header().Get(RequestIDConfig.TargetHeader)
	}
	return
}

// New creates an instance of Echop.
//
// It will set the following echo middlewares:
//
// middleware.RequestIDWithConfig
// middleware.RequestLoggerWithConfig
func New() (ep *Echop) {
	e := echo.New()
	ep = &Echop{e}
	e.Validator = NewValidator()
	e.HTTPErrorHandler = ep.DefaultHTTPErrorHandler
	e.Use(middleware.RequestIDWithConfig(RequestIDConfig))
	e.Use(OverrideContext())
	e.Use(middleware.RequestLoggerWithConfig(RequestLoggerConfig))
	return
}

func (ep *Echop) DefaultHTTPErrorHandler(err error, c echo.Context) {
	LogErrorWithContext(c, "http error handler", zap.Error(err))

	// If error is not an echo.HTTPError, wrap it with echo.HTTPError.
	//if _, ok := err.(*echo.HTTPError); !ok {
	//	err = &echo.HTTPError{
	//		Code:    http.StatusInternalServerError,
	//		Message: err.Error(),
	//	}
	//}

	ep.Echo.DefaultHTTPErrorHandler(err, c)
}

func (ep *Echop) CONNECT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.CONNECT(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.DELETE(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.GET(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) HEAD(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.HEAD(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) OPTIONS(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.OPTIONS(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) PATCH(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.PATCH(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.POST(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.PUT(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) TRACE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.TRACE(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) RouteNotFound(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Add(RouteNotFound, path, h, m...)
}

func (ep *Echop) Any(path string, h HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return ep.Echo.Any(path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) Match(methods []string, path string, h HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return ep.Echo.Match(methods, path, WrapHandlerFunc(h), m...)
}

func (ep *Echop) Group(prefix string, m ...echo.MiddlewareFunc) *Group {
	return &Group{ep.Echo.Group(prefix, m...)}
}

func (ep *Echop) URI(handler HandlerFunc, params ...interface{}) string {
	return ep.Echo.URI(WrapHandlerFunc(handler), params...)
}

func (ep *Echop) URL(handler HandlerFunc, params ...interface{}) string {
	return ep.Echo.URL(WrapHandlerFunc(handler), params...)
}

func (ep *Echop) Add(method, path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return ep.Echo.Add(method, path, WrapHandlerFunc(h), m...)
}
