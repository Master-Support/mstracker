package eventos

import (
	"net/http"

	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetObjetos(ctx *gin.Context) {
	var envios []models.StatusObjeto

	query := "SELECT * FROM status_objetos"
	result := h.DB.Raw(query).Scan(&envios)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, envios)
}

//WHERE status_objeto = 'Objeto entregue ao destinatário'

//WHERE status_objeto = 'Objeto em transferência - por favor aguarde'
