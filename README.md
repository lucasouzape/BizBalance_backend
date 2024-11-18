# BizBalance_backend
BizBalance Backend
Este é o backend do projeto BizBalance, uma aplicação para gerenciar finanças e fluxo de caixa, com funcionalidades específicas para a gestão de produtos da Sonho de Mel.

Funcionalidades
Cadastro, atualização e remoção de produtos como Pão de Mel.
Gestão de transações financeiras.
Integração com banco de dados PostgreSQL.
Suporte a testes de rotas via Postman.

Tecnologias Utilizadas
Linguagem de Programação: Go (Golang)
Banco de Dados: PostgreSQL
Containerização: Docker e Docker Compose
Ferramentas de Teste: Postman
Gerenciamento de Dependências: Go Modules

Pré-requisitos
Certifique-se de ter os seguintes itens instalados no sistema:

Go (Golang) - Versão >= 1.19
Docker - Versão >= 20.10
Postman (opcional, para testar a API)
_____________________________________________Configuração e Execução__________________________
Passo 1: Clone o Repositório
bash
Copy code
git clone git@github.com:lucasouzape/BizBalance_backend.git
cd BizBalance_backend
Passo 2: Configure as Variáveis de Ambiente
Copie o arquivo .env_example e configure as credenciais do banco de dados:

bash
Copy code
cp .env_example .env
Exemplo de conteúdo do .env:

plaintext
Copy code
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=bizbalance
PORT=8080


Passo 3: Inicie o Banco de Dados com Docker
bash
Copy code
docker-compose up -d

Passo 4: Instale as Dependências
bash
Copy code
go mod tidy

Passo 5: Rode o Backend
bash
Copy code
go run main.go
Passo 6: Acesse a API
A API estará disponível em: http://localhost:8080

Testes
Usando o Postman
Importe o arquivo SonhoDeMel.postman_collection.json no Postman.
Teste as rotas disponíveis, como:
GET /pao_de_mel: Lista produtos.
POST /pao_de_mel: Adiciona um produto.