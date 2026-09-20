package handler

import (
	"net/http"
	"strconv"
	"task-api/internal/model"
	"task-api/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Handler *service.TaskService
}

func NewTaskHandler(handler *service.TaskService) *TaskHandler {
	return &TaskHandler{
		Handler: handler,
	}
}

func (h *TaskHandler) GetAllTask(ctx *gin.Context) {
	id , id_err := strconv.Atoi(ctx.Param("id"))

	if id_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"error":"Invalid ID",
		})
		return
	}

	tasks ,task_err := h.Handler.GetAll(id)

	if task_err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error" : "database error",
		})
		return
	}

	ctx.JSON(http.StatusOK,tasks)
}

func (h *TaskHandler) PostTask(ctx *gin.Context) {
	var task_in model.CreateTaskRequest

	if json_err := ctx.ShouldBindJSON(&task_in) ; json_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"error":"Invalid JSON",
		})
		return
	}

	task,task_err := h.Handler.PostTask(task_in)

	if task_err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"database error",
		})
		return
	}

	ctx.JSON(http.StatusCreated,task)
}

func (h *TaskHandler) UpdateTask(ctx *gin.Context) {
	var task_in model.UpdateTaskRequest

	if json_err := ctx.ShouldBindJSON(&task_in) ; json_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"error":"Invalid JSON",
		})
		return
	}

	task , task_err := h.Handler.UpdateTask(task_in)

	if task_err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error" :"database error",
		}) 
		return
	}

	if task == nil {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error" :"User not found",
		}) 
		return
	}

	ctx.JSON(http.StatusOK,task)
}

func (h * TaskHandler) DeleteTask(ctx *gin.Context) {
	id , id_err := strconv.Atoi(ctx.Param("id"))

	if id_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"error":"Invalid id",
		})
		return
	}

	row_affected , err := h.Handler.DeleteTask(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error" :"database error",
		}) 
		return
	}

	if row_affected == 0 {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error" :"User not found",
		}) 
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"message":"task deleted",
	})
}



