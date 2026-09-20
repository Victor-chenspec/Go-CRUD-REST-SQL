package model

type Task struct {
	ID        int    `json:"id"`
	TITLE     string `json:"title"`
	COMPLETED bool   `json:"completed"`
	USERID    int    `json:"user_id"`
}

type CreateTaskRequest struct {
	TITLE  string `json:"title" binding:"required"`
	USERID int    `json:"user_id" binding:"required"`
}

type UpdateTaskRequest struct {
	ID        int    `json:"id" binding:"required"`
	TITLE     string `json:"title" binding:"required"`
	COMPLETED bool   `json:"completed"`
	USERID    int    `json:"user_id" binding:"required"`
}