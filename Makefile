build:
	go build

fmt:
	gofmt -w .
	goimports -w .
