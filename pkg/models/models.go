package models

type TaskResponse struct {
	UUID        string `json:"uuid"`
	CreatedTime string `json:"timestamp"`
	Status      string `json:"status"`
}

type TaskStatus struct {
	CreatedTime string `json:"timestamp"`
	Status      string `json:"status"`
}