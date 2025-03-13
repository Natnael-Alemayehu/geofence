# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Geofense
run-geofence:
	go run api/service/geofence/main.go | go run api/tooling/logfmt/main.go

run-status:
	curl http -i -X GET localhost:3000/v1/status

run-verify-inside:
	curl -i -X POST \
	-H 'Content-Type: application/json' \
	-d '{"latitude":9.02921925586169,"longitude":38.741409590890214}' \
	localhost:3000/v1/verify_location

run-verify-outside:
	curl -i -X POST \
	-H 'Content-Type: application/json' \
	-d '{"latitude":9.02921925586169,"longitude":40.741409590890214}' \
	localhost:3000/v1/verify_location

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all


# ==============================================================================
build: 
	sudo docker build -t geofence .

compose-up:
	docker compose -f zarf/compose/docker_compose.yaml up -d	

compose-down:
	docker compose -f zarf/compose/docker_compose.yaml down

compose-restart: build compose-down compose-up


# ==============================================================================
# Administration

migrate:
	export GEOFENCE_DB_HOST=localhost; go run api/tooling/admin/main.go migrate

seed: migrate
	export GEOFENCE_DB_HOST=localhost; go run api/tooling/admin/main.go seed

pgcli:
	pgcli postgresql://postgres:postgres@localhost