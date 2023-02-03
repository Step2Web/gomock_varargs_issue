BUILD_OPTS=
DEV_BUILD_OPTS=$(BUILD_OPTS) -buildvcs=false
PROD_BUILD_OPTS=$(BUILD_OPTS) -tags=timetzdata -trimpath

TEST_OPTS=--race --tags=all -coverprofile=coverage.txt -covermode=atomic

.PHONY: build
build:
	go install $(DEV_BUILD_OPTS) ./...

.PHONY: build-prod
build-prod:
	go install $(PROD_BUILD_OPTS) ./...

.PHONY: mocks
mocks:
	mockgen -package=main github.com/willfaught/gockle Session,Query > gocql_mock.go

.PHONY: test
test:
	go test $(TEST_OPTS) ./...

.PHONY: install-tools
install-tools:
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: fmt
fmt:
	# Fixup modules
	go mod tidy
	# Format the Go sources:
	go fmt ./...

.PHONY: lint
lint:
	# Lint the Go source:
	go vet ./...
	staticcheck ./...

.PHONY: update-dependencies
update-dependencies: 
	go get -u
	go mod tidy

.PHONY: setup
setup: install-tools build mocks
