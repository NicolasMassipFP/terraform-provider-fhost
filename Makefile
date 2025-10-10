PROVIDER_NAME=terraform-provider-fhost
PROVIDER_VERSION=0.0.1
PLUGIN_DIR=plugins/registry.terraform.io/training-sop-fsmc/fhost/${PROVIDER_VERSION}/linux_amd64
RUN=./scripts/run_go

all: build

# Show available targets and file management commands
.PHONY: help
help:
	@echo "SMC Terraform Provider Build Commands:"
	@echo "  make all                - Full build (docker + codegen + go build)"
	@echo "  make build              - Same as 'all'"
	@echo "  make go-build           - Build the Go binary"
	@echo "  make clean              - Clean all build artifacts"
	@echo "  make help               - Show this help message"


.PHONY: build
build: docker-build go-build

.PHONY: go-build
go-build:
	mkdir $(PLUGIN_DIR) 2>/dev/null; true
	$(RUN) go build  -o $(PLUGIN_DIR)/terraform-provider-fhost .
	# $(RUN) go build  -o $(PLUGIN_DIR)/terraform-provider-fhost ./...

.PHONY: install
install:
	@go install .

.PHONY: check
check: docker-build build
	@mkdir -p .cache .cache/go-build
	@scripts/run_go golangci-lint run

.PHONY: test
test: docker-build
	@scripts/run_go go test ./...

.PHONY: test-verbose
test-verbose: docker-build
	@scripts/run_go go test -v ./...

.PHONY: test-coverage
test-coverage: docker-build
	@scripts/run_go go test -cover ./...

.PHONY: test-coverage-html
test-coverage-html: docker-build
	@scripts/run_go go test -coverprofile=coverage.out ./...
	@scripts/run_go go tool cover -html=coverage.out -o coverage.html

.PHONY: docker-build
docker-build:
	docker build -q -t go-docker docker/

.PHONY: fmt
fmt: build
	@scripts/run_go gofmt -w .

.PHONY: update-dependencies
update-dependencies:
	go mod tidy

.PHONY: clean
clean:
	rm -f coverage.out coverage.html
	chmod -R u+w .cache/gomod || true
	rm -fr .cache $(PROVIDER_NAME)
