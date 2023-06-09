# Genie-Gen
## Description
A simple cli tool to generate a Go Backend project using gin-gonic, sqlc, viper and postgresql.
It will generate a basic structure only, some modifications will be required to make it production ready and error free.
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
![genie_demo](https://user-images.githubusercontent.com/67036708/226205523-f653b3f1-82ad-44cd-8661-467d937a6778.gif)
