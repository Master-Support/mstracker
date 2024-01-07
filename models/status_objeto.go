package models

import "gorm.io/gorm"

type StatusObjeto struct {
	gorm.Model
	IdObjeto              uint64 `json:"idObjeto" gorm:"primaryKey"`
	CodigoObjeto          string `json:"codigoObjeto" gorm:"unique;not null"`
	NomeObjeto            string `json:"nomeObjeto" gorm:"not null"`
	DataPrevistaDeEntrega string `json:"dataPrevistaDeEntrega"`
	StatusObjeto          string `json:"statusObjeto"`
	Localizacao           string `json:"localizacao"`
}
