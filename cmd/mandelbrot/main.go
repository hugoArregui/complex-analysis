package main

import (
	"flag"
	"image"
	"log"
	"math/rand"

	"github.com/hugoArregui/complex-analysis/internal/mandelbrot"
	"github.com/hugoArregui/complex-analysis/internal/utils"
)

func main() {
	output := flag.String("output", "build/out.png", "")
	xMin := flag.Float64("xMin", -2.0, "")
	xMax:= flag.Float64("xMax", 0.75, "")
	yMin := flag.Float64("yMin", -1.5, "")
	yMax := flag.Float64("yMax", 1.5, "")
	seed := flag.Int64("seed", 1, "Random seed")

	flag.Parse()

	width := 1000
	height := 1000
	im := image.NewRGBA(image.Rect(0, 0, width, height))

	rand.Seed(*seed)
	p := mandelbrot.Parameters{
		MaxIter: 1000,
		XMin:    *xMin,
		XMax:    *xMax,
		YMin:    *yMin,
		YMax:    *yMax,
	}

	mandelbrot.Draw(im, p)

	if err := utils.Save(im, *output); err != nil {
		log.Fatal(err)
	}
}
