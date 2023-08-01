postgres_run:
	docker run --name postgres-go-masterclass -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2

createdb:
	docker exec -it postgres-go-masterclass createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-go-masterclass dropdb simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migrate_up1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up 1

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

migrate_down1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down 1

server:
	go run main.go

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/almacitunaberk/go_masterclass/db/sqlc Store

.PHONY: createdb dropdb postgres_run migrate_up migrate_down migrate_up1 migrate_down1 sqlc test server mockdb