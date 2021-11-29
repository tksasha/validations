all: test

test:
	@go test

fmt:
	@find . -name \*.go -exec go fmt {} \;

run:
	@go run examples/*.go
