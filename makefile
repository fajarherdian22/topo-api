create_migrate:
	migrate create -ext .sql -dir db/migration -seq name_migrate
create_db:
	docker exec -it mysql mysql -uadmin -padmin1234 -e "CREATE DATABASE dashboard;"
migrateup:
	migrate -path db/migration -database "mysql://admin:admin1234@tcp(localhost:3306)/dashboard" --verbose up
migratedown:
	migrate -path db/migration -database "mysql://admin:admin1234@tcp(localhost:3306)/dashboard" --verbose down
migrateup1:
	migrate -path db/migration -database "mysql://admin:admin1234@tcp(localhost:3306)/dashboard" --verbose up 1
migratedown1:
	migrate -path db/migration -database "mysql://admin:admin1234@tcp(localhost:3306)/dashboard" --verbose down 1
server:
	go run main.go
sqlc:
	sqlc generate
