
# Resumo do Projeto: API REST em Go com Gin, DDD e Docker

Descrição Geral:

Este projeto consiste em desenvolver uma API REST em Go utilizando a biblioteca Gin para gerenciar as requisições HTTP. A estrutura do código segue o padrão de arquitetura DDD (Domain-Driven Design), separando o projeto em camadas: Domínio (regras de negócio e entidades), Aplicação (casos de uso), Interface (controladores e adaptadores) e Infraestrutura (conexão com banco de dados e configurações).

O banco de dados utilizado é o PostgreSQL, que será executado em um contêiner Docker, facilitando a configuração e o gerenciamento do ambiente. Essa abordagem permite criar uma aplicação modular e escalável, com uma clara separação de responsabilidades.


Estruta de pastas: 

```
└── 📁go_rest_api
    └── 📁cmd
        └── main.go
    └── 📁controller
        └── drink_controller.go
    └── 📁db
        └── conn.go
    └── 📁model
        └── drinks.go
    └── 📁repository
        └── drinks_repository.go
    └── 📁usecase
        └── drinks_usecase.go
    └── .gitignore
    └── docker-compose.yml
    └── go.mod
    └── go.sum
```



## Instalação

Instale as dependencias do projeto utilizando o 

```bash
  go mod tidy
```

então rode o para subir o docker 

```bash
  docker-compose up
```

e o comando para entrar na pasta 

```bash
  cd cmd
```

o comando para subir a aplicação principal
  
```bash
  go run main.go
``` 