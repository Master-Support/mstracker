package rotas

import (
	"github.com/gin-gonic/gin"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI    string
	Metodo string
	Funcao func(*gin.Context)
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *gin.Engine) *gin.Engine {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		switch rota.Metodo {
		case "GET":
			r.GET(rota.URI, rota.Funcao)
		case "POST":
			r.POST(rota.URI, rota.Funcao)
		case "DELETE":
			r.DELETE(rota.URI, rota.Funcao)
		}
	}

	return r
}
