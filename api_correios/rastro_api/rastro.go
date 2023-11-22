package rastroapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	tokenapi "example.com/mstracker/api_correios/token_api"
	"example.com/mstracker/models"
)

type Response struct {
	Objetos       []Objeto  `json:"objetos"`
}

type Objeto struct {
	CodObjeto    string     `json:"codObjeto"`
	DtPrevista   string     `json:"dtPrevista"`
	Contrato     string     `json:"contrato"`
	Largura      int        `json:"largura"`
	Comprimento  int        `json:"comprimento"`
	Altura       int        `json:"altura"`
	Peso         float64    `json:"peso"`
	Formato      string     `json:"formato"`
	Modalidade   string     `json:"modalidade"`
	Eventos      []Evento   `json:"eventos"`
}

type TipoPostal struct {
	Sigla      string `json:"sigla"`
	Descricao  string `json:"descricao"`
	Categoria  string `json:"categoria"`
}

type Evento struct {
	Codigo         string    `json:"codigo"`
	Tipo           string    `json:"tipo"`
	DtHrCriado     string    `json:"dtHrCriado"`
	Descricao      string    `json:"descricao"`
	Unidade        Unidade   `json:"unidade"`
	UnidadeDestino Unidade   `json:"unidadeDestino"`
}

type Unidade struct {
	Endereco struct {
		Cidade string `json:"cidade"`
		UF     string `json:"uf"`
	} `json:"endereco"`
}

func ObterStatus(codigoObjeto string) (models.StatusObjeto, error) {
	url := "https://api.correios.com.br/srorastro/v1/objetos?codigosObjetos=" + codigoObjeto + "&resultado=U"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.StatusObjeto{}, err
	}

	// Adiciona o token de acesso ao cabeçalho "Authorization"
	token, _ := tokenapi.GenerateToken()

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	// Cria um cliente HTTP e envia a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.StatusObjeto{}, err
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.StatusObjeto{},  err
	}

	// Decodificar o JSON da resposta
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.StatusObjeto{}, err
	}

	// Verifica se há eventos antes de acessar os dados
	if len(response.Objetos) == 0 || len(response.Objetos[0].Eventos) == 0 {
		return models.StatusObjeto{}, fmt.Errorf("Nenhum evento encontrado para o objeto %s", codigoObjeto)
	}

	// Cria a instância da struct Status
	status := models.StatusObjeto{
		CodigoObjeto:            codigoObjeto,                            
		DataPrevistaDeEntrega: response.Objetos[0].DtPrevista,         
		StatusObjeto:          response.Objetos[0].Eventos[0].Descricao,
		Localizacao:           response.Objetos[0].Eventos[0].Unidade.Endereco.Cidade,
	}


	return status, nil
}
