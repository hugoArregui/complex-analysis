build:
	go build -o build/julia ./cmd/julia

fmt:
	gofmt -w .
	goimports -w .

bench: build
	go test -bench=. github.com/hugoArregui/complex-analysis/internal/julia

examples: build
	build/julia --output=examples/example1.png --real=0.25
	build/julia --output=examples/example2.png --im=1
	build/julia --output=examples/example3.png --real=-1
	build/julia --output=examples/example4.png --real=1

.PHONY: build fmt
