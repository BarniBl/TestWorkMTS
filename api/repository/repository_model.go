package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
)

type RepositoryStruct struct {
	connectionString string
	DataBase         *sql.DB
}

type RepositoryInterface interface {

}
