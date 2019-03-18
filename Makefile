build:
	go build -o build/julia ./cmd/julia

fmt:
	gofmt -w .
	goimports -w .

bench: build
	go test -bench=. github.com/hugoArregui/julia-sets/internal/julia

.PHONY: build fmt
