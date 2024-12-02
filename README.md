```markdown
# BizBalance Backend

O BizBalance Backend é uma aplicação desenvolvida em Go (Golang) para gerenciar finanças e fluxo de caixa, com funcionalidades específicas para a gestão de produtos da Sonho de Mel. Ele oferece integração com o banco de dados PostgreSQL e suporte a cálculos customizados para produtos. Este projeto foi desenvolvido com foco em modernidade e escalabilidade, sendo configurável para ambientes de desenvolvimento e produção.

## Funcionalidades

- Cadastro, atualização e remoção de produtos como Pão de Mel.
- Gestão de transações financeiras.
- Integração com banco de dados PostgreSQL.
- Suporte a testes de rotas via Postman.
- Endpoint customizado `/calculate` para cálculos financeiros.

## Tecnologias Utilizadas

- **Linguagem de Programação**: Go (Golang)
- **Banco de Dados**: PostgreSQL
- **Containerização**: Docker e Docker Compose
- **Ferramentas de Teste**: Postman
- **Gerenciamento de Dependências**: Go Modules

## Pré-requisitos

Certifique-se de ter os seguintes itens instalados no sistema:
- **Go (Golang)** - Versão >= 1.19
- **Docker** - Versão >= 20.10
- **Postman** (opcional, para testar a API)

## Configuração e Execução

### Passo 1: Clone o Repositório
```bash
git clone git@github.com:lucasouzape/BizBalance_backend.git
cd BizBalance_backend
```

### Passo 2: Configure as Variáveis de Ambiente
Copie o arquivo `.env_example` e configure as credenciais do banco de dados:
```bash
cp .env_example .env
```

Exemplo de conteúdo do `.env`:
```plaintext
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=bizbalance
SERVER_PORT=8080
```

### Passo 3: Inicie o Banco de Dados com Docker
Certifique-se de que o Docker está instalado e execute o comando para iniciar o banco de dados:
```bash
docker-compose up -d
```

### Passo 4: Instale as Dependências
Use o comando abaixo para instalar as dependências do projeto:
```bash
go mod tidy
```

### Passo 5: Rode o Backend
Execute o backend com o seguinte comando:
```bash
go run main.go
```

### Passo 6: Acesse a API
A API estará disponível em:
```
http://localhost:8080
```

## Endpoints Disponíveis

### Produtos Pão de Mel
- **GET /api/pao_de_mel**
  - Lista todos os produtos cadastrados.

- **POST /api/pao_de_mel/add**
  - Adiciona um novo produto.
  - Exemplo de payload:
    ```json
    {
      "sabor": "Chocolate",
      "quantidade": 50,
      "preco_custo": 3.50,
      "preco_venda": 5.00
    }
    ```

### Cálculos de Retorno
- **POST /api/calculate**
  - Realiza o cálculo de retorno baseado nos dados enviados.
  - Exemplo de payload:
    ```json
    {
      "quantidade_vendida": 10,
      "preco_custo": 3.50,
      "preco_venda": 5.00
    }
    ```
  - Resposta:
    ```json
    {
      "retorno": 15.00
    }
    ```

## Testes com Postman

1. Importe o arquivo `SonhoDeMel.postman_collection.json` no Postman.
2. Teste os seguintes endpoints:
   - **GET** `/api/pao_de_mel`: Lista todos os produtos.
   - **POST** `/api/pao_de_mel/add`: Adiciona um novo produto.
   - **POST** `/api/calculate`: Realiza o cálculo de retorno.

## Atualizações Recentes

- **Adição do endpoint `/calculate`**:
  - Permite o cálculo customizado de retornos.
  - Totalmente integrado com o controlador `ItemController`.
- **Melhorias gerais na configuração do backend**:
  - Configuração para rodar na porta especificada no arquivo `.env` ou na padrão (8080).

## Observações

- Utilize o Postman ou outra ferramenta de teste de API para validar o funcionamento das rotas.
- Certifique-se de que o banco de dados está rodando antes de iniciar o backend.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir um pull request ou relatar problemas na página de issues do repositório.

## Licença

Este projeto está licenciado sob os termos da licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.
```