package eventos

import (
	"net/http"

	rastroapi "example.com/mstracker/api_correios/rastro_api"
	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

type CodigoParaRealziarUpdate struct {
	CodigoObjeto *string `json:"codigoObjeto"`
}

func (h handler) UpdateStatus(ctx *gin.Context) {
	codigo := CodigoParaRealziarUpdate{}

	if err := ctx.BindJSON(&codigo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := rastroapi.ObterStatus(*codigo.CodigoObjeto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	id := ctx.Param(body.CodigoObjeto)
	var s models.StatusObjeto

	if result := h.DB.First(&s, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	s.CodigoObjeto = body.CodigoObjeto
	s.NomeObjeto = "TESTANDO UPDATE NOVAMENTE"
	s.DataPrevistaDeEntrega = body.DataPrevistaDeEntrega
	s.StatusObjeto = body.StatusObjeto
	s.Localizacao = body.Localizacao

	h.DB.Save(&s)

	ctx.JSON(http.StatusOK, &s)

}

// codigo := CodigoUpdate{}

// // body := UpdateStatusRequestBody{}

// // if err := ctx.BindJSON(&body); err != nil {
// //     ctx.AbortWithError(http.StatusBadRequest, err)
// //     return
// // }

// body, err := rastroapi.ObterStatus(*codigo.CodigoObjeto)
// if err != nil {
// 	ctx.AbortWithError(http.StatusInternalServerError, err)
// 	return
// }
// fmt.Println(body)

// id := ctx.Param("codigo_objeto")

// var updateStatus models.StatusObjeto

// if result := h.DB.First(&updateStatus, id); result.Error != nil {
// 	ctx.AbortWithError(http.StatusNotFound, result.Error)
// 	return
// }

// updateStatus.StatusObjeto = "TESTANDO"
// updateStatus.Localizacao = "TESTANDO"

// h.DB.Save(&updateStatus)

// ctx.JSON(http.StatusOK, &updateStatus)
