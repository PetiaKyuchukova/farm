migrate:
	goose -dir backend/postgres/migrations -table _db_version mysql '$(DSN)' up

#seed:
#	@goose -dir db/seeds -table _db_seeds mysql '$(DSN)' up

sql:
	sqlc generate
