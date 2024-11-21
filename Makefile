export MYSQL_URL = "mysql://root:superSecretPassword@tcp(localhost:3306)/situs_forum"

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down

migrate-force:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations force 7

migrate-create-schema-migrations:
	@ migrate create -ext sql -dir scripts/migrations/schema_migrations -seq $(name)

migrate-down-schema-migrations:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations/schema_migrations down

run:
	@ go run cmd/main.go