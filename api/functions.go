package functions

import (
	"fmt"
	"github.com/BarniBl/TestWorkMTS/api/repository"
	"github.com/BarniBl/TestWorkMTS/pkg/consts"
	"time"
)

func DoWork(taskuuid string, worker repository.RepositoryInterface) {
	var params1 []interface{}
	params1 = append(params1,"running", taskuuid)
	if _, err := worker.Update(consts.UpdateTaskStatus, params1); err != nil {
		fmt.Println(err)
	}
	time.Sleep(2*time.Minute)
	var params2 []interface{}
	params2 = append(params2, "finished", taskuuid)
	if _, err := worker.Update(consts.UpdateTaskStatus, params2); err != nil {
		fmt.Println(err)
	}
}
