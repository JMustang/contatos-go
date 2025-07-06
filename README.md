# 📞 Contatos-Go

> **API REST para gerenciamento de contatos desenvolvida em Go**

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-Web%20Framework-green.svg)](https://gin-gonic.com/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 🚀 Sobre o Projeto

Contatos-Go é uma API REST completa para gerenciamento de contatos, desenvolvida em Go utilizando o framework Gin. O projeto implementa operações CRUD (Create, Read, Update, Delete) com validação de dados, tratamento de erros robusto e sistema de logging estruturado.

### ✨ Funcionalidades Implementadas

- ✅ **Estrutura de Rotas CRUD**
  - `GET /getUserById/:userId` - Busca usuário por ID
  - `GET /getUserByEmail/:userEmail` - Busca usuário por email
  - `POST /createUser` - Cria novo usuário
  - `PUT /updateUser/:userId` - Atualiza usuário existente
  - `DELETE /deleteUser/:userId` - Remove usuário

- ✅ **Sistema de Validação Avançado**
  - Validação de campos obrigatórios
  - Validação de formato de email
  - Validação de senha (mínimo 6 caracteres, deve conter caracteres especiais)
  - Validação de nome (4-100 caracteres)
  - Validação de idade (1-140 anos)
  - Tradução automática de mensagens de erro

- ✅ **Tratamento de Erros Robusto**
  - Erros de validação estruturados
  - Erros de tipo de dados
  - Erros de requisição malformada
  - Códigos de status HTTP apropriados
  - Mensagens de erro detalhadas

- ✅ **Sistema de Logging Estruturado**
  - Logs em formato JSON
  - Configurável via variáveis de ambiente
  - Diferentes níveis de log (INFO, ERROR, DEBUG)
  - Output configurável (stdout, arquivo)

- ✅ **Estrutura Modular e Organizada**
  - Separação clara de responsabilidades
  - Controllers para lógica de negócio
  - Models para estruturas de dados
  - Configurações centralizadas

### 🛠️ Tecnologias Utilizadas

- **Go 1.24.4** - Linguagem principal
- **Gin** - Framework web para APIs REST
- **Validator v10** - Validação de dados com tradução
- **Zap** - Sistema de logging estruturado
- **Godotenv** - Gerenciamento de variáveis de ambiente

---

## 📋 Pré-requisitos

- Go 1.24.4 ou superior
- Git

## 🚀 Como Executar

1. **Clone o repositório**

   ```bash
   git clone https://github.com/JMustang/contatos-go.git
   cd contatos-go
   ```

2. **Instale as dependências**

   ```bash
   go mod tidy
   ```

3. **Configure as variáveis de ambiente (opcional)**

   ```bash
   cp .env.example .env
   # Edite o arquivo .env com suas configurações
   ```

   Variáveis disponíveis:
   - `LOG_OUTPUT` - Saída dos logs (stdout, arquivo)
   - `LOG_LEVEL` - Nível de log (info, error, debug)

4. **Execute o servidor**

   ```bash
   go run main.go
   ```

O servidor estará rodando em `http://localhost:8080`

---

## 📚 Documentação da API

### Modelo de Usuário

```json
{
  "email": "test@test.com",
  "password": "password#@#@!2121",
  "name": "John Doe",
  "age": 30
}
```

### Validações Implementadas

- **Email**: Formato válido de email
- **Password**: Mínimo 6 caracteres, deve conter caracteres especiais (!@#$%*)
- **Name**: Entre 4 e 100 caracteres
- **Age**: Entre 1 e 140 anos

### Endpoints Disponíveis

| Método | Endpoint | Descrição | Status |
|--------|----------|-----------|--------|
| `GET` | `/getUserById/:userId` | Busca usuário por ID | ✅ Implementado |
| `GET` | `/getUserByEmail/:userEmail` | Busca usuário por email | ✅ Implementado |
| `POST` | `/createUser` | Cria novo usuário | ✅ Implementado |
| `PUT` | `/updateUser/:userId` | Atualiza usuário existente | ✅ Implementado |
| `DELETE` | `/deleteUser/:userId` | Remove usuário | ✅ Implementado |

### Exemplo de Resposta de Erro

```json
{
  "message": "Some fields are invalid",
  "error": "bad_request",
  "code": 400,
  "causes": [
    {
      "field": "email",
      "message": "email must be a valid email address"
    },
    {
      "field": "password",
      "message": "password must be at least 6 characters long"
    }
  ]
}
```

---

## 📁 Estrutura do Projeto

```zsh
contatos-go/
├── main.go                 # Ponto de entrada da aplicação
├── go.mod                  # Dependências do Go
├── go.sum                  # Checksums das dependências
├── README.md              # Documentação do projeto
├── LICENSE                # Licença MIT
├── src/
│   ├── config/            # Configurações
│   │   ├── error_handler/ # Sistema de tratamento de erros
│   │   │   └── error_handler.go
│   │   ├── validation/    # Sistema de validação
│   │   │   └── validate_user.go
│   │   └── logger/        # Sistema de logging
│   │       └── logger.go
│   ├── controller/        # Controladores da API
│   │   ├── createUser.go  # Criação de usuário
│   │   ├── findUser.go    # Busca de usuários
│   │   ├── updateUser.go  # Atualização de usuário
│   │   ├── deleteUser.go  # Remoção de usuário
│   │   ├── model/         # Modelos de request/response
│   │   │   ├── request/
│   │   │   │   └── user_request.go
│   │   │   └── response/
│   │   └── routes/        # Definição de rotas
│   │       └── routes.go
│   ├── model/             # Modelos de dados
│   ├── test/              # Testes
│   └── view/              # Views (se aplicável)
```

---

## 🔧 Configuração

### Variáveis de Ambiente

O projeto suporta as seguintes variáveis de ambiente:

- `LOG_OUTPUT`: Define onde os logs serão exibidos (padrão: "stdout")
- `LOG_LEVEL`: Define o nível de log (padrão: "info")

### Logging

O sistema de logging utiliza o Zap para logs estruturados em JSON:

```json
{
  "level": "info",
  "time": "2024-01-15T10:30:00Z",
  "message": "Starting server",
  "caller": "main.go:15"
}
```

---

## 🤝 Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## 👨‍💻 Autor

**JMustang**

- GitHub: [@JMustang](https://github.com/JMustang)

---

## 🙏 Agradecimentos

- [Gin Framework](https://gin-gonic.com/) - Framework web para Go
- [Go Playground Validator](https://github.com/go-playground/validator) - Validação de dados
- [Zap](https://github.com/uber-go/zap) - Sistema de logging estruturado
- [Godotenv](https://github.com/joho/godotenv) - Gerenciamento de variáveis de ambiente
