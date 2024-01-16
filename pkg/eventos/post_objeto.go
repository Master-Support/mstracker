package eventos

import (
	rastroapi "example.com/mstracker/api_correios/rastro_api"
	"example.com/mstracker/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Codigo struct {
	CodigoObjeto *string `json:"codigoObjeto"`
	NomeObjeto   string  `json:"nomeObjeto"`
}

func (h handler) PostObjeto(ctx *gin.Context) {
	codigo := Codigo{}

	if err := ctx.BindJSON(&codigo); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := rastroapi.ObterStatus(*codigo.CodigoObjeto)
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, "O objeto nao esta cadastrado na base de dados dos correios!")
		return
	}

	var s models.StatusObjeto

	data, _ := time.Parse("2006-01-02T15:04:05", body.DataPrevistaDeEntrega)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	dataFormatada := data.Format("02/01/2006")

	s.NomeObjeto = codigo.NomeObjeto
	s.CodigoObjeto = *codigo.CodigoObjeto
	s.DataPrevistaDeEntrega = dataFormatada
	s.StatusObjeto = body.StatusObjeto
	s.Localizacao = body.Localizacao

	if result := h.DB.Create(&s); result.Error != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	h.DB.Updates(&s)

	ctx.JSON(http.StatusCreated, &s)
}
