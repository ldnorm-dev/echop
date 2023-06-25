# Echop

[![Go Reference](https://pkg.go.dev/badge/echop.ldnorm.dev.svg)](https://pkg.go.dev/echop.ldnorm.dev)

Echop is a wrapper of echo framework, which provides some useful features.

## Feature Overview

- Standard response body format
- Logger
- Context adds logging methods and standard response body methods
- Logging errors caught by error handling functions
- Request logger middleware and request id middleware are added by default

## Installation

```shell
go get -u echop.ldnorm.dev
```

## [Document](https://pkg.go.dev/echop.ldnorm.dev)

## Example

```go
package main

import (
	echop "echop.ldnorm.dev"
	"github.com/labstack/echo/v4"
)

func main() {
	echop.AppName = "app-name"
	echop.RequestIDConfig.TargetHeader = "Request-Id"
	ep := echop.New()

	group := ep.Group("/api")

	type Payload struct {
		Name string `json:"name" validate:"required"`
	}

	group.POST("/ping", func(c echop.Context) error {
		payload := new(Payload)
		if err := c.BindAndValidate(payload); err != nil {
			// Error will be caught and logged
			return err
		}

		if payload.Name != "test" {
			return c.JsonFail("", "name error")
		} else {
			return c.JsonSuccess(echo.Map{
				"ping": "ping ping pong pong",
			}, "success msg")
		}

	})

	ep.Echo.Start(":1323")
}
```

