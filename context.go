package echop

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type context struct {
	echo.Context
}
type Context interface {
	echo.Context
	Json(httpCode int, resp JSONResponse) error
	JsonSuccess(i any, msg string) error
	JsonFail(i any, msg string) error
	JsonSuccessWithCode(code any, data any, msg string) error
	JsonFailWithCode(code any, data any, msg string) error
	LogInfo(msg string, fields ...zap.Field)
	LogError(msg string, fields ...zap.Field)
	LogWarn(msg string, fields ...zap.Field)
	BindAndValidate(payload any) error
}

type JSONResponse struct {
	Code    any    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func WrapHandlerFunc(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(Context))
	}
}

func (c *context) Json(httpCode int, resp JSONResponse) error {
	return c.Context.JSON(httpCode, resp)
}

func (c *context) JsonSuccess(data any, msg string) error {
	return c.JsonSuccessWithCode(DefaultResponseSuccessCode, data, msg)
}

func (c *context) JsonSuccessWithCode(code any, data any, msg string) error {
	if msg == "" {
		msg = DefaultResponseSuccessMsg
	}

	return c.Json(http.StatusOK, JSONResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func (c *context) JsonFail(data any, msg string) error {
	return c.JsonFailWithCode(DefaultResponseFailCode, data, msg)
}

func (c *context) JsonFailWithCode(code any, data any, msg string) error {
	if msg == "" {
		msg = DefaultResponseFailMsg
	}
	return c.Json(http.StatusOK, JSONResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func (c *context) LogInfo(msg string, fields ...zap.Field) {
	LogInfoWithContext(c, msg, fields...)
}

func (c *context) LogError(msg string, fields ...zap.Field) {
	LogErrorWithContext(c, msg, fields...)
}

func (c *context) LogWarn(msg string, fields ...zap.Field) {
	LogWarnWithContext(c, msg, fields...)
}

func (c *context) BindAndValidate(payload any) error {
	if err := c.Context.Bind(payload); err != nil {
		return err
	}
	if err := c.Context.Validate(payload); err != nil {
		return err
	}
	return nil
}
