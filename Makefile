PROVIDER_NAME=fhost
PROVIDER_FULLNAME=terraform-provider-fhost
PROVIDER_VERSION=0.1.4
PROVIDER_ORGANIZATION=NicolasMassipFP
PLUGIN_DIR=plugins/registry.terraform.io/${PROVIDER_ORGANIZATION}/${PROVIDER_NAME}/${PROVIDER_VERSION}/linux_amd64
RUN=./scripts/run_go

all: build

# Show available targets and file management commands
.PHONY: help
help:
	@echo "SMC Terraform Provider Build Commands:"
	@echo "  make all                - Full build (docker + codegen + go build)"
	@echo "  make build              - Same as 'all'"
	@echo "  make release            - Prepare the provider release (build + goreleaser)"
	@echo "  make go-build           - Build the Go binary"
	@echo "  make go-release         - Build multi-platform binaries with goreleaser (no dirty edit allowed, tag must be set on HEAD)"
	@echo "  make go-release-snapshot- Build snapshot without publishing (no validation, dirty edit allowed, tag must be set on HEAD)"
	@echo "  make docs               - Generate provider documentation"
	@echo "  make clean              - Clean all build artifacts"
	@echo "  make help               - Show this help message"


.PHONY: build
build: docker-build go-build

.PHONY: release
release: docker-build go-release

.PHONY: go-build
go-build:
	mkdir $(PLUGIN_DIR) 2>/dev/null; true
	$(RUN) go build  -o $(PLUGIN_DIR)/${PROVIDER_FULLNAME} .

.PHONY: docs
docs:
	mkdir docs 2>/dev/null; true
	@echo "\nWARNING: us local go command, not docker go image\n"
	@echo "Retrieving dependencies..."
	cd tools && go mod tidy
	@echo "Formating examples...(not working currenlty --> generate errors )"
	@echo " NOT ENABLED: cd tools && terraform fmt -recursive ../examples/"
	@echo "Generating provider documentation... ()"
	# using docker image in sub folder do not work ( run_go rely o, $PWD to load cache )
	cd tools && go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-dir .. -provider-name smc

.PHONY: go-release-snapshot
go-release-snapshot: docker-build
	$(RUN) goreleaser build --snapshot --clean --skip=validate

.PHONY: go-release
go-release:
	$(RUN) ./scripts/init_pgp_and_make_release.sh

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
	rm -fr .cache $(PROVIDER_FULLNAME) dist/
