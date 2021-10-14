VERSION=0.0.1
COVER_FILE=.coverage

.PHONY: setup
setup:
	@go mod download

.PHONY: fmt
fmt:
	@gofmt -l -w -e .

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: test cover cover_output

TEST_PARAMS = -v -race -failfast -covermode=atomic

test:
	@git clean -fdx ${COVER_FILE}
	@go test ${TEST_PARAMS} -coverprofile=${COVER_FILE} -coverpkg ./...

cover: test
	@go tool cover -func ${COVER_FILE}

cover_output: test
	@go tool cover -html ${COVER_FILE}





