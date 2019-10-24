package delivery

import (
	"errors"
	"fmt"
	"github.com/BarniBl/TestWorkMTS/pkg/consts"
	"github.com/BarniBl/TestWorkMTS/pkg/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"time"
)

func (h *HandlersStruct) HandlerCreateTask(ctx echo.Context) (Err error) {
	taskuuid := uuid.NewV4()
	return ctx.JSON(202, taskuuid.String())
/*	var params []interface{}
	params = append(params, taskuuid.String(), "created", time.Now())
	if _, err := h.RepositoryWorker.Insert(consts.InsertNewTask, params); err != nil {
		return err
	}
	go functions.DoWork(taskuuid.String(), h.RepositoryWorker)
	return ctx.JSON(202, taskuuid.String())*/
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
	return ctx.JSON(200, uuidSlice[0])
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
	if uuidSlice[0].Status == "finished" {
		return ctx.JSON(200, uuidSlice[0])
	}
	currentStatus := uuidSlice[0].Status
	ticker := time.NewTicker(5 * time.Second)
	i := 0
	for range ticker.C {
		i++
		var err error
		var params []interface{}
		var uuidSlice []models.TaskStatus
		params = append(params, taskuuid)
		if uuidSlice, err = h.RepositoryWorker.SelectTasksStatus(consts.SelectTaskStatus, params); err != nil {
			ctx.Logger().Error(err)
		}
		if len(uuidSlice) == 0 {
			message := fmt.Sprintf("Задача с %s не найдена", taskuuid)
			ctx.Logger().Error(message)
		}
		if len(uuidSlice) != 1 {
			ctx.Logger().Error("several equal tasks")
		}
		if uuidSlice[0].Status == "finished" {
			currentStatus = uuidSlice[0].Status
			ticker.Stop()
			break
		}
		if i >= 60 {
			ticker.Stop()
			break
		}
	}
	if currentStatus != "finished" {
		return &echo.HTTPError{Code: 408, Message: "Ожидание превысило порог в 5 минут"}
	}
	uuidSlice[0].Status = "finished"
	return ctx.JSON(200, uuidSlice[0])
}
