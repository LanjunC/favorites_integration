package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CtxJson(ctx *gin.Context, statusCode int) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode" : statusCode,
		"statusMsg" : GetStatusMsg(statusCode),
	})
}

func CtxJsonOfData(ctx *gin.Context,statusCode int, key string, value interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode" : statusCode,
		"statusMsg" : GetStatusMsg(statusCode),
		key : value,
	})
}
