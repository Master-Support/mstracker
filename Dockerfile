FROM golang:1.21.0

# Crie o diretório de trabalho e defina-o como diretório de trabalho
WORKDIR /app

# Copie os arquivos de dependências go.mod e go.sum
COPY go.mod go.sum ./

# Baixe os módulos Go
RUN go mod download

# Copie os arquivos Go
COPY *.go ./

RUN go build -o /mstracker

# Opcional: Expor a porta que o aplicativo vai escutar
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["/mstracker"]
