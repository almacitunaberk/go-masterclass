postgres_run:
	docker run --name postgres-go-masterclass -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2

createdb:
	docker exec -it postgres-go-masterclass createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-go-masterclass dropdb simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

server:
	go run main.go

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres_run migrate_up migrate_down sqlc test server