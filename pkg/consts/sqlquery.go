package consts

const (
	InsertNewTask    = "INSERT INTO testAPI.tasks (uuid, status, time) values ($1,$2,$3) RETURNING id"
	SelectTaskStatus = "SELECT FROM testAPI.tasks tasks.status, tasks.time  where testAPI.uuid = $1"
)
