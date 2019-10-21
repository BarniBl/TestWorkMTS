package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var ConnStr string = "user=postgres password=7396 dbname=sunrise_db sslmode=disable"

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
