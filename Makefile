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

setFlyEnvStg:
	ifeq ($(key),)
	$(error key is not set)
	endif
	ifeq ($(value),)
	$(error value is not set)
	endif
	flyctl -c $(stg-config) secrets set $(key)=$(value)

unsetFlyEnvStg:
	ifeq ($(key),)
	$(error key is not set)
	endif
	flyctl -c $(stg-config) secrets unset $(key)

setFlyEnvPrd:
	ifeq ($(key),)
	$(error key is not set)
	endif
	ifeq ($(value),)
	$(error value is not set)
	endif
	flyctl -c $(prd-config) secrets set $(key)=$(value)

unsetFlyEnvPrd:
	ifeq ($(key),)
	$(error key is not set)
	endif
	flyctl -c $(prd-config) secrets unset $(key)

# memo
# cat .env.prd | flyctl secrets import -c ./.github/workflows/fly.production.toml
# cat .env.stg | flyctl secrets import -c ./.github/workflows/fly.staging.toml
