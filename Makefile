DIR=${CURDIR}

db-up:
	docker-compose up --build user-postgres

migrations-up:
	docker run --rm -v ${DIR}/pkg/repository/postgres/migrations:/migrations \
	--network login_default migrate/migrate -path=/migrations \
	-database "postgres://dev:12345@user-postgres:5432/postgres?sslmode=disable" up

db-down:
	docker-compose down
