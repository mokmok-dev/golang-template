.PHONY: dev
dev:
	@go run github.com/cosmtrek/air -c .air.toml

.PHONY: format
format:
	$(call format)

.PHONY: generate.mock
generate.mock:
	@go generate ./domain/...
	$(call format)

.PHONY: generate.sqlc
generate.sqlc:
	@go run github.com/kyleconroy/sqlc/cmd/sqlc generate -f sqlc.yaml
	$(call format)

.PHONY: generate.wire
generate.wire:
	@go run github.com/google/wire/cmd/wire ./...
	$(call format)

.PHONY: migrate.plan
migrate.plan:
	@go run github.com/k0kubun/sqldef/cmd/psqldef -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USER} -W ${DATABASE_PASSWORD} ${DATABASE_NAME} --dry-run < schema.sql

.PHONY: migrate.apply
migrate.apply:
	@go run github.com/k0kubun/sqldef/cmd/psqldef -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USER} -W ${DATABASE_PASSWORD} ${DATABASE_NAME} < schema.sql

define format
	@go run mvdan.cc/gofumpt -l -w .
	@go run golang.org/x/tools/cmd/goimports -w .
	@go mod tidy
endef
