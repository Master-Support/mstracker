package eventos

import (
	"github.com/gin-gonic/gin"
)

func (h handler) Options(ctx *gin.Context) {
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(200)
		return
	}
}
