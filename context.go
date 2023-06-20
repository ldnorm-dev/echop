package echop

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type context struct {
	echo.Context
}

// Context is a wrapper of echo.Context
type Context interface {
	echo.Context
	Json(httpCode int, resp JSONResponse) error
	// JsonSuccess is a shortcut of JsonSuccessWithCode with DefaultResponseSuccessCode
	JsonSuccess(i any, msg string) error
	// JsonFail is a shortcut of JsonFailWithCode with DefaultResponseFailCode
	JsonFail(i any, msg string) error
	// JsonSuccessWithCode is a shortcut of Json with http code is http.StatusOK and default msg is DefaultResponseSuccessMsg
	JsonSuccessWithCode(code any, data any, msg string) error
	// JsonFailWithCode is a shortcut of Json with http code is http.StatusOK and default msg is DefaultResponseFailMsg
	JsonFailWithCode(code any, data any, msg string) error
	// LogInfo is a shortcut of LogInfoWithContext
	LogInfo(msg string, fields ...zap.Field)
	// LogError is a shortcut of LogErrorWithContext
	LogError(msg string, fields ...zap.Field)
	// LogWarn is a shortcut of LogWarnWithContext
	LogWarn(msg string, fields ...zap.Field)
	// BindAndValidate is a shortcut of echo context Bind and Validate
	BindAndValidate(payload any) error
}

// JSONResponse is a standard response format
type JSONResponse struct {
	Code    any    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// WrapHandlerFunc wraps HandlerFunc to echo.HandlerFunc
func WrapHandlerFunc(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(Context))
	}
}

func (c *context) Json(httpCode int, resp JSONResponse) error {
	return c.Context.JSON(httpCode, resp)
}

// JsonSuccess is a shortcut of JsonSuccessWithCode with DefaultResponseSuccessCode
func (c *context) JsonSuccess(data any, msg string) error {
	return c.JsonSuccessWithCode(DefaultResponseSuccessCode, data, msg)
}

// JsonSuccessWithCode is a shortcut of Json with http code is http.StatusOK and default msg is DefaultResponseSuccessMsg
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

// JsonFailWithCode is a shortcut of Json with http code is http.StatusOK and default msg is DefaultResponseFailMsg
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

// LogInfo is a shortcut of LogInfoWithContext
func (c *context) LogInfo(msg string, fields ...zap.Field) {
	LogInfoWithContext(c, msg, fields...)
}

// LogError is a shortcut of LogErrorWithContext
func (c *context) LogError(msg string, fields ...zap.Field) {
	LogErrorWithContext(c, msg, fields...)
}

// LogWarn is a shortcut of LogWarnWithContext
func (c *context) LogWarn(msg string, fields ...zap.Field) {
	LogWarnWithContext(c, msg, fields...)
}

// BindAndValidate is a shortcut of echo context Bind and Validate
func (c *context) BindAndValidate(payload any) error {
	if err := c.Context.Bind(payload); err != nil {
		return err
	}
	if err := c.Context.Validate(payload); err != nil {
		return err
	}
	return nil
}
