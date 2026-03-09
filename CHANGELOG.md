# Changelog for Go Server

Previous changelog before 1.7 will not be noted here.

This changelog is for the server only.

## [3.1.0] - 2026-03-09

### Breaking Changes

- **`/data` response shape changed** — endpoint now returns all sensors as `{"sensors": [...]}` instead of a single sensor's `{temperature, humidity, lastUpdated}`

### Added

- **Multi-sensor `/data` endpoint** — returns an array of all registered sensors, each with `id`, `name`, `temperature`, `humidity`, `lastUpdated`, and `online` fields
- **`GetAllSensors()` service function** — iterates all registered sensors and returns public details

### Changed

- **Index page redesigned** — dark-themed responsive UI showing all rooms as cards with online/offline badges; no external CSS dependencies
- **Server-side template rendering removed** — `index.html` is served as static HTML; all data is fetched client-side via `/data`

## [3.0.0] - 2026-03-05

### Breaking Changes

- **Removed Prometheus** — `/metrics` endpoint and all `prometheus/client_golang` dependencies removed
- **Requires TimescaleDB** — new `DATABASE_URL` environment variable required

### Added

- **TimescaleDB push metrics** — sensor readings are now pushed to a `sensor_readings` hypertable on every update
- **Downtime event tracking** — `sensor_downtime_events` table with `down`/`resolved` pairs for Grafana-friendly visualization
- **SQL migration** — `sql/0001_create_sensor_readings.sql` with hypertable creation
- **OpenAPI spec** — build-time generated via `swaggo/swag`; served at `/openapi.json` and `/openapi.yaml`
- **Scalar API docs** — `/scalar` serves interactive API documentation UI
- **Makefile** — `dev`, `build`, `test`, `lint`, `fmt`, `tidy`, `docs`, `docker-build`, `migrate` targets
- **Makefile `.env` loading** — automatically loads `.env` file when present
- **`is_migrated` column** — `sensor_readings` includes a boolean flag for data migrated from Prometheus

### Changed

- **Go version** bumped from 1.21 to 1.23
- **Docker base images** bumped to `golang:1.23-alpine3.21` / `alpine:3.21`
- **All static assets embedded** — `index.html`, `scalar.html`, `swagger.json`, `swagger.yaml` are compiled into the binary via `go:embed`; Docker image only needs the executable
- **Dependencies** updated (`testify` v1.10.0, added `pgx/v5` v5.7.4)

### Fixed

- **Health check false alarm on startup** — sensors now initialize `lastUpdated` to current time, preventing immediate alert escalation
- **OpenAPI spec missing `required`** — all non-nullable fields in `DataResponse` and `UpdateRequest` are now marked as required

## [2.4.0] - 2023-12-24

- feat: don't put uptime kuma into logs

## [2.3.0] - 2023-12-23

- feat: only allow local ip in some path for system that don't use nginx

## [2.2.1] - 2023-12-21

- feat: add arm64 docker image

## [2.2.0] - 2023-12-18

- server: handle 404

## [2.1.0] - 2023-12-18

- fix: set to NaN when sensor down

## [2.0.0] - 2023-12-16

- refactor: restructure project to be much better
- refactor: safer alert level
- feat: support more than one sensor
- misc improvements

## [1.7.0] - 2023-09-13

- fix: Prometheus now set to NaN
- chore: change threshold time
