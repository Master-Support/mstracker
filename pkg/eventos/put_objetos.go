package eventos

import (
	rastroapi "example.com/mstracker/api_correios/rastro_api"
	"example.com/mstracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) UpdateStatus(ctx *gin.Context) {
	var codigos []string

	query := "SELECT codigo_objeto FROM status_objetos WHERE status_objeto = 'Objeto entregue ao destinatário'"
	result := h.DB.Raw(query).Scan(&codigos)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for _, codigoObjeto := range codigos {
		body, err := rastroapi.ObterStatus(codigoObjeto)

		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var s models.StatusObjeto
		if result := h.DB.First(&s, "codigo_objeto = ?", body.CodigoObjeto); result.Error != nil {
			ctx.AbortWithError(http.StatusNotFound, result.Error)
			return
		}

		s.CodigoObjeto = body.CodigoObjeto
		s.NomeObjeto = body.NomeObjeto
		s.DataPrevistaDeEntrega = body.DataPrevistaDeEntrega
		s.StatusObjeto = body.StatusObjeto
		s.Localizacao = body.Localizacao

		h.DB.Save(&s)

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Atualização realizada com sucesso"})
}
