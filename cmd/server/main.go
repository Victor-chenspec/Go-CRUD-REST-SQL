package main

import (
	"context"
	"log"
	"os"
	"task-api/internal/handler"
	"task-api/internal/repository"
	"task-api/internal/service"
	"task-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	//Env
	env_err := godotenv.Load()

	if env_err != nil {
		log.Fatal(env_err)
	}

	//Database
	db, db_err := pgx.Connect(context.Background(),os.Getenv("SQL"))

	if db_err != nil {
		log.Fatal(db_err)
	}

	defer db.Close(context.Background())

	//User
	UserRepository := repository.NewUserRepository(db)
	UserService := service.NewUserService(UserRepository)
	UserHandler := handler.NewUserHandler(UserService)

	//Task
	TaskRepository := repository.NewTaskRepository(db)
	TaskService := service.NewTaskService(TaskRepository)
	TaskHandler := handler.NewTaskHandler(TaskService)

	router := gin.Default()

	//Middleware
	router.Use(middleware.ErrorHandler())

	//User Route
	router.GET("/user",UserHandler.GetAllUsers)
	router.GET("/user/:id",UserHandler.GetOneUsers)
	router.POST("/user",UserHandler.PostUser)
	router.DELETE("/user/:id",UserHandler.DeleteUser)
	router.PUT("/user",UserHandler.UpdateUser)

	//Task Route
	router.GET("/task/:id",TaskHandler.GetAllTask)
	router.POST("/task",TaskHandler.PostTask)
	router.PUT("/task",TaskHandler.UpdateTask)
	router.DELETE("/task/:id",TaskHandler.DeleteTask)

	router.Run(os.Getenv("PORT"))
}