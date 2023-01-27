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

migrateupmysql:
	migrate -path db/schemas -database "jdbc:mysql://docker-user-golang-database-mysql-1:3307/user_golang?sslmode=disable" --verbose up

migrateupmysql:
	migrate -path db/schemas -database "jdbc:mysql://docker-user-golang-database-mysql-1:3307/user_golang?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v ${CURDIR}:/src -w /src kjconroy/sqlc generate

test: 
	go test -v -cover ./...
