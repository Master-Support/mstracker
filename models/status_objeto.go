package models

type StatusObjeto struct {
	CodigoObjeto          string `json:"codigoObjeto" gorm:"primaryKey"`
	NomeObjeto            string `json:"nomeObjeto"`
	DataPrevistaDeEntrega string `json:"dataPrevistaDeEntrega"`
	StatusObjeto          string `json:"statusObjeto"`
	Localizacao           string `json:"localizacao"`
}
