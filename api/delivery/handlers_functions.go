package delivery

import (
	"github.com/go-park-mail-ru/2019_2_Solar/pinterest/usecase"
	"github.com/labstack/echo"
)

func (h *HandlersStruct) NewHandlers(e *echo.Echo) {
	e.GET("/task", h.HandleEmpty)
	e.GET("/task/:taskId", h.HandleListUsers)
	e.GET("//task/:taskId/finished", h.HandleGetUserByEmail)
}
