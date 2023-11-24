package eventos

import (
	"net/http"

	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetObjetos(ctx *gin.Context) {
	var envios []models.StatusObjeto  // Renomeie a variável para 'envios'

	if result := h.DB.Find(&envios); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"envios": envios})  // Use "envios" como chave no JSON
}