# Challenge Clean Architecture - Desafio: Listagem de Orders

## Descrição
Implementar listagem de orders utilizando REST, gRPC e GraphQL. Banco de dados deve ser provisionado via Docker. 

# Desafio de Listagem de Orders

Este repositório contém a implementação da listagem de pedidos (orders) via REST, gRPC e GraphQL, utilizando Docker para orquestração dos serviços.

---

## ✅ Etapa 1: API REST (GET /order)

- [X] Criar estrutura base do projeto em Go
- [X] Configurar dependências (Go modules, etc.)
- [X] Criar camada de model e repository de Order
- [X] Criar service de listagem de orders
- [ ] Criar handler REST com endpoint `GET /order`
- [ ] Criar arquivo `api.http` com exemplo de requisição REST
- [ ] Testar retorno da listagem via REST

---

## ✅ Etapa 2: Serviço gRPC (ListOrders)

- [ ] Criar definição do serviço gRPC (proto file)
- [ ] Gerar os arquivos gRPC (protobuf + Go)
- [ ] Implementar o serviço `ListOrders` no servidor gRPC
- [ ] Conectar o service gRPC com a camada de service/repository
- [ ] Expor o serviço na porta definida
- [ ] Adicionar testes ou client para validação local

---

## ✅ Etapa 3: Query GraphQL (ListOrders)

- [ ] Configurar o servidor GraphQL
- [ ] Criar schema com a query `listOrders`
- [ ] Mapear resolver para a query usando a camada de service
- [ ] Expor endpoint GraphQL
- [ ] Testar listagem com playground ou cliente GraphQL

---

## ✅ Etapa 4: Banco de Dados e Docker

- [ ] Criar Dockerfile da aplicação
- [ ] Criar `docker-compose.yaml` com:
  - Serviço da aplicação
  - Banco de dados (sqlite)
- [ ] Criar scripts de migração para a tabela `orders`
- [ ] Garantir que `docker compose up`:
  - Sobe aplicação
  - Cria e migra o banco de dados
- [ ] Documentar variáveis de ambiente e configuração do DB

---

## ✅ Etapa 5: Documentação

- [ ] Criar `README.md` com:
  - Instruções de build e execução
  - Portas de cada serviço (REST, gRPC, GraphQL)
  - Exemplo de requisições (ou apontar para `api.http`)
- [ ] Atualizar checklist conforme for completando cada etapa

---

## Observações
- REST: Porta 8080
- gRPC: Porta 50051
- GraphQL: Porta 8081 (ou ajustar conforme necessário)

