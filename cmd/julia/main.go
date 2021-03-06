package main

import (
	"flag"
	"image"
	"log"
	"math/rand"

	"github.com/hugoArregui/complex-analysis/internal/julia"
	"github.com/hugoArregui/complex-analysis/internal/utils"
)

func main() {
	output := flag.String("output", "build/out.png", "")
	realPart := flag.Float64("real", 0, "")
	imPart := flag.Float64("im", 0, "")
	xMin := flag.Float64("xMin", -3.0, "")
	xMax:= flag.Float64("xMax", 3.0, "")
	yMin := flag.Float64("yMin", -3.0, "")
	yMax := flag.Float64("yMax", 3.0, "")
	seed := flag.Int64("seed", 1, "Random seed")

	flag.Parse()

	width := 1000
	height := 1000
	im := image.NewRGBA(image.Rect(0, 0, width, height))

	rand.Seed(*seed)

	p := julia.Parameters{
		C:       complex(*realPart, *imPart),
		MaxIter: 1000,
		XMin:    *xMin,
		XMax:    *xMax,
		YMin:    *yMin,
		YMax:    *yMax,
	}

	julia.Draw(im, p)

	if err := utils.Save(im, *output); err != nil {
		log.Fatal(err)
	}
}
