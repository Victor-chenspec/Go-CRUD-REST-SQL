package model

type User struct {
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	EMAIL string `json:"email"`
}

type CreateUserRequest struct {
	NAME  string `json:"name" binding:"required"`
	EMAIL string `json:"email" binding:"required,email"`
}

type UpdateUserRequest struct {
	ID    int    `json:"id" binding:"required"`
	NAME  string `json:"name" binding:"required"`
	EMAIL string `json:"email" binding:"required,email"`
}