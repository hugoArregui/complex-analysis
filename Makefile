SEED=1

build:
	go build -o build/julia ./cmd/julia
	go build -o build/mandelbrot ./cmd/mandelbrot

fmt:
	gofmt -w .
	goimports -w .

bench: build
	go test -bench=. github.com/hugoArregui/complex-analysis/internal/julia

examples: build
	build/julia --seed=${SEED} --output=examples/example1.png --real=0.25
	build/julia --seed=${SEED} --output=examples/example2.png --im=1
	build/julia --seed=${SEED} --output=examples/example3.png --real=-1
	build/julia --seed=${SEED} --output=examples/example4.png --real=1
	build/mandelbrot --seed=${SEED} --output=examples/example5.png
	build/mandelbrot --seed=${SEED} --xMin=0 --yMin=0 --xMax=0.5 --yMax=0.5 --output=examples/example6.png
	build/mandelbrot --seed=${SEED} --xMin=0.3 --yMin=0 --xMax=0.5 --yMax=0.5 --output=examples/example7.png
	build/mandelbrot --seed=${SEED} --xMin=0.43 --yMin=0 --xMax=0.5 --yMax=0.5 --output=examples/example8.png
	build/mandelbrot --seed=${SEED} --xMin=0.43 --yMin=0.2 --xMax=0.5 --yMax=0.3 --output=examples/example9.png
	build/mandelbrot --seed=${SEED} --xMin=0.43 --yMin=0.2 --xMax=0.5 --yMax=0.25 --output=examples/example10.png
	build/mandelbrot --seed=${SEED} --xMin=0.43 --yMin=0.2 --xMax=0.45 --yMax=0.25 --output=examples/example11.png

play: build
	build/mandelbrot --seed=${SEED} --xMin=0 --yMin=0 --xMax=0.5 --yMax=0.5 --output=playground/p1.png
	build/mandelbrot --seed=${SEED} --xMin=0.3 --yMin=0 --xMax=0.5 --yMax=0.5 --output=playground/p2.png
	build/mandelbrot --seed=${SEED} --xMin=0.43 --yMin=0 --xMax=0.5 --yMax=0.5 --output=playground/p3.png
	build/mandelbrot --seed=${SEED} --xMin=0.43 --yMin=0.2 --xMax=0.5 --yMax=0.3 --output=playground/p4.png
	build/mandelbrot --seed=${SEED} --xMin=0.43 --yMin=0.2 --xMax=0.5 --yMax=0.25 --output=playground/p5.png
	build/mandelbrot --seed=${SEED} --xMin=0.43 --yMin=0.2 --xMax=0.45 --yMax=0.25 --output=playground/p6.png
	firefox playground/*.png

.PHONY: build fmt
