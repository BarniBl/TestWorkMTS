package delivery

import (
	"errors"
	"fmt"
	"github.com/BarniBl/TestWorkMTS/pkg/consts"
	"github.com/BarniBl/TestWorkMTS/pkg/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
)

func (h *HandlersStruct) HandlerCreateTask(ctx echo.Context) (Err error) {
	taskuuid := uuid.NewV4()
	var params []interface{}
	params = append(params, taskuuid.String(), "created")
	if _, err := h.RepositoryWorker.Insert(consts.InsertNewTask, params); err != nil {
		return err
	}
	return ctx.JSON(202, taskuuid.String())
}

func (h *HandlersStruct) HandlerGetTaskStatus(ctx echo.Context) (Err error) {
	taskuuid := ctx.Param("taskId")
	var err error
	var params []interface{}
	var uuidSlice []models.TaskStatus
	params = append(params, taskuuid)
	if uuidSlice, err = h.RepositoryWorker.SelectTasksStatus(consts.SelectTaskStatus, params); err != nil {
		return err
	}
	if len(uuidSlice) == 0 {
		message := fmt.Sprintf("Задача с %s не найдена", taskuuid)
		return &echo.HTTPError{Code: 404, Message: message}
	}
	if len(uuidSlice) != 1 {
		return errors.New("several equal tasks")
	}
	return nil
}

func (h *HandlersStruct) HandleGetTaskFinished(ctx echo.Context) (Err error) {
	taskuuid := ctx.Param("taskId")
	var err error
	var params []interface{}
	var uuidSlice []models.TaskStatus
	params = append(params, taskuuid)
	if uuidSlice, err = h.RepositoryWorker.SelectTasksStatus(consts.SelectTaskStatus, params); err != nil {
		return err
	}
	if len(uuidSlice) == 0 {
		message := fmt.Sprintf("Задача с %s не найдена", taskuuid)
		return &echo.HTTPError{Code: 404, Message: message}
	}
	if len(uuidSlice) != 1 {
		return errors.New("several equal tasks")
	}
	return nil
}
