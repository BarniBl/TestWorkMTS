package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	e.Use(customMiddlewares.PanicMiddleware)
	e.HTTPErrorHandler = customMiddlewares.CustomHTTPErrorHandler
	handlers := delivery.HandlersStruct{}
	handlers.NewHandlers(e)
	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	err = e.Start(consts.HostAddress)
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}
	e.Logger.Warnf("shutdown")

}
