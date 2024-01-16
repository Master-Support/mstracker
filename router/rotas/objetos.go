package rotas

import (
	controllers "example.com/mstracker/controller"
)

var rotasUsuarios = []Rota{
	{
		URI:    "/api/v1/envio",
		Metodo: "POST",
		Funcao: controllers.CriarObjeto,
	},

	{
		URI:    "/api/v1/envio/objetos",
		Metodo: "GET",
		Funcao: controllers.BuscarObjetos,
	},

	{
		URI:    "/api/v1/envio/:objetoId",
		Metodo: "GET",
		Funcao: controllers.BuscarObjeto,
	},

	{
		URI:    "/api/v1/envio/:objetoId",
		Metodo: "POST",
		Funcao: controllers.AtualizarObjeto,
	},

	{
		URI:    "/api/v1/envio/:objetoId",
		Metodo: "DELETE",
		Funcao: controllers.DeletarObjeto,
	},
}
