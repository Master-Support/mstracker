package eventos

import (
	"net/http"

	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetObjeto(ctx *gin.Context) {
	id := ctx.Param("id")

	var s models.StatusObjeto

	if result := h.DB.First(&s, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &s)
}
