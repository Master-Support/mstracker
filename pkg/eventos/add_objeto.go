package eventos

import (
	"net/http"
	"time"
	rastroapi "example.com/mstracker/api_correios/rastro_api"
	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
)

type Codigo struct {
	CodigoObjeto *string `json:"codigoObjeto"`
	NomeObjeto string `json:"nomeObjeto"`
}

func (h handler) AddStatus(ctx *gin.Context) {
	codigo := Codigo{}

	if err := ctx.BindJSON(&codigo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := rastroapi.ObterStatus(*codigo.CodigoObjeto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var s models.Status

	data, _ := time.Parse("2006-01-02T15:04:05", body.DataPrevistaDeEntrega)

	dataFormatada := data.Format("02/01/2006")

	s.NomeObjeto = codigo.NomeObjeto
	s.CodigoObjeto = *codigo.CodigoObjeto
	s.DataPrevistaDeEntrega = dataFormatada
	s.StatusObjeto = body.StatusObjeto
	s.Localizacao = body.Localizacao

	if result := h.DB.Create(&s); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &s)
}