# Goportunities

## Visão geral

Goportunities é uma API REST em Go para gerenciar vagas de emprego. O projeto usa o framework Gin para rotas HTTP, GORM com SQLite para armazenamento local e Swaggo para gerar documentação Swagger.

## Arquitetura do projeto

A organização do código segue uma estrutura simples e modular:

- `main.go` — inicializa logger, configurações e o roteador.
- `config/` — contém inicialização do banco de dados SQLite, logger e configurações gerais.
- `router/` — define rotas da API e a rota do Swagger.
- `handler/` — implementa os handlers HTTP e tipos de request/response.
- `schemas/` — define o modelo de dados `Opening` e structs de resposta.
- `db/` — pasta onde o banco SQLite local é criado (`db/main.db`).
- `docs/` — arquivos gerados pelo `swag init` para documentação Swagger.

## Tecnologias usadas

- Go 1.26
- Gin Web Framework (`github.com/gin-gonic/gin`)
- GORM ORM (`gorm.io/gorm`)
- SQLite via Glebarez (`github.com/glebarez/sqlite`)
- Swagger / Swaggo (`github.com/swaggo/swag`, `github.com/swaggo/gin-swagger`, `github.com/swaggo/files`)

## Pré-requisitos

- Go instalado (versão suportada: 1.26.x)
- Git Bash, PowerShell ou terminal compatível
- `go` disponível no PATH

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/HudsonJR777/gopportunities.git
cd gopportunities/gopportunities
```

2. Instale dependências Go:

```bash
go mod download
```

3. Gere a documentação Swagger (opcional, mas recomendado):

```bash
swag init --generalInfo main.go --output docs
```

## Execução

Execute a aplicação com:

```bash
go run main.go
```

A API inicializa na porta configurada em `router/router.go`.

> Observação: o projeto atual está configurado para rodar em `http://localhost:8081`.

## Rotas disponíveis

A base path usada é `/api/v1`.

- `POST /api/v1/opening` — criar uma nova vaga
- `GET /api/v1/openings` — listar todas as vagas
- `GET /api/v1/opening/:id` — buscar vaga por ID
- `PUT /api/v1/opening/:id` — atualizar vaga por ID
- `DELETE /api/v1/opening/:id` — remover vaga por ID
- `GET /swagger/*any` — documentação Swagger UI

## Swagger

Após iniciar a API, acesse:

```text
http://localhost:8081/swagger/index.html
```

Se preferir, use o arquivo `docs/swagger.json` para ferramentas compatíveis com OpenAPI.

## Configuração do ambiente

O projeto não exige variáveis de ambiente especiais. O banco local é criado automaticamente em `./db/main.db`.

Se quiser mudar a porta ou o caminho do banco, edite:

- `router/router.go` — porta HTTP
- `config/sqlite.go` — caminho do arquivo SQLite

## Exemplos de request / response

### Criar vaga

```bash
curl -X POST http://localhost:8081/api/v1/opening \
  -H "Content-Type: application/json" \
  -d '{
    "role": "Backend Developer",
    "company": "Goportunities",
    "location": "Remote",
    "remote": true,
    "link": "https://example.com/job/123",
    "salary": 12000
  }'
```

Resposta:

```json
{
  "message": "operation from hendler: create-opening successfull",
  "data": {
    "id": 1,
    "created_at": "2026-05-08T00:00:00Z",
    "updated_at": "2026-05-08T00:00:00Z",
    "deleted_at": null,
    "role": "Backend Developer",
    "company": "Goportunities",
    "location": "Remote",
    "remote": true,
    "link": "https://example.com/job/123",
    "salary": 12000
  }
}
```

### Listar vagas

```bash
curl http://localhost:8081/api/v1/openings
```

Resposta:

```json
{
  "message": "operation from hendler: list-openings successfull",
  "data": [
    {
      "id": 1,
      "created_at": "2026-05-08T00:00:00Z",
      "updated_at": "2026-05-08T00:00:00Z",
      "deleted_at": null,
      "role": "Backend Developer",
      "company": "Goportunities",
      "location": "Remote",
      "remote": true,
      "link": "https://example.com/job/123",
      "salary": 12000
    }
  ]
}
```

### Buscar vaga por ID

```bash
curl http://localhost:8081/api/v1/opening/1
```

Resposta:

```json
{
  "message": "operation from hendler: show-opening successfull",
  "data": {
    "id": 1,
    "created_at": "2026-05-08T00:00:00Z",
    "updated_at": "2026-05-08T00:00:00Z",
    "deleted_at": null,
    "role": "Backend Developer",
    "company": "Goportunities",
    "location": "Remote",
    "remote": true,
    "link": "https://example.com/job/123",
    "salary": 12000
  }
}
```

### Atualizar vaga

```bash
curl -X PUT http://localhost:8081/api/v1/opening/1 \
  -H "Content-Type: application/json" \
  -d '{
    "location": "Sao Paulo",
    "salary": 13000
  }'
```

Resposta:

```json
{
  "message": "operation from hendler: update-opening successfull",
  "data": {
    "id": 1,
    "created_at": "2026-05-08T00:00:00Z",
    "updated_at": "2026-05-08T00:00:00Z",
    "deleted_at": null,
    "role": "Backend Developer",
    "company": "Goportunities",
    "location": "Sao Paulo",
    "remote": true,
    "link": "https://example.com/job/123",
    "salary": 13000
  }
}
```

### Excluir vaga

```bash
curl -X DELETE http://localhost:8081/api/v1/opening/1
```

Resposta:

```json
{
  "message": "operation from hendler: delete-opening successfull",
  "data": {
    "id": 1,
    "created_at": "2026-05-08T00:00:00Z",
    "updated_at": "2026-05-08T00:00:00:00Z",
    "deleted_at": null,
    "role": "Backend Developer",
    "company": "Goportunities",
    "location": "Sao Paulo",
    "remote": true,
    "link": "https://example.com/job/123",
    "salary": 13000
  }
}
```

## Observações

- A API salva os dados em `./db/main.db` automaticamente.
- Se os exemplos retornarem erro de validação, verifique o payload JSON e os campos obrigatórios.
- A documentação Swagger é gerada a partir dos comentários `// @...` nos handlers.

