package eventos

import (
	"net/http"

	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

// func (h handler) GetObjetos(ctx *gin.Context) {
// 	var envios []models.StatusObjeto
// 	var objetos_entrgues []models.StatusObjeto

// 	if result := h.DB.Find(&envios); result.Error != nil {
// 		ctx.AbortWithError(http.StatusNotFound, result.Error)
// 		return
// 	}

// 	for _, envio := range envios {
// 		if envio.StatusObjeto != "Objeto entregue ao destinatário"{
// 			objetos_entrgues = append(objetos_entrgues, envio)
// 			//quando um objeto já estiver sido entregue ele ira para outra categoria de objetos entregues.
// 			//caso ao contrario o objeto será exibido em envios
// 		}
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"envios": envios})
// }

func (h handler) GetObjetos(ctx *gin.Context) {
	var envios []models.StatusObjeto
	var objetos_entregues []models.StatusObjeto

	// Exemplo de consulta SQL usando Raw
	query := "SELECT * FROM status_objetos WHERE status_objeto = 'Objeto entregue ao destinatário'"
	result := h.DB.Raw(query).Scan(&objetos_entregues)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// Restante do código...

	ctx.JSON(http.StatusOK, gin.H{"envios": envios, "objetos_entregues": objetos_entregues})
}
