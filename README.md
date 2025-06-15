# Challenge Clean Architecture - Desafio: Listagem de Orders

## Descrição
Implementar listagem de orders utilizando REST, gRPC e GraphQL. Banco de dados deve ser provisionado via Docker. 

# Desafio de Listagem de Orders

Este repositório contém a implementação da listagem de pedidos (orders) via REST, gRPC e GraphQL, utilizando Docker para orquestração dos serviços.


# EXECUÇÃO DO APP

## STEPS

1. Rode o `compose`
~~~shell
docker compose up --build
~~~

2. Testando o serviço `rest`

- POST
~~~shell
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": "123",
    "product_id": "456",
    "quantity": 2,
    "total_price": 99.90,
    "status": "pending"
  }'
~~~

- GET
~~~shell
curl -X GET http://localhost:8080/orders
~~~

3. Testando o `gRPC`

- Rode o comando do evans
~~~shell
docker compose exec grpc evans -r repl --host 0.0.0.0 --port 50051 --proto /proto/order_list.proto
~~~

- No envans rode os seguintes comandos
~~~shell
package pb
service OrderService
call ListOrders
~~~

4. Testando o `graphQl`

- Acesse a url `http://localhost:8081/`
- Siga exemplo de `query`
~~~shell
query queryOrders{
  orders{
    quantity,
    status,
    created_at,
    id,
    customer_id,
    product_id
  }
}
~~~

## Variáveis 

- Já definido no `.env`
~~~
REST_PORT=8080
GRPC_PORT=50051
GRAPHQL_PORT=8081
~~~

# Checklist
---

## ✅ Etapa 1: API REST (GET /order)

- [X] Criar estrutura base do projeto em Go
- [X] Configurar dependências (Go modules, etc.)
- [X] Criar camada de model e repository de Order
- [X] Criar service de listagem de orders
- [X] Criar handler REST com endpoint `GET /order`
- [X] Criar arquivo `api.http` com exemplo de requisição REST
- [X] Testar retorno da listagem via REST

---

## ✅ Etapa 2: Serviço gRPC (ListOrders)

- [X] Criar definição do serviço gRPC (proto file)
- [X] Gerar os arquivos gRPC (protobuf + Go)
- [X] Implementar o serviço `ListOrders` no servidor gRPC
- [X] Conectar o service gRPC com a camada de service/repository
- [X] Expor o serviço na porta definida
- [X] Adicionar testes ou client para validação local
  - Evans: https://github.com/ktr0731/evans

---

## ✅ Etapa 3: Query GraphQL (ListOrders)

- [X] Configurar o servidor GraphQL
  - Doc: https://gqlgen.com/
- [X] Criar schema com a query `listOrders`
  - `go run github.com/99designs/gqlgen generate`
- [X] Mapear resolver para a query usando a camada de service
- [X] Expor endpoint GraphQL
- [X] Testar listagem com playground ou cliente GraphQL

---

## ✅ Etapa 4: Banco de Dados e Docker

- [X] Criar Dockerfile da aplicação
- [X] Criar `docker-compose.yaml` com:
  - Serviço da aplicação
  - Banco de dados (sqlite)
- [X] Criar scripts de migração para a tabela `orders`
- [X] Garantir que `docker compose up`:
  - Sobe aplicação
  - Cria e migra o banco de dados
- [X] Documentar variáveis de ambiente e configuração do DB

---

## ✅ Etapa 5: Documentação

- [X] Criar `README.md` com:
  - Instruções de build e execução
  - Portas de cada serviço (REST, gRPC, GraphQL)
  - Exemplo de requisições (ou apontar para `api.http`)
- [X] Atualizar checklist conforme for completando cada etapa

