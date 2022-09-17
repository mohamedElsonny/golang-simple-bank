PG_USER=root
PG_PASS=postgres
PG_PORT=5555
PG_DB=simple_bank
PG_TEST_DB=simple_bank_test
PG_URL="postgresql://${PG_USER}:${PG_PASS}@localhost:${PG_PORT}/${PG_DB}?sslmode=disable"
PG_TEST_URL="postgresql://${PG_USER}:${PG_PASS}@localhost:${PG_PORT}/${PG_TEST_DB}?sslmode=disable"

run_postgres:
	docker run --name postgres12 -p ${PG_PORT}:5432 -e POSTGRES_USER=${PG_USER} -e POSTGRES_PASSWORD=${PG_PASS} -d postgres:12-alpine

run_pgadmin:
	docker run -p 5050:80 -e PGADMIN_DEFAULT_EMAIL='pgadmin4@pgadmin.org' -e PGADMIN_DEFAULT_PASSWORD='admin' -d --name pgadmin4 dpage/pgadmin4

stop_postgres:
	docker stop postgres12

stop_pgadmin:
	docker stop pgadmin4

createdb:
	docker exec -it postgres12 createdb --username=${PG_USER} --owner=${PG_USER} ${PG_DB}

createdb_test:
	docker exec -it postgres12 createdb --username=${PG_USER} --owner=${PG_USER} ${PG_TEST_DB}

migrate_up:
	migrate -path db/migration -database=${PG_URL} -verbose up

migrate_down:
	migrate -path db/migration -database=${PG_URL} -verbose down


migrate_up_test:
	migrate -path db/migration -database=${PG_TEST_URL} -verbose up

migrate_down_test:
	migrate -path db/migration -database=${PG_TEST_URL} -verbose down

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

dropdb:
	docker exec -it postgres12 dropdb simple_bank

dev:
	go run main.go 
