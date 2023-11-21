package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	NomeObjeto            string `json:"nomeObjeto"`
	CodigoObjeto          string `json:"codigoObjeto"`
	DataPrevistaDeEntrega string `json:"dataPrevistaDeEntrega"`
	StatusObjeto          string `json:"statusObjeto"`
	Localizacao           string `json:"localizacao"`
}
