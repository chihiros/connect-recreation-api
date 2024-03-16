up:
	DOCKER_BUILDKIT=1 docker compose up --build $(filter-out $@,$(MAKECMDGOALS))

%:
	@:

down:
	docker compose down

db-in:
	docker compose exec db bash -c "psql \"user=postgres password=postgres_pw host=localhost port=5432 dbname=postgres\""

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
