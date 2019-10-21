package main

import (
	"github.com/BarniBl/TestWorkMTS/api/delivery"
	"github.com/BarniBl/TestWorkMTS/api/repository"
	"github.com/BarniBl/TestWorkMTS/pkg/consts"
	customMiddleware "github.com/BarniBl/TestWorkMTS/pkg/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: consts.LoggerFormat}))
	e.Use(customMiddleware.PanicMiddleware)
	e.HTTPErrorHandler = customMiddleware.CustomHTTPErrorHandler
	var worker repository.RepositoryStruct
	if err := worker.NewDataBaseWorker(); err != nil {
		e.Logger.Errorf("server error: %s", err)
		return
	}
	handlers := delivery.HandlersStruct{RepositoryWorker: &worker}
	handlers.NewHandlers(e)
	e.Logger.Warnf("start listening on %s", consts.HostAddress)
	err := e.Start(consts.HostAddress)
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}
	e.Logger.Warnf("shutdown")

}
