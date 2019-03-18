package main

import (
	"github.com/fogleman/gg"
	"github.com/hugoArregui/julia-sets/internal/julia"
)

func main() {
	dc := gg.NewContext(1000, 1000)

	p := julia.Parameters{
		C:       complex(1.0/4.0, 0),
		MaxIter: 1000,
		XMin:    -5.0,
		XMax:    5.0,
		YMin:    -5.0,
		YMax:    5.0,
	}

	julia.DrawJuliaSet(dc, p)
	dc.SavePNG("build/out.png")
}
