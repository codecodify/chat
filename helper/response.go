package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
	if data != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   data,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "000",
		})
	}

}

func Error(ctx *gin.Context, msg error) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "error",
		"msg":    msg.Error(),
	})
	ctx.Abort()
}
