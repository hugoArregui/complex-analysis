package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/hugoArregui/complex-analysis/internal/julia"
)

func main() {
	output := flag.String("output", "build/out.png", "")
	realPart := flag.Float64("real", 0, "")
	imPart := flag.Float64("im", 0, "")

	flag.Parse()
	width := 1000
	height := 1000
	im := image.NewRGBA(image.Rect(0, 0, width, height))

	p := julia.Parameters{
		C:       complex(*realPart, *imPart),
		MaxIter: 1000,
		XMin:    -3.0,
		XMax:    3.0,
		YMin:    -3.0,
		YMax:    3.0,
	}

	julia.DrawJuliaSet(im, p)

	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}

	if err = png.Encode(file, im); err != nil {
		log.Fatal(err)
	}

	file.Close()
}
