postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d -p 5432:5432 postgres:12-alpine
createdb:
	docker exec -it postgres12  createdb --username=root --owner=root simple_social_media
dropdb:
	docker exec -it postgres12  dropdb simple_social_media
cleardb:
	docker exec -it postgres12  psql -U root -d simple_social_media -c "TRUNCATE users, posts, comments RESTART IDENTITY;"
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_social_media?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_social_media?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server cleardb
