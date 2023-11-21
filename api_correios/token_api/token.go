package tokenapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RespostaToken struct {
	Token string `json:"token"`
}

// GenerateToken gera um token para a aplicação.
func GenerateToken() (string, error) {
	url := "https://api.correios.com.br/token/v1/autentica/cartaopostagem"
	body := []byte(`{"numero": "0077126548"}`)

	username := os.Getenv("CORREIOS_USERNAME")
	password := os.Getenv("CORREIOS_PASSWORD")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("Erro ao criar requisição: %v", err)
	}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Erro ao ler o corpo da resposta: %v", err)
	}

	// Decodifica a resposta JSON para extrair o token
	var resposta RespostaToken
	err = json.Unmarshal(respBody, &resposta)
	if err != nil {
		return "", fmt.Errorf("Erro ao decodificar resposta JSON: %v", err)
	}

	// Retorna o token
	return resposta.Token, nil
}