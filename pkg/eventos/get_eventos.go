package eventos

import (
	"net/http"

	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetEventos(ctx *gin.Context) {
	var s []models.Status

	if result := h.DB.Find(&s); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &s)
}
