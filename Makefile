VERSION=0.0.1
COVER_FILE=./coverage.html

.PHONY: setup
setup:
	@go mod download

.PHONY: fmt lint
fmt:
	@gofmt -l -w -e .

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: test cover cover_output

TEST_PARAMS = -v -race -failfast -covermode=atomic

test:
	@git clean -fdx ${COVER_FILE}
	@go test ${TEST_PARAMS} -coverprofile=${COVER_FILE} -coverpkg=./... -timeout=10s ./...

cover: test
	@go tool cover -func ${COVER_FILE}

cover_output: test
	@go tool cover -html ${COVER_FILE}





