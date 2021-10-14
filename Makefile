COVER_FILE=coverage.txt

.PHONY: setup
setup:
	@go mod download

.PHONY: fmt
fmt:
	@gofmt -l -w -e .
	@goimports -w .

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: test cover cover_output
test:
	@git clean -fdx ${COVER_FILE}
	@go test -v -race -covermode=atomic -coverprofile=${COVER_FILE} -coverpkg ./...

cover: test
	@go tool cover -func ${COVER_FILE}

cover_output: test
	@go tool cover -html ${COVER_FILE}





