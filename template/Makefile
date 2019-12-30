APP_NAME={{Name}}
DEFAULT_PORT={{Port}}
.PHONY: setup init build dev test db-migrate-up db-migrate-down

setup:
	cd ~ && go get -v github.com/rubenv/sql-migrate/...
	cd ~ && go get github.com/golang/mock/gomock
	cd ~ && go get github.com/golang/mock/mockgen
	cp .env.sample .env && vim .env

init: remove-infras
	docker-compose up -d
	@echo "Waiting for database connection..."
	@while ! docker exec {{Name}}_db pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	make db-migrate-up

remove-infras:
	docker-compose stop; docker-compose rm -f

build:
	env GOOS=darwin GOARCH=amd64 go build -o bin/server $(shell pwd)/cmd/server/

dev:
	go run ./cmd/server/main.go

admin:
	go run ./cmd/admin/*.go

test:
	@PROJECT_PATH=$(shell pwd) go test -cover ./...

db-migrate-up:
	sql-migrate up -env=local

db-migrate-down:
	sql-migrate down -env=local

docker-build:
	docker build \
	--build-arg DEFAULT_PORT="${DEFAULT_PORT}" \
	-t ${APP_NAME}:latest .

seed-db-local:
	@docker cp  migrations/postgres/seed/seed.sql  {{Name}}_db:/seed.sql
	@docker exec -t {{Name}}_db sh -c "PGPASSWORD=postgres psql -U postgres -d {{Name}}_db -f /seed.sql"

gen-mock:
	@mockgen -source=./pkg/service/repo/repo.go -destination=./pkg/service/repo/mocks/repo.go
