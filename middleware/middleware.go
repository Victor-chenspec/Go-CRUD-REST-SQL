package middleware

import (
	"errors"
	"task-api/apperror"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return 
		}

		err := ctx.Errors.Last().Err

		var appErr *apperror.AppError

		if errors.As(err , &appErr) {
			ctx.JSON(appErr.Status,gin.H{
				"error":appErr.Message,
			})
			return 
		}

		ctx.JSON(500,gin.H{
			"error":"Internal server error",
		})
	}
}