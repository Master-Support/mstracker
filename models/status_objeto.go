package models

type StatusObjeto struct {
	IdObjeto              uint64 `json:"idObjeto" gorm:"primaryKey"`
	CodigoObjeto          string `json:"codigoObjeto" gorm:"unique"`
	NomeObjeto            string `json:"nomeObjeto"`
	DataPrevistaDeEntrega string `json:"dataPrevistaDeEntrega"`
	StatusObjeto          string `json:"statusObjeto"`
	Localizacao           string `json:"localizacao"`
}
