# Sistema de Ticket

Este projeto é um sistema de ticket desenvolvido em Go. Ele permite a criação, gerenciamento e uso de tickets para diferentes promoções e usuários.

## Estrutura do Projeto

A estrutura do projeto é organizada da seguinte forma:

```
.vscode/
    launch.json
cmd/
    main.go
config/
    db.go
docker-compose.yml
go.mod
go.sum
internal/
    domain/
        company/
            handler.go
            service.go
        promotion/
            handler.go
            service.go
        user/
            handler.go
            service.go
    infra/
        adapters/
            postgresql/
        ports/
            repo/
    pkg/
        common/
            errors.go
        entity/
            company/
            promotion/
            user/
            ...
        utils/
README.md
scripts/
    db/
        dbconfig.yml
        migrations/
        setup_db.sh
test/
    integration/
    testhelpers/
        testhelpers.go
    unit/
        service/
```

## Arquitetura do Sistema

### .vscode/

Contém configurações específicas do Visual Studio Code, como o arquivo `launch.json` para facilitar o debug.

### cmd/

Contém o ponto de entrada principal da aplicação, `main.go`.

### config/

Contém arquivos de configuração, como `db.go` que configura a conexão com o banco de dados.

### internal/

Contém a lógica interna do sistema, dividida em subpastas:

- **domain/**: Contém a lógica de negócios para cada domínio (empresa, promoção, usuário).
- **infra/**: Contém adaptadores e portas para infraestrutura, como repositórios de banco de dados.
- **pkg/**: Contém pacotes reutilizáveis, como entidades, utilitários e erros comuns.

### scripts/

Contém scripts para configuração e migração do banco de dados.

### test/

Contém testes de integração e unitários, além de helpers para testes.

## Funcionalidades

- **Criação de Empresas**: Permite a criação de empresas com nome e ID fiscal.
- **Criação de Promoções**: Permite a criação de promoções associadas a empresas.
- **Criação de Usuários**: Permite a criação de usuários com nome de usuário, email e senha.
- **Gerenciamento de Tickets**: Permite a criação e uso de tickets para promoções específicas.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação principal.
- **Fiber**: Framework web para Go.
- **Bun**: ORM para Go.
- **PostgreSQL**: Banco de dados relacional.
- **Docker**: Para containerização do banco de dados.

## Configuração do Ambiente

### Pré-requisitos

- Docker
- Go 1.23.4 ou superior

### Passos para Configuração

1. Clone o repositório:

   ```sh
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

2. Configure o banco de dados usando Docker:

   ```sh
   docker-compose up -d
   ```

3. Execute as migrações do banco de dados:

   ```sh
   ./scripts/db/setup_db.sh
   ```

4. Inicie o servidor:
   ```sh
   go run cmd/main.go
   ```

## Testes

Para executar os testes, use o comando:

```sh
go test ./...
```

## Endpoints

### Empresa

- **POST /company**: Cria uma nova empresa.
  - Request Body:
    ```json
    {
      "name": "Nome da Empresa",
      "tax_id": "123456789"
    }
    ```

### Promoção

- **POST /promotion**: Cria uma nova promoção.
  - Request Body:
    ```json
    {
      "company_id": "UUID da Empresa",
      "name": "Nome da Promoção",
      "text_message_in_progress": "Mensagem de progresso",
      "text_message_success": "Mensagem de sucesso",
      "start_date": "2021-01-01",
      "end_date": "2021-01-02",
      "qty_max_users": 100,
      "vouchers_per_user": 10
    }
    ```

### Usuário

- **POST /user**: Cria um novo usuário.
  - Request Body:
    ```json
    {
      "username": "Nome de Usuário",
      "email": "email@exemplo.com",
      "password": "senha"
    }
    ```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
