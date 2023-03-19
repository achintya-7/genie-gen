# Genie-Gen
## Description
A simple cli tool to generate a Go Backend project using gin-gonic, sqlc, viper and postgresql.
It will generate a project with the following structure:
```
├── pkgName
│   │── api
│   │   ├── server.go
│   │── db
│   │   ├── migration
│   │   ├── query
│   │   ├── sqlc
│   │   │   ├── store.go 
│   ├── util
│   │   ├── config.go
│   │── app.env
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── Makefile
│   ├── sqlc.yaml
```
1. The `Makefile` will have various useful commands to run the project, tests, sqlc, docker image, etc.
2. Use `golang-migrate` to manage the database migrations.
3. Write all your sql queries in the `query` folder.
4. The `sqlc.yaml` will have the sqlc configuration to generate the go code from the sql queries.
5. The `sqlc` dir will have the generated go code from the sql queries.
6. The `config.go` will have the viper configuration to read the environment variables.
7. Use `app.env` to store the environment variables.
8. The `server.go` will have the gin-gonic server configuration.


## Installation
```bash
go install github.com/achintya-7/genie-gen@latest
```

## Usage
```bash
genie-gen init
```
Follow the prompts to generate the project.

## Requirements
1. Go 1.18+
2. Docker
3. golang-migrate (optional)
4. sqlc (optional)

## Demo
