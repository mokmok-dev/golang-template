.PHONY: dev
dev:
	@go run github.com/cosmtrek/air -c .air.toml

.PHONY: format
format:
	$(call format)

.PHONY: generate.wire
generate.wire:
	@go run github.com/google/wire/cmd/wire ./...
	$(call format)

define format
	@go run mvdan.cc/gofumpt -l -w .
	@go run golang.org/x/tools/cmd/goimports -w .
	@go mod tidy
endef
