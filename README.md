
# Battle Game API 🎮

![Go](https://img.shields.io/badge/Go-1.23.2-blue)
![MySQL](https://img.shields.io/badge/MySQL-5.7-blue)
![JWT](https://img.shields.io/badge/JWT-Auth-blue)

## Descrição 📝
Battle Game API é uma aplicação backend desenvolvida em Go que permite aos usuários registrar-se, adicionar personagens e participar de batalhas.

## Endpoints

### Registro de Usuário 🚀
- **URL**: `/user/register`
- **Método**: `POST`
- **Corpo da Requisição**:
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Resposta de Sucesso**:
  - **Código**: `200 OK`
  - **Conteúdo**: `"User created successfully"`

### Adicionar Personagens 🧙‍♂️
- **URL**: `/users/{userID}/characters`
- **Método**: `POST`
- **Parâmetros de URL**:
  - `userID`: ID do usuário
- **Corpo da Requisição**:
  ```json
  [
    {
      "name": "string",
      "health": "int",
      "strength": "int",
      "type": "string",
      "species": "string",
      "species": "string",
      "special_attack": "string",
      "special_defense": "string",
      "defense": "string",
      "weakness": "string"
    }
  ]
  ```
- **Resposta de Sucesso**:
  - **Código**: `201 CREATED`
  - **Conteúdo**: `"Character assigned successfully"`

### Editar Personagens 📝
- **URL**: `/users/{userID}/characters/edit`
- **Método**: `POST`
- **Parâmetros de URL**:
  - `userID`: ID do usuário
- **Corpo da Requisição**:
  ```json
  [
    {
      "name": "string",
      "health": "int",
      "strength": "int",
      "type": "string",
      "species": "string",
      "special_attack": "string",
      "special_defense": "string",
      "defense": "string",
      "weakness": "string"
    }
  ]
  ```
- **Resposta de Sucesso**:
  - **Código**: `200 OK`
  - **Conteúdo**: `"User character edited successfully"`

### Batalha ⚔️
- **URL**: `/user/battle`
- **Método**: `POST`
- **Corpo da Requisição**:
  ```json
  {
    "user_id_1": "uint",
    "user_id_2": "uint"
  }
  ```
- **Resposta de Sucesso**:
  - **Código**: `200 OK`
  - **Conteúdo**: `{"result": "string", "character1": "string", "character2": "string"}`

### Listar Usuários 👥
- **URL**: `/users`
- **Método**: `GET`
- **Resposta de Sucesso**:
  - **Código**: `200 OK`
  - **Conteúdo**:
    ```json
    [
      {
        "id": "uint",
        "username": "string",
        "characters": [
          {
            "name": "string",
            "health": "int",
            "strength": "int",
            "type": "string",
            "species": "string",
            "special_attack": "string",
            "special_defense": "string",
            "defense": "string",
            "weakness": "string"
          }
        ]
      }
    ]
    ```

## Configuração do Ambiente 🔧
Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
```
DB_USER=fictitious_user
DB_PASSWORD=fictitious_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=fictitious_db
JWT_SECRET=fictitious_secret
```

## Como Executar 🚀
1. Clone o repositório:
   ```
   git clone https://github.com/seu-usuario/battle-game.git
   cd battle-game
   ```
2. Instale as dependências:
   ```
   go mod tidy
   ```
3. Execute a aplicação:
   ```
   go run cmd/main/main.go
   ```

A aplicação estará disponível em `http://localhost:8080`.

## Tecnologias Utilizadas 💻
- Go
- MySQL
- JWT
- Gorilla Mux
- Viper

## Contribuição 🤝
1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/fooBar`).
3. Commit suas mudanças (`git commit -am 'Add some fooBar'`).
4. Push para a branch (`git push origin feature/fooBar`).
5. Crie um novo Pull Request.

## Licença 📄
Distribuído sob a licença MIT. Veja `LICENSE` para mais informações.
