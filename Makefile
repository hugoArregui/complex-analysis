build:
	go build -o build/julia ./cmd/julia

fmt:
	gofmt -w .
	goimports -w .

bench: build
	go test -bench=. github.com/hugoArregui/complex-analysis/internal/julia

.PHONY: build fmt
