package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"greeting" : "hello",
	})
}
