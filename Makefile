postgresInit:
	docker run --name postgresforChat	-p  5430:5432	-e POSTGRES_USER=root -e POSTGRES_PASSWORD=test -d postgres:latest

postgres:
	docker exec -it postgresforChat psql

createdb:
	docker exec -it postgresforChat createdb --username=root  --owner=root go-chat

dropdb:
	docker exec -it postgresforChat dropdb  go-chat

createRedisdb:
	docker exec -it redisChat createdb --username=root  --owner=root go-chat

migrateup:
	migrate -path db/migrations -database "postgres://root:test@localhost:5430/go-chat?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migrations -database "postgres://root:test@localhost:5430/go-chat?sslmode=disable" --verbose down

dockerup:
	docker-compose	up

dockerdown:
	docker-compose down

build:
	docker build -t omachat .

omachat:
	docker run  --name  omatest -d  omachat

.PHONY: postgresInit postgres createdb dropdb migrateup migratedown build