migrate:
	migrate -database ${DB_POSTGRES_URL} -path ./db/migrations up