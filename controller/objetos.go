package controllers

import "github.com/gin-gonic/gin"

// CriarObjeto insere um objeto no banco de dados
func CriarObjeto(c *gin.Context) {
	c.String(200, "Criando Objeto")
}

// BuscarObjetos procura por todos os objetos no banco de dados
func BuscarObjetos(c *gin.Context) {
	c.String(200, "Buscando todos os objetos")
}

// BuscarObjeto procura por um objeto no banco de dados
func BuscarObjeto(c *gin.Context) {
	c.String(200, "Buscando um Objeto")
}

// AtualizarObjeto atualiza informações sobre um objeto no banco de dados
func AtualizarObjeto(c *gin.Context) {
	c.String(200, "Atualizando Objeto")
}

// DeletarObjeto deleta um objeto no banco de dados
func DeletarObjeto(c *gin.Context) {
	c.String(200, "Deletando Objeto")
}
