package models

type Status struct {
	NomeObjeto string    `json:"nomeObjeto"`
	CodigoObjeto string    `json:"codigoObjeto"`
	DataPrevistaDeEntrega string `json:"dataPrevistaDeEntrega"`
	StatusObjeto string `json:"statusObjeto"`
	Localizacao string `json:"localizacao"`
}
