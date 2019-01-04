OUTFILE = ./bin/jsonstream
CLI_PATH = ./jsonstream-cli

PROJECT_PKGS := $$(go list ./... | egrep -v -e '(/vendor/|/mocks)')

DEV_PACKAGES := \
	github.com/golangci/golangci-lint/cmd/golangci-lint

install-dev-deps:
	$(foreach pkg,$(DEV_PACKAGES),go get -u $(pkg);)

vendor:
	go mod vendor

lint: 
	golangci-lint run --exclude-use-default=false ./...

test:
	go test -race ./...

cover:
	go test -cover ./...
	
bench:
	go test -bench=. -benchmem ./...

build:
	go build -o $(OUTFILE) $(CLI_PATH)

