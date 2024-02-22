migrate:
	migrate -database ${DB_POSTGRES_URL} -path ./migrations up

run-test:
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"