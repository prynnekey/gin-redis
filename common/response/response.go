package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, data, msg)
}

func Fail(ctx *gin.Context, msg string) {
	Response(ctx, http.StatusBadRequest, nil, msg)
}
