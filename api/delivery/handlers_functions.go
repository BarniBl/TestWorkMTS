package delivery

import (
	"github.com/labstack/echo"
)

func (h *HandlersStruct) NewHandlers(e *echo.Echo) {
	e.GET("/task", h.HandlerCreateTask)
	e.GET("/task/:taskId", h.HandlerGetTaskStatus)
	e.GET("/task/:taskId/finished", h.HandleGetTaskFinished)
}
