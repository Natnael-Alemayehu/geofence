# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Geofense
start-server:
	go run api/service/geofence/main.go | go run api/tooling/logfmt/main.go

run-status:
	curl http -i -X GET localhost:3000/v1/status

run-verify-inside:
	curl -i -X POST \
	-H 'Content-Type: application/json' \
	-d '{"location_id":"f4d1ce59-b2ab-4879-8038-5903236a15c7","latitude":8.352624746,"longitude":38.0320905}' \
	localhost:3000/v1/verify_location

run-verify-outside:
	curl -i -X POST \
	-H 'Content-Type: application/json' \
	-d '{"location_id":"f4d1ce59-b2ab-4879-8038-5903236a15c7","latitude":8.352624746,"longitude":40.0320905}' \
	localhost:3000/v1/verify_location

run-search-location-by-id:
	curl -i -X GET localhost:3000/v1/location/id/bb75c5ae-d162-43b0-b4b3-ae31f191fb44

run-search-location-by-name:
	curl -i -X GET localhost:3000/v1/location/name/Enge%20Health%20center

run-create-location:
	curl -i -X POST \
	-H 'Content-Type: application/json' \
	-d '{"name":"New Zone","geojson":{"type":"Polygon", "coordinates":[[[38.752323744414866,9.03534632727542],[38.752323744414866,9.034837522071527],[38.75271335641881,9.034837522071527],[38.75271335641881,9.03534632727542],[38.752323744414866,9.03534632727542]]]}}' \
	localhost:3000/v1/location

run-delete-location:
	curl -i -X GET localhost:3000/v1/location/delete/test_del

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

restart:
	docker compose -f zarf/compose/docker_compose.yaml restart geofence


# ==============================================================================
# Administration

migrate:
	export GEOFENCE_DB_HOST=localhost; go run api/tooling/admin/main.go migrate

seed: migrate
	export GEOFENCE_DB_HOST=localhost; go run api/tooling/admin/main.go seed

pgcli:
	pgcli postgresql://postgres:postgres@localhost

# ==============================================================================
# Start Frontend
start-frontend:
	cd api/frontend && npm start

stop-frontend:
	@pkill -f "react-scripts" || echo "No frontend process running"


# ==============================================================================
# Start Development Images

dev-db-up:
	@docker rm -f database >/dev/null 2>&1 || true
	@docker network create mynet >/dev/null 2>&1 || true
	@mkdir -p docker-entrypoint-initdb.d
	@cp business/sdk/migrate/sql/seed.sql docker-entrypoint-initdb.d/02_seed.sql
	@cp business/sdk/migrate/sql/migrate.sql docker-entrypoint-initdb.d/01_migrate.sql
	
	docker run -d --name database --network mynet \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=postgres \
		-p 5432:5432 \
		-v "$(shell pwd)/docker-entrypoint-initdb.d:/docker-entrypoint-initdb.d:ro" \
		postgis/postgis:17-3.4
	@sleep 5
	@docker logs database
	@rm -rf docker-entrypoint-initdb.d

dev-tile-up:
	docker run -p 9851:9851 -d --name tile tile38/tile38 

dev-up: dev-db-up dev-tile-up


# ==============================================================================
# Stop Development Docker Images
dev-db-down:
	@docker rm -f database >/dev/null 2>&1 || true

stop-tile-down:
	@docker rm -f tile >/dev/null 2>&1 || true

dev-down: dev-db-down stop-tile-down