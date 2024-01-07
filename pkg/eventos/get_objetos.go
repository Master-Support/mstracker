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

	for _, envio := range envios {
		if envio.StatusObjeto != "Objeto entregue ao destinatário"{
			//quando um objeto já estiver sido entregue ele ira para outra categoria de objetos entregues.
			//caso ao contrario o objeto será exibido em envios
			return 
		}
	}


	ctx.JSON(http.StatusOK, gin.H{"envios": envios})  // Use "envios" como chave no JSON
}