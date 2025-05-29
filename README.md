# Challenge Clean Architecture - Desafio: Listagem de Orders

## Descrição
Implementar listagem de orders utilizando REST, gRPC e GraphQL. Banco de dados deve ser provisionado via Docker. 

---

## Passos para execução do desafio

### 1. Setup do ambiente
- [ ] Criar `Dockerfile` da aplicação
- [ ] Criar `docker-compose.yaml` com:
  - Banco de dados (ex: PostgreSQL)
  - Serviço da aplicação
- [ ] Garantir que `docker compose up` sobe tudo corretamente

---

### 2. Modelagem e migração
- [ ] Definir model `Order`
- [ ] Criar as migrations para a tabela de orders
- [ ] Configurar ferramenta de migrations (ex: golang-migrate)

---

### 3. Implementar camada de repositório
- [ ] Criar repositório para acesso ao banco de dados
- [ ] Implementar método para listar orders

---

### 4. Implementar camada de usecase
- [ ] Criar o usecase `ListOrders`
- [ ] Conectar o usecase com o repositório

---

### 5. Implementar API REST
- [ ] Criar rota `GET /order`
- [ ] Chamar o usecase `ListOrders` na rota

---

### 6. Implementar serviço gRPC
- [ ] Criar definição `.proto` para o serviço `ListOrders`
- [ ] Gerar código com `protoc`
- [ ] Implementar handler gRPC chamando o usecase

---

### 7. Implementar query GraphQL
- [ ] Configurar servidor GraphQL
- [ ] Criar query `listOrders`
- [ ] Conectar com o usecase `ListOrders`

---

### 8. Arquivo de requests
- [ ] Criar `api.http` com exemplos de:
  - Requisição para criar order (se necessário)
  - Requisição para listar orders via REST

---

### 9. Documentação
- [ ] Criar ou completar `README.md` com:
  - Como rodar a aplicação
  - Portas de cada serviço (REST, gRPC, GraphQL)
  - Como testar os endpoints

---

## Observações
- REST: Porta 8080
- gRPC: Porta 50051
- GraphQL: Porta 8081 (ou ajustar conforme necessário)

