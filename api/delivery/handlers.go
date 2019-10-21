package delivery

import (
	"fmt"
	"github.com/BarniBl/TestWorkMTS/pkg/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"time"
)

func (h *HandlersStruct) HandlerCreateTask(ctx echo.Context) (Err error) {
	taskuuid := uuid.NewV4()

	taskResponse := models.TaskResponse{UUID: taskuuid.String(), CreatedTime: time.Now().String(), Status: "created"}
	fmt.Println(taskuuid)
	return ctx.JSON(202, taskResponse)
}

func (h *HandlersStruct) HandlerGetTaskStatus(ctx echo.Context) (Err error) {

	return nil
}

func (h *HandlersStruct) HandleGetTaskFinished(ctx echo.Context) (Err error) {

	return nil
}
