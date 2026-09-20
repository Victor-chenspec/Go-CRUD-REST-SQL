package handler

import (
	"net/http"
	"strconv"
	"task-api/internal/model"
	"task-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Handler *service.UserService
}

func NewUserHandler(handler *service.UserService) *UserHandler {
	return &UserHandler{
		Handler: handler,
	}
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users,err := h.Handler.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Database error",
		})
		return
	}

	ctx.JSON(http.StatusOK,users)
}

func (h *UserHandler) GetOneUsers(ctx *gin.Context) {
	id , id_err := strconv.Atoi(ctx.Param("id"))

	if id_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"error":"Invalid Id",
		})
		return
	}

	user , user_err := h.Handler.GetOneUsers(id)

	if user_err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Database error",
		})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error":"User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK,user)
}
func (h *UserHandler) PostUser(ctx *gin.Context) {
	var user_in model.CreateUserRequest

	if in_err := ctx.ShouldBindJSON(&user_in) ; in_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"error":"Invalid JSON",
		})
		return
	}

	user , user_err := h.Handler.PostUser(user_in)

	if user_err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Database error",
		})
		return
	}

	ctx.JSON(http.StatusCreated,user)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id , id_err := strconv.Atoi(ctx.Param("id"))

	if id_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"error":"Invalid ID",
		})
		return
	}

	row_affected , err := h.Handler.DeleteUser(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Database error",
		})
		return
	}

	if row_affected == 0 {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error":"User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"message":"Deleted user",
	})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var user_in model.UpdateUserRequest

	if in_err := ctx.ShouldBindJSON(&user_in) ; in_err != nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"error":"Invalid JSON",
		})
		return
	}

	user , err := h.Handler.UpdateUser(user_in)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Database error",
		})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error":"User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK,user)
}