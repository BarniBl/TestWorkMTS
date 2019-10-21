package repository

import (
	"database/sql"
	"github.com/BarniBl/TestWorkMTS/pkg/models"
)

type RepositoryStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type RepositoryInterface interface {
	Insert(executeQuery string, params []interface{}) (string, error)
	Update(executeQuery string, params []interface{}) (int, error)
	SelectTasksStatus(executeQuery string, params []interface{}) ([]models.TaskStatus, error)
}
