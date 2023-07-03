# Clean Arch

Aplicação desenvolvida para cumprir o desafio proposto no module de Clean Arch do curso Go Expert.

A aplicações expoe 3 APIS,

# Pré requisitos

- Docker 20+
- Docker Compose 2.17+
- Golang 1.20+

# Como executar

- Após clonar o projeto executar o comando `docker-compose up -d` para subir o container com Mysql.
- Entrar no diretório `cmd/order` e executar o comando `go run main.go wire_gen.go` para subir a aplicação.
- Após a aplicação subir é possivel testar seus serviços pelos endpoints REST, pelos serviços GRPC e por queries no GraphQL.
    - Por REST, executar os arquivos .http dentro do diretório api
    - Por GRPC deverá ser configurado algum client que suporte o arquivo de definição protobuf dentro do diretorio `internal/infra/grpc/protofiles/` para fazer as requisições nos serviços.
    - Por GraphQL, acessar `http://localhost:8080/query` e executar as seguintes queries no console que abrira no navegador
    ```
    mutation createOrder {
        createOrder(input: {id: "e", Price: 10.0, Tax: 1.0}) {
            id
            Price
            FinalPrice
        }
    }

    query listOrders {
        orders {
            id
            Price
            Tax
            FinalPrice
        }
    }
    ```