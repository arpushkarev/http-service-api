.PHONY: run
run:
	go run -v ./cmd

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=zt-service user=zt-service-user password=zt-service-password sslmode=disable"

.PHONY: install-goose
.install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: .local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: Local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v
