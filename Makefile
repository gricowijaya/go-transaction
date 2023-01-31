postgres:
	docker run --name postgres12 -p 5000:5432 -e POSTGRES_USER=root POSTGRES_PASSWORD=secret -d postgres:14.1-alpine

createdb:
	docker exec -it docker-user-golang-database-1 createdb --username=postgres --owner=postgres user_golang

dropdb:
	docker exec -it docker-user-golang-database-1 dropdb --username=postgres user_golang

migrateup:
	migrate -path db/schemas -database "postgresql://postgres:password@localhost:5000/user_golang?sslmode=disable" --verbose up

migratedown:
	migrate -path db/schemas -database "postgresql://postgres:password@localhost:5000/user_golang?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v ${CURDIR}:/src -w /src kjconroy/sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
