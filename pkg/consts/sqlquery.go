package consts

const (
	InsertNewTask    = "INSERT INTO testMTS.public.tasks (uuid, status, starttime) values ($1,$2,$3) RETURNING uuid"
	SelectTaskStatus = "SELECT task.status, task.starttime FROM testmts.public.tasks as task where uuid = $1"
	UpdateTaskStatus = "UPDATE testMTS.public.tasks SET status = $1 where uuid = $2"
)
