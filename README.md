# 📞 Contatos-Go

> **API REST para gerenciamento de contatos desenvolvida em Go**

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-Web%20Framework-green.svg)](https://gin-gonic.com/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 🚀 Sobre o Projeto

Contatos-Go é uma API REST completa para gerenciamento de contatos, desenvolvida em Go utilizando o framework Gin. O projeto implementa operações CRUD (Create, Read, Update, Delete) com validação de dados e tratamento de erros.

### ✨ Funcionalidades

- ✅ **CRUD Completo de Usuários**
- 🔍 **Busca por ID e Email**
- 📝 **Validação de Dados**
- 🛡️ **Tratamento de Erros**
- 🚀 **API REST com Gin Framework**
- 📦 **Estrutura Modular e Organizada**

### 🛠️ Tecnologias Utilizadas

- **Go** - Linguagem principal
- **Gin** - Framework web para APIs
- **Validator** - Validação de dados
- **Godotenv** - Gerenciamento de variáveis de ambiente

---

## 📋 Pré-requisitos

- Go 1.21 ou superior
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

3. **Configure as variáveis de ambiente**

   ```bash
   cp .env.example .env
   # Edite o arquivo .env com suas configurações
   ```

4. **Execute o servidor**

   ```bash
   go run main.go
   ```

O servidor estará rodando em `http://localhost:8080`

---

## 📚 Documentação da API

### Endpoints Disponíveis

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/getUserById/:userId` | Busca usuário por ID |
| `GET` | `/getUserByEmail/:userEmail` | Busca usuário por email |
| `POST` | `/createUser` | Cria novo usuário |
| `PUT` | `/updateUser/:userId` | Atualiza usuário existente |
| `DELETE` | `/deleteUser/:userId` | Remove usuário |

---

## 📁 Estrutura do Projeto

```zsh
contatos-go/
├── main.go                 # Ponto de entrada da aplicação
├── go.mod                  # Dependências do Go
├── README.md              # Documentação do projeto
├── src/
│   ├── config/            # Configurações
│   │   ├── error_handler/ # Tratamento de erros
│   │   └── validation/    # Validação de dados
│   ├── controller/        # Controladores da API
│   │   ├── model/         # Modelos de request/response
│   │   └── routes/        # Definição de rotas
│   ├── model/             # Modelos de dados
│   ├── test/              # Testes
│   └── view/              # Views (se aplicável)
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
- [Godotenv](https://github.com/joho/godotenv) - Gerenciamento de variáveis de ambiente
