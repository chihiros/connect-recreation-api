%:
	@:

up: docker-up-db goose-migration-up goose-seed-up docker-up-api

docker-up:
	docker compose up -d

docker-up-db:
	docker compose up -d db
	@echo "Waiting for DB to be ready..."
	@until docker compose exec db pg_isready -U postgres > /dev/null 2>&1; do \
		sleep 1; \
		echo -n "."; \
	done
	@echo "DB is ready."

docker-up-api:
	docker compose up api

build-up:
	DOCKER_BUILDKIT=1 docker compose up --build $(filter-out $@,$(MAKECMDGOALS))

down:
	docker compose down

db-in:
	docker compose exec db bash -c "psql \"user=postgres password=postgres host=localhost port=5432 dbname=postgres\""

prune:
	docker system prune

gen:
	cd app; \
	go generate ./ent

gomod-update:
	cd app; \
	go get -u && go mod tidy

stg-config := ./.github/workflows/fly.staging.toml
prd-config := ./.github/workflows/fly.production.toml

deploy-stg:
	flyctl deploy --config $(stg-config) --build-target deploy --remote-only

deploy-prod:
	flyctl deploy --config $(prd-config) --build-target deploy --remote-only

golint:
	cd app; \
	golangci-lint run ./...

# OpenAPI Generatorでコード生成が出来るかを確認するためのコマンド
ogen:
	docker run --rm \
		-v $$PWD:/local openapitools/openapi-generator-cli generate \
		-i /local/docs/openapi.yml \
		-g typescript-axios \
		--additional-properties supportsES6=true,withInterfaces=true \
		-o /local/docs/test-generate
	rm -rf docs/test-generate

goose-create-migration:
	goose -dir ./schema/migrations create $(filter-out $@,$(MAKECMDGOALS)) sql

goose-create-seed:
	goose -dir ./schema/seeds create $(filter-out $@,$(MAKECMDGOALS)) sql

goose-migration-up:
	goose -dir ./schema/migrations postgres "user=postgres password=postgres host=localhost port=5432 dbname=postgres" up

goose-migration-down:
	goose -dir ./schema/migrations postgres "user=postgres password=postgres host=localhost port=5432 dbname=postgres" down

goose-seed-up:
	goose -dir ./schema/seeds postgres "user=postgres password=postgres host=localhost port=5432 dbname=postgres" up

goose-seed-down:
	goose -dir ./schema/seeds postgres "user=postgres password=postgres host=localhost port=5432 dbname=postgres" down
