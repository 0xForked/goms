# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

# --- Tooling & Variables ----------------------------------------------------------------
include ./misc/make/tools.Makefile

install-deps: gotestsum mockery golangci-lint
deps: $(GOTESTSUM) $(MOCKERY) $(GOLANGCI)
deps:
	@echo "Required Tools Are Available"

mock:
	mockery --all --output=mocks/genmocks --outpkg=mocks

api-spec: tests
	@ echo "Re-generate Swagger File (API Spec docs)"
	@ swag init --parseDependency --parseInternal \
		--parseDepth 4 -g ./cmd/bff/main.go
	@ echo "generate swagger file done"

tests: $(GOTESTSUM) lint
	@ echo "Run tests"
	@ gotestsum --format pkgname-and-test-fails \
		--hide-summary=skipped \
		-- -coverprofile=cover.out ./...
	@ rm cover.out

lint: $(GOLANGCI)
	@ echo "Applying linter"
	@ golangci-lint cache clean
	@ golangci-lint run -c .golangci.yaml ./...

proto:
	protoc --go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:.  \
        ./pkg/pb/store.proto \
        ./pkg/pb/book.proto
