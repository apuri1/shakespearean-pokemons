clean:
	@rm -r bin .fmtcheck .test .tools || true

fmt:
	bin/golangci-lint run --fix

.fmtcheck: .tools $(go_files)
	bin/golangci-lint run
	@touch .fmtcheck

.tools:
	curl -sSfLk https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.31.0
	@touch .tools

test: .fmtcheck $(go_files)
	go test -coverprofile=coverage.out -json ./... > testreport.json

build: test $(go_files)
	CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/app ./cmd/...

default: build

.DEFAULT_GOAL := default

.PHONEY: default clean%
