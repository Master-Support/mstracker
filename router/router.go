package router

import (
	"example.com/mstracker/router/rotas"
	"github.com/gin-gonic/gin"
)

// Gerar retorna um router com as rotas configuradas
func Gerar() *gin.Engine {
	r := gin.Default()

	return rotas.Configurar(r)
}
