# Use uma imagem base do Go
FROM golang:1.22.2

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie o código-fonte para o diretório de trabalho
COPY . .

# Baixe as dependências e compile o aplicativo
RUN go mod tidy
RUN go build -o main .

# Exponha a porta em que o servidor vai rodar
EXPOSE 8080

# Execute o aplicativo
CMD ["./main.go"]