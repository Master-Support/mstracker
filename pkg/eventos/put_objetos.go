package eventos

import (
	"fmt"
	"net/http"

	rastroapi "example.com/mstracker/api_correios/rastro_api"
	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

type UpdateStatusRequestBody struct {
	StatusObjeto          string `json:"statusObjeto"`
	Localizacao           string `json:"localizacao"`
}

type CodigoUpdate struct {
	CodigoObjeto *string `json:"codigoObjeto"`
}

func (h handler) UpdateStatus(ctx *gin.Context) {
	codigo := CodigoUpdate{}

    
    // body := UpdateStatusRequestBody{}

    // if err := ctx.BindJSON(&body); err != nil {
    //     ctx.AbortWithError(http.StatusBadRequest, err)
    //     return
    // }

	body, err := rastroapi.ObterStatus(*codigo.CodigoObjeto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(body)

	id := ctx.Param("id")

    var updateStatus models.StatusObjeto

    if result := h.DB.First(&updateStatus, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    updateStatus.StatusObjeto = "TESTANDO"
    updateStatus.Localizacao = "TESTANDO"

    h.DB.Save(&updateStatus)

    ctx.JSON(http.StatusOK, &updateStatus)
}