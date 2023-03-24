
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root DB_NAME

dropdb:
	docker exec -it postgres12 dropdb DB_NAME

migrateup:
	migrate -path migrations -database "postgres://root:secret@localhost:5432/DB_NAME?sslmode=disable" up

migratedown:
	migrate -path migrations -database "postgres://root:secret@localhost:5432/DB_NAME?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test ./...

sever:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server