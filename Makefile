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
	-d '{"location_id":"delivery_zone_1","latitude":9.02921925586169,"longitude":38.741409590890214}' \
	localhost:3000/v1/verify_location

run-verify-outside:
	curl -i -X POST \
	-H 'Content-Type: application/json' \
	-d '{"location_id":"delivery_zone_1","latitude":9.02921925586169,"longitude":40.741409590890214}' \
	localhost:3000/v1/verify_location

run-search-location:
	curl -i -X GET localhost:3000/v1/location/delivery_zone_1

run-create-location:
	curl -i -X POST \
	-H 'Content-Type: application/json' \
	-d '{"id":"new_zone","geojson":{"type":"Polygon", "coordinates":[[[38.752323744414866,9.03534632727542],[38.752323744414866,9.034837522071527],[38.75271335641881,9.034837522071527],[38.75271335641881,9.03534632727542],[38.752323744414866,9.03534632727542]]]}}' \
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
# Start Development Images

dev-db-up:
	@docker rm -f database >/dev/null 2>&1 || true
	@docker network create mynet >/dev/null 2>&1 || true
	@mkdir -p docker-entrypoint-initdb.d
	@printf "%s\n" \
		"-- Version: 1.01" \
		"-- Description: Create table geolocation with PostGIS support" \
		"" \
		"CREATE EXTENSION IF NOT EXISTS postgis;" \
		"" \
		"CREATE TABLE geolocation (" \
		"    location_id TEXT NOT NULL," \
		"    geojson     GEOMETRY NOT NULL," \
		"    PRIMARY KEY (location_id)" \
		");" > docker-entrypoint-initdb.d/01_migrate.sql
	
	@printf "%s\n" \
		"INSERT INTO geolocation (location_id, geojson) VALUES (" \
		"    'delivery_zone_1'," \
		"    ST_GeomFromGeoJSON('{ \"type\": \"Polygon\", \"coordinates\": [[[38.74256704424312,9.033138471223111],[38.736467448797924,9.032775582701646],[38.73709210616215,9.030017618002177],[38.738194442688865,9.025989500072342],[38.74322844615821,9.02439275618923],[38.74894222381815,9.028021709336926],[38.747435697232305,9.032194960307649],[38.74256704424312,9.033138471223111]]]}')" \
		"),( " \
		"	'test_del'," \
		"    ST_GeomFromGeoJSON('{ \"type\": \"Polygon\", \"coordinates\": [[[38.74256704424312,9.033138471223111],[38.736467448797924,9.032775582701646],[38.73709210616215,9.030017618002177],[38.738194442688865,9.025989500072342],[38.74322844615821,9.02439275618923],[38.74894222381815,9.028021709336926],[38.747435697232305,9.032194960307649],[38.74256704424312,9.033138471223111]]]}')" \
		");" > docker-entrypoint-initdb.d/02_seed.sql
	
	docker run -d --name database --network mynet \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=postgres \
		-p 5432:5432 \
		-v "$(shell pwd)/docker-entrypoint-initdb.d:/docker-entrypoint-initdb.d" \
		postgis/postgis:17-3.4
	@sleep 15
	@docker logs database
	@rm -rf docker-entrypoint-initdb.d

dev-tile-up:
	docker run -p 9851:9851 -d --name tile tile38/tile38 

dev-up: dev-db-up start-tile


# ==============================================================================
# Stop Development Docker Images
dev-db-down:
	@docker rm -f database >/dev/null 2>&1 || true

stop-tile-down:
	@docker rm -f tile >/dev/null 2>&1 || true

dev-down: dev-db-down stop-tile-down