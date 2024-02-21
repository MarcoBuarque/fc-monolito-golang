migrate:
	migrate -database ${DB_POSTGRES_URL} -path ./migrations up