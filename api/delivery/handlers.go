package delivery

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
)

func (h *HandlersStruct) HandlerCreateTask(ctx echo.Context) (Err error) {
	taskuuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	fmt.Println(taskuuid)
	return nil
}
