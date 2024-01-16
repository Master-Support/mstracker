package eventos

import (
	"net/http"

	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetObjetos(ctx *gin.Context) {
	var envios []models.StatusObjeto
	var objetos_entregues []models.StatusObjeto

	query := "SELECT * FROM status_objetos WHERE status_objeto = 'Objeto entregue ao destinatário'"
	result := h.DB.Raw(query).Scan(&objetos_entregues)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"envios": envios, "objetos_entregues": objetos_entregues})
}
