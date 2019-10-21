package repository

import (
	"database/sql"
	"github.com/BarniBl/TestWorkMTS/pkg/models"
	_ "github.com/lib/pq"
)

var ConnStr string = "user=postgres password=7396 dbname=testmts sslmode=disable"

func (RS *RepositoryStruct) NewDataBaseWorker() error {
	RS.connectionString = ConnStr
	var err error = nil

	RS.DataBase, err = sql.Open("postgres", ConnStr)
	if err != nil {
		return err
	}
	RS.DataBase.SetMaxOpenConns(10)
	err = RS.DataBase.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (RS *RepositoryStruct) Insert(executeQuery string, params []interface{}) (string, error) {
	var id string
	err := RS.DataBase.QueryRow(executeQuery, params...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (RS *RepositoryStruct) Update(executeQuery string, params []interface{}) (int, error) {
	result, err := RS.DataBase.Exec(executeQuery, params...)
	if err != nil {
		return 0, err
	}
	rowsEdit, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsEdit), nil
}

func (RS *RepositoryStruct) SelectTasksStatus(executeQuery string, params []interface{}) (Sl []models.TaskStatus, Err error) {
	taskStatusSlice := make([]models.TaskStatus, 0)
	rows, err := RS.DataBase.Query(executeQuery, params...)
	if err != nil {
		return taskStatusSlice, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			Err = err
		}
	}()
	for rows.Next() {
		taskStatus := models.TaskStatus{}
		err := rows.Scan(&taskStatus.Status, &taskStatus.CreatedTime)
		if err != nil {
			return taskStatusSlice, err
		}
		taskStatusSlice = append(taskStatusSlice, taskStatus)
	}
	return taskStatusSlice, nil
}
