install:
	@go version
	@go get -v golang.org/x/lint/golint
.PHONY: install

test:
	@go version
	@golint -set_exit_status ./...
	@go vet ./...
	@go test -cover ./...
.PHONY: test

coverage:
	@bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage
